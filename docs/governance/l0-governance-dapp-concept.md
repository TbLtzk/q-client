# Technical Concept: Better UX for L0 Governance

## Goal

Make Layer 0 governance usable for non-technical root nodes by replacing the current `geth attach` / `gov.*` console workflow with a dapp-driven signing flow that feels similar to existing MetaMask-based interactions in HQ.

The desired user experience is:

1. A root node opens an L0 governance page in HQ.
2. The dapp shows pending root-list or exclusion-list proposals.
3. The root node reviews the change in a human-friendly format.
4. The root node signs the proposal with MetaMask.
5. The signed message is submitted to any reachable Q node and propagated through the existing governance network.

This document answers:

- whether signing outside the client is feasible,
- what must change in `q-client`,
- and what a safe implementation could look like.

## Problem Statement

Today, L0 governance is operationally difficult for many root nodes because it requires direct console access and low-level commands on the `gov` interface, as documented in the [QGOV L0 governance guide](https://docs.qgov.io/layer0-governance/#layer-0-governance-of-root-nodes).

In practice, many root nodes do not reliably:

- inspect pending proposals,
- accept valid proposals,
- propose updates themselves,
- or verify what exactly they are signing.

The current UX is therefore a governance participation bottleneck and an operational risk.

## Current State In `q-client`

### What already exists

`q-client` already has the core mechanics needed for off-chain L0 governance:

- canonical data structures for root lists and validator exclusion lists in `common/governance.go`,
- deterministic hashing and signature recovery in `governance/root_list.go` and `governance/exclusion_set.go`,
- a dedicated `qgov` p2p protocol that propagates signed governance sets in `governance/handler.go`,
- read RPC endpoints via public namespace `govPub`,
- mutating RPC endpoints via private namespace `gov`.

The important observation is that propagation is already implemented. Once a node accepts a signed proposal, it can relay that proposal to peers.

### What is missing

The current mutating flow assumes that the receiving node itself:

- is a root node,
- has the relevant account locally available,
- and can sign the governance payload inside the client.

This is implemented through methods such as:

- `ProposeRootListUpdate()`
- `ProposeExclusionListUpdate()`
- `AcceptProposedRootList()`
- `AcceptProposedExclusionList()`

Those methods route into `RootManager` logic that signs locally and therefore does not support "submit an already signed governance message from outside the client".

Additionally, `gov` is intentionally blocked from exposure over external HTTP/WS APIs in `cmd/utils/flags.go`.

## Feasibility Assessment

## 1. Can L0 governance messages be signed outside the client?

Yes, in principle.

The current signature model is simple:

- a root-list signature is a signature over the deterministic root-list hash,
- an exclusion-list signature is a signature over the deterministic exclusion-list hash,
- signatures are verified by recovering the signer from the hash and signature.

So there is no fundamental requirement that signatures must be produced inside `q-client`.

However, there is an important UX and compatibility caveat:

- the current implementation expects a raw signature over the raw governance hash,
- MetaMask does not provide a clean UX for arbitrary raw-hash signing,
- and commonly used wallet methods such as `personal_sign` or `eth_signTypedData_v4` produce different signature semantics than the current implementation expects.

Conclusion:

- external signing is technically feasible,
- MetaMask-native signing is not a clean fit with the current raw-signature format,
- therefore a good dapp UX likely requires a new wallet-friendly signing format, not just a new RPC endpoint.

## 2. Can signed L0 governance messages be propagated via any node, including a public RPC endpoint?

Yes, with client changes.

The p2p layer already shows that a node can ingest and rebroadcast a signed governance set even if the proposal did not originate on that node. In other words, the missing piece is not network propagation but public ingestion.

Today there is no public RPC for:

- submitting a signed root list,
- submitting a signed exclusion list,
- merging additional signatures into an existing pending proposal.

So the concept is feasible, but `q-client` needs a dedicated external submission API.

## Recommended Product Direction

Implement L0 governance in HQ as a dedicated subpage for root nodes, backed by new submission APIs in `q-client`.

For good UX and security, the preferred direction is:

- keep propagation on the existing `qgov` p2p network,
- add public read and submission APIs for L0 governance,
- use an explicit typed message format for MetaMask-friendly signing,
- preserve compatibility with **`qgov/5`** peers while introducing **`qgov/6`** for typed relay where needed ([phased rollout](#phased-rollout-qgov-versions-typed-relay-and-raw-signatures)).

## Design Options

## Option A: Minimal-change path

### Summary

Keep the existing signature format and add RPC methods that accept already signed governance payloads.

### How it works

1. The dapp constructs the same canonical root list or exclusion list payload as today.
2. An external signer signs the raw governance hash.
3. The dapp submits the signed payload to a node through a new RPC method.
4. The receiving node validates it and injects it into the same internal flow used for p2p-received governance sets.
5. The node propagates it through `qgov`.

### Advantages

- smallest change to core governance logic,
- minimal changes to hashing and validation,
- reuses current p2p propagation as-is.

### Disadvantages

- weak wallet UX,
- poor MetaMask fit,
- less user transparency about exactly what is being signed,
- harder to present clear signing prompts in the dapp.

### Verdict

Good as an engineering stepping stone, but not ideal as the long-term UX solution.

## Option B: Wallet-friendly path

### Summary

Introduce a new explicit signing scheme for dapp-based L0 governance, ideally based on typed structured data, and let any node accept and relay those signed submissions.

### How it works

1. The dapp fetches read-heavy governance state from the existing HQ/indexer monitoring APIs where available.
2. The dapp builds a proposal object and shows a diff to the root node.
3. MetaMask signs a typed governance message.
4. The dapp submits the signed typed payload to a public submission endpoint.
5. The node verifies the typed signature. **Network-wide quorum still uses raw signatures over the governance list hash on the existing `RootListMsg` / `ExclusionListMsg` rails** until a later rollout phase (see [Phased rollout: `qgov` versions, typed relay, and raw signatures](#phased-rollout-qgov-versions-typed-relay-and-raw-signatures)). Until then, typed material is relayed on a separate path (e.g. `qgov/6` only) and root nodes that recognize their own attestation mint the corresponding raw signature locally for ingestion and legacy propagation.

### Advantages

- best UX for root nodes,
- clear and auditable signing prompts,
- lower risk of users signing opaque hashes,
- natural fit for an HQ-integrated workflow.

### Disadvantages

- more code changes,
- requires a versioned signing format,
- requires careful compatibility design between legacy console flow and dapp flow.

### Verdict

This is the recommended target architecture.

## Phased rollout: `qgov` versions, typed relay, and raw signatures

This section refines **Option B** for **fleet compatibility**: MetaMask-friendly **typed** signing is the UX target, while **raw** signatures over the list hash remain what existing `qgov` list messages verify until a coordinated later phase.

### `qgov/5` vs `qgov/6` (typed relay gating)

| Version | Typed attestation relay to peer | Notes |
| ------- | --------------------------------- | ----- |
| **`qgov/5`** | **No** (do not send typed-relay messages on this sub-protocol) | Matches today’s deployed behavior; peers must not receive payloads they cannot interpret. |
| **`qgov/6`** | **Yes** (dedicated typed-relay messages between peers that negotiated `6`) | New wire message(s) or codes; only forwarded along connections where both sides speak `6`. |

New clients **implement both** protocol versions so they remain compatible with **`qgov/5`-only** peers while opening **`qgov/6`** sessions to upgraded peers. Typed relay traffic is **not** forwarded onto `qgov/5` links.

*Implementation note:* `governance/protocol.go` currently exposes only **`qgov/5`**; adding **`qgov/6`** is the deliberate interface bump for this behavior.

### Signature deduplication (explicit rule)

- **Storage / merge deduplication** (e.g. `signers` map keys, `mergeSignatures`): keep the **existing rule** — **one entry per recovered signer address** (the Ethereum address recovered from the signature bytes). Do **not** introduce a new deduplication key scheme unless a later phase explicitly changes it.
- **Quorum and threshold counting** (e.g. `isAcceptable`, alias-aware known signers): keep **existing** active-root and **alias** rules; they are unchanged by this specification.

When **raw** and **typed** encodings can both represent participation by the **same logical root slot** (e.g. different recovered addresses for alias vs main), preventing **double counting** for threshold purposes may require extra rules on top of “map key = recovered address” — see [Root-node alias and signing identity](#root-node-alias-and-signing-identity) and [Open decisions](#open-decisions-after-this-spec).

### Dual verification on list messages (`Signatures[]`)

Applies to `RootListMsg` / `ExclusionListMsg` when a client accepts more than one scheme in `Signatures[]` (defensive in Phase 1; production-mixed in Phase 2). **`qgov/6` typed relay** uses **typed-only** verification (same as RPC); this section does not apply there.

**Try order (single pass — do not run the second scheme if the first succeeds):**

| Phase | First try | Fallback |
| ----- | --------- | -------- |
| **1** | **Raw** over `list.Hash` | **Typed** over the versioned payload digest |
| **2** | **Typed** over the versioned payload digest | **Raw** over `list.Hash` |

Rationale: match the likely on-wire encoding and avoid building/hashing typed payloads on every signature when p2p list traffic is overwhelmingly raw (Phase 1).

**Dual success (very unlikely):** If a blob would verify under both schemes, treat it as the scheme that was **tested first** (raw in Phase 1, typed in Phase 2). There is no need to always run both verifications.

**Per-entry misses:** If neither scheme verifies, **skip** that `Signatures[]` entry and continue with the rest. Do **not** fail the whole message for a dual-verify miss.

**Whole-message failure:** Fail only on **structural** errors (e.g. invalid or empty node sets, declared `list.Hash` mismatch with calculated hash, decode failures) — the same class of hard rejects as today before merge/import proceeds.

### Root-node alias and signing identity

**Decision:** When an active root node has a **distinct alias** (`root != alias` in the active root set’s alias map), **external typed signing must use the alias address in MetaMask** (the same address the node uses for in-client governance signing via `signRootSet`). Signing with the **main/root account** when an alias is configured is **invalid** for typed submissions and typed relay.

**Rationale:** In-client signing always uses the alias key when one exists (`governance/root_manager.go`, `signRootSet`). Requiring the same identity off-chain avoids ambiguous “two keys, one slot” behavior and matches quorum counting (`knownSigners` only treats alias signatures as votes when `root != alias`).

**`q-client` behavior today (Epic B):**

| Path | Main-account sig when alias exists |
| ---- | ---------------------------------- |
| **Typed RPC** (`SubmitTypedSigned*`) | **Rejected** with error (e.g. `typed root list contains unknown signer`) — recovered main address is not in `active.knownSigners()`. Tests: `typed signature from root key instead of alias is rejected`. |
| **Typed relay** (`qgov/6`, planned) | Must apply the **same** membership rule: **do not relay** attestations that fail typed signer checks. |
| **Raw p2p / `newRootSet`** | **Not rejected at crypto parse time** — a valid ECDSA sig over `list.Hash` from the main key is stored in `signers`. It **does not count** toward threshold (`isAcceptable` / `knownSigners`). A **new** proposal whose only signatures are main-key may be **silently ignored** (`importRootSet` default branch: `len(knownSigners)==0` → debug log, `nil` error). Merges on an existing hash can still **store** a main-key entry in `signers` without advancing quorum. |

**Product / API expectations:**

- Dapps should **detect alias vs main** (indexer / `govPub` helpers) and tell the user which address to connect in MetaMask before signing.
- Typed RPC should return a **clear, stable error** when the recovered signer is a known root main address but not the required alias (implementation improvement over generic “unknown signer”).
- Phase 1 **self-bridge** (RN mints raw after seeing its own typed attestation): use the **alias** account when configured, same as `signRootSet` — not the main account from a mistaken MetaMask connection.

**When `root == alias`:** No separate alias; main and alias are the same address — typed and raw may use that single address.

### Phase 1 — Typed relay only

**Goals:** deliver typed attestations between **upgraded** peers without breaking **`qgov/5`-only** nodes; keep **raw** signatures as the only votes that merge into list state for threshold **on the legacy list messages** in production.

- New clients can **receive and verify** typed attestations (same membership / replay guards as for raw, as implemented for each path).
- New clients **propagate** typed attestations **only** to peers with **`qgov/6`** negotiated (see above).
- New clients **do not** treat typed attestations as full substitutes for raw signatures in internal **quorum state** in this phase: they may track relayed attestations (e.g. dedupe by content hash + signer + proposal id) to **avoid rebroadcast storms**, but **threshold / upgrade** continues to follow **raw** signatures on the normal list ingestion path.
- On the **existing** root-list / exclusion-list messages (`RootListMsg`, etc.), new clients use [dual verification on list messages](#dual-verification-on-list-messages-signatures) (**raw first**, then typed; skip per-entry misses). **Production nodes do not place typed signature bytes into `Signatures[]` in Phase 1**; dual verification exists for **tests**, **defense in depth**, and forward compatibility.
- When a new client receives a valid typed attestation whose recovered signer is **this node’s governance signing account** (the **alias** if configured, otherwise the root address), it **mints the corresponding raw signature** with that same account (matching `signRootSet`) and **ingests** it into the normal set / merge path — so the vote can propagate on **`qgov/5`** rails without requiring the operator to RPC their own validator.

### Phase 2 — Typed ingestion on list messages

- New clients **fully ingest** typed signatures into the same internal structures as raw and propagate merged lists on the **existing** raw-hash rails (both may appear in `Signatures[]` from different signers). Per-signature verification on lists uses [dual verification](#dual-verification-on-list-messages-signatures) with **typed first**, then raw fallback.
- The **dedicated typed relay** path may still exist for **ingress** (e.g. first receipt from dapp or peer) but **forwarding** typed relay messages to peers is de-emphasized once list messages carry typed bytes — exact “relay vs list only” split is an implementation detail to lock during Phase 2.
- **Old (`qgov/5`-only) clients** that cannot verify typed bytes inside `Signatures[]` will **fail** validation — this phase is therefore a **fleet coordination** milestone (upgrade threshold before enabling production mixed lists).
- **Same signer must not** contribute both a raw and a typed signature if that would **double-count** toward threshold; enforcement must align with the [deduplication](#signature-deduplication-explicit-rule) rules and any [open decisions](#open-decisions-after-this-spec) on identity slots.

### Phase 3 — Abandon raw-over-hash for verification

- New clients **no longer accept** legacy raw-over-`list.Hash` signatures for new submissions (verification hardening / deprecation).
- Dedicated **typed-only relay** helpers may be **removed** if redundant.
- Requires an **activation policy** (network config, flag date, or minimum peer version) so operators can plan upgrades.

## Proposed Architecture

## Components

### 1. HQ dapp L0 governance module

Responsibilities:

- discover whether the connected wallet is an active root node,
- reuse the existing root-node monitoring/indexer read model for active / proposed root and exclusion lists where possible,
- show diffs against on-chain or active state without duplicating the current monitoring dashboard,
- create proposals,
- collect signatures,
- submit signed payloads to a public node endpoint,
- show signature status and threshold progress.

### 2. HQ/indexer read model

HQ already has a root-node monitoring page backed by indexer endpoints for L0 root lists, exclusion lists, signer data, and operational metrics. The L0 governance UX should treat those indexer-backed APIs as the preferred source for read-heavy dashboard state because they can cache and aggregate data that would otherwise be redundantly fetched by every UI client from RPC nodes.

The dedicated governance page should therefore be an action and review layer on top of the existing monitoring approach, not a parallel read dashboard. It should reuse existing monitoring components and helpers where they fit, and only add new read endpoints when the current indexer model cannot answer an action-critical question.

### 3. Public governance submission API in `q-client`

New externally safe RPC methods should be introduced for signed submissions and canonical signing helpers.

Suggested examples:

- `govPub_l0GovernanceCapabilities`
- `govPub_submitSignedRootList`
- `govPub_submitSignedExclusionList`
- `govPub_appendRootListSignature`
- `govPub_appendExclusionListSignature`
- `govPub_getGovernanceSigningPayload`
- `govPub_getGovernanceProposalStatus`

These methods should be limited to signed payload ingestion, canonical signing payload construction, and narrowly scoped status checks that are required for safe signing. They should not expose arbitrary internal `gov` mutators and should not replace the indexer as the default source for read-heavy dashboard data.

The dapp should probe the RPC endpoint configured for the connected wallet before enabling signing or submission controls. If the node does not expose the required L0 governance capabilities, the dapp should keep monitoring/read-only views usable, disable action controls, and explain that the selected RPC endpoint does not yet support external L0 governance signing/submission.

### govPub JSON wire format for list signatures

For `govPub_submitSignedRootList`, `govPub_submitTypedSignedRootList`, `govPub_submitSignedExclusionList`, and `govPub_submitTypedSignedExclusionList`, each element of `signatures` must be a **`0x`-prefixed hex string** (65 bytes for a standard ECDSA signature), consistent with `hash` and `address` fields elsewhere in JSON-RPC. **Base64 is not accepted** for `signatures` on these payloads.

Example (typed root list submit):

```json
"signatures": ["0x<r><s><v>"]
```

Pass the value returned by `eth_signTypedData_v4` / wallet `signTypedData` **without** re-encoding the recovery byte: q-client normalizes Yellow Paper `v` (`27` / `28`) to the internal `0` / `1` form before verification.

List responses over JSON-RPC marshal `signatures` the same way (`0x` hex). **p2p / RLP** list messages are unchanged.

### 4. Existing `qgov` p2p propagation

The existing `RootListMsg` / `ExclusionListMsg` / handshake payloads remain the **canonical carrier for raw ECDSA signatures** over the deterministic list hash (`list.Hash`), because that is what all deployed clients verify today (`governance/root_list.go`, `governance/exclusion_set.go`).

**Typed governance attestations must not be placed inside `Signatures[]` on those messages until enough of the fleet supports verifying them there** (see phased rollout below). Until then, typed relay uses an **additional** negotiation surface.

**`qgov` protocol versions (compatibility gate):**

- **`qgov/5`** — behavior as today: root/exclusion list messages and existing message codes only; **no** forwarding of typed-relay messages to peers (and no expectation that peers accept them).
- **`qgov/6`** — extends the protocol with dedicated wire messages (or codes) for **typed attestation relay** between peers that negotiated this version. New clients **register both** `qgov/5` and `qgov/6` handlers so they interoperate with old peers on `5` while using `6` with upgraded peers for typed relay.

Today `q-client` only implements **`qgov/5`** in `governance/protocol.go`; introducing **`qgov/6`** is the explicit bump that gates “relay typed to peers who can accept it.” RPC capability discovery (`govPub_l0GovernanceCapabilities`, etc.) remains complementary for **submission and signing helpers**, not a substitute for this peer negotiation.

See [Phased rollout: `qgov` versions, typed relay, and raw signatures](#phased-rollout-qgov-versions-typed-relay-and-raw-signatures) for how `5` vs `6` interact with phased ingestion.

### 5. Governance state ingestion layer

Internally, `q-client` should gain a new "import signed governance set" path that:

- validates the payload,
- validates signatures,
- checks signer membership against the active root set,
- applies spam protection and quota rules,
- merges signatures when the payload already exists locally,
- upgrades `proposed` to `desired` or `active` when thresholds are met,
- emits the same broadcast events already used by peer-originated governance messages.

## Required `q-client` Changes

## 1. Add a public submission API

Files primarily affected:

- `governance/api.go`
- `governance/governance.go`
- `internal/web3ext/web3ext.go`
- potentially `ethclient/ethclient.go`

Changes:

- add public RPC methods for signed governance submission,
- add a public capabilities method that advertises supported L0 governance submission methods, signing schemes, signing-payload versions, and status-helper support,
- keep them in `govPub` or another explicitly public namespace,
- do not expose the existing unrestricted `gov` mutators externally.

## 2. Refactor local-signing and external-ingestion paths

Files primarily affected:

- `governance/root_manager.go`
- `governance/handler.go`

Changes:

- extract a shared ingestion path from the existing p2p handlers,
- reuse it for both peer-received and RPC-submitted signed payloads,
- separate "create and sign locally" from "validate and import externally signed proposal".

This is the most important code-structure improvement.

## 3. Introduce a versioned signing format

Files primarily affected:

- `common/governance.go`
- `governance/root_list.go`
- `governance/exclusion_set.go`
- new helper files for typed payload creation and verification

Changes:

- define a versioned signing payload for root lists and exclusion lists,
- include chain ID / network ID and purpose in the signed data to avoid replay and cross-context confusion,
- support canonical serialization rules,
- preserve legacy verification for existing console-driven flows if backward compatibility is needed.

Recommended fields for signing payload:

- `domain`: `"Q-L0-Governance"`
- `version`
- `chainId` or `networkId`
- `proposalType`: `rootList` or `exclusionList`
- `timestamp`
- `payloadHash`
- human-relevant fields such as node addresses or validator block ranges

## 4. Support MetaMask-oriented signing

If the goal is real HQ integration, the preferred client-side contract should be:

- the dapp asks MetaMask to sign typed data,
- `q-client` verifies that typed data signature,
- `q-client` maps the recovered signer to the active root aliases / addresses.

If the project wants the smallest possible first step, raw-hash signing can be supported first, but this should be considered transitional.

## 5. Keep spam and abuse protection

Files primarily affected:

- `governance/handler.go`
- `governance/root_manager.go`

Changes:

- reuse existing quota checks for RPC-submitted proposals,
- add request-size limits and rate limiting at the RPC boundary,
- reject malformed payloads early,
- reject signatures from non-root signers,
- reject stale proposals based on timestamp and current state.

This matters especially if public RPC nodes are allowed to accept L0 governance submissions.

## 6. Add action-critical governance helper APIs for dapp UX

HQ should continue to use its existing indexer-backed monitoring APIs for read-heavy dashboard state such as active/proposed L0 root lists, exclusion lists, signer data, and operational metrics. This avoids many UI clients repeatedly asking public RPC nodes for the same data and keeps caching/aggregation in the indexer layer.

`q-client` may still need narrowly scoped read helpers for action-critical signing flows:

- endpoint capability discovery,
- proposal diff preview,
- signature threshold progress,
- signer list with alias resolution,
- normalized active / proposed / desired status,
- "needs your signature" status for the connected root node.

These helpers should be treated as canonical signing/status checks, not as a replacement for the existing monitoring read model.

## Signature Strategy Recommendation

## Recommended approach

Use a typed, versioned off-chain signature format for L0 governance.

Rationale:

- MetaMask supports it well,
- users can review structured content before signing,
- the payload can include anti-replay metadata,
- it provides a safer UX than signing opaque raw hashes.

## Backward compatibility

The client should support both:

- legacy raw-signature governance messages from console-driven nodes,
- new typed-signature governance messages from dapp-driven nodes.

There are two reasonable **long-term** compatibility models:

**Fleet behavior** (typed relay, `qgov/5` vs `qgov/6`, when typed bytes may appear in `Signatures[]`) is specified in [Phased rollout: `qgov` versions, typed relay, and raw signatures](#phased-rollout-qgov-versions-typed-relay-and-raw-signatures) above.

### Model 1: RPC accepts typed signatures; list messages stay raw-first during early phases

In this model (aligned with **Phase 1** in that section):

- the public API verifies typed signatures where exposed;
- root nodes **self-bridge** to raw for votes that must propagate on **`qgov/5`**;
- typed attestations **relay on `qgov/6` only** and are not placed in production `Signatures[]` until **Phase 2** coordination.

This avoids breaking **`qgov/5`-only** peers while MetaMask-first UX rolls out.

### Model 2: Governance protocol evolves to carry signature metadata versioning

In this model:

- both RPC and p2p support multiple signature schemes explicitly on the wire (**Phase 2+**).

This is cleaner long-term, but larger in scope.

## Recommendation

Start with **Model 1** during **Phase 1–2**, then converge toward a single verification path in **Phase 3** if the project commits to deprecating raw-over-hash.

## Suggested HQ / product rollout

The **phase numbers below** are **product / engineering milestones** (RPC, refactor, HQ). They are **not** the same as [governance network phases](#phased-rollout-qgov-versions-typed-relay-and-raw-signatures) (typed relay, `qgov/6`, list ingestion, raw deprecation); align both timelines when planning releases.

## Phase 1: Technical enablement

- add public submission endpoints for signed root and exclusion lists,
- refactor internal ingestion logic so RPC and p2p share the same validation path,
- keep legacy raw-signature scheme for an initial prototype.

Outcome:

- externally signed proposals become technically possible,
- even before the full HQ integration is complete.

## Phase 2: Wallet-friendly signing

- define typed signing payloads,
- implement verification in `q-client`,
- expose payload-building helpers for the dapp,
- add status endpoints optimized for UI use.

Outcome:

- MetaMask-based signing becomes practical.

## Phase 3: HQ integration

- add an L0 governance subpage,
- reuse existing root-node monitoring/indexer data for read-heavy L0 state,
- show proposal diffs and signer progress in a focused action/review UX,
- support propose, accept, and co-sign flows,
- guide users with warnings and confirmations.

Outcome:

- root nodes can perform L0 governance without console access.

## Security Considerations

- Public submission endpoints must not be equivalent to exposing the private `gov` namespace.
- Signed submissions should be the only mutating operations exposed publicly.
- The dapp should disable signing/submission controls when the connected RPC endpoint does not advertise the required L0 governance capabilities.
- Payloads must contain network/domain separation data to prevent replay across networks or contexts.
- Timestamp and staleness validation must remain strict.
- Signature counting must continue to respect active root membership and alias resolution.
- Rate limiting and payload size limits are recommended for public endpoints.

## Risks And Open Questions

### Open decisions (after this spec)

The following are **not** fully pinned by this document; they need explicit decisions before implementation tickets close:

1. **`qgov/6` wire shape** — Exact new message codes, payload schema, max size, and whether typed relay is **request/response** or **fire-and-forget gossip** (and how it interacts with `StatusMsg` / handshake).
2. ~~**Phase 1/2 dual-verify on list messages**~~ — **Resolved:** single-pass order (**raw → typed** in Phase 1, **typed → raw** in Phase 2); if both would verify (very unlikely), credit the **first** scheme tried (no need to run both); **skip** entries that fail both; **fail whole message** only on structural errors. See [Dual verification on list messages](#dual-verification-on-list-messages-signatures). Implemented in `governance/signature_parse.go` (`dualVerifyTypedFirst` flips in Phase 2).
3. ~~**External vs alias typed signer**~~ — **Resolved:** use **alias in MetaMask** when an alias is configured; reject main-only typed sigs at RPC and typed relay; self-bridge mints raw with alias. See [Root-node alias and signing identity](#root-node-alias-and-signing-identity).
4. **Phase 2 double-counting** — Reject or collapse **main + alias** encodings for the same logical root where only alias counts for quorum; Phase 1 policy avoids typed-from-main, but raw main-key bytes may still appear on legacy merges until stricter rejection is added.
5. **Phase 2→3 activation** — Network-wide flag, minimum peer fraction on `qgov/6`, calendar date, or on-chain config; who flips it.
6. **Unknown `qgov` message handling on `5`** — Ensure old peers never receive `6`-only messages on the wrong session; define behavior if a buggy peer sends them anyway (ignore vs disconnect).
7. **Typed relay reachability** — Typed traffic only flows on **`qgov/6`** edges. If the graph of `6` sessions is disconnected, some root nodes may not receive attestations until they upgrade peers or use RPC ingress; decide whether **multi-hop relay** on `6`, **dedicated gateways**, or **operator expectations** are acceptable.

### Other risks

- MetaMask signing semantics need to be pinned down early for **typed** payloads (domain, types, chain id).
- Public RPC nodes may be unwilling to expose any governance-related mutators unless the submission API is tightly scoped and abuse-resistant.
- Alias policy for typed signing is fixed in [Root-node alias and signing identity](#root-node-alias-and-signing-identity); HQ must connect the wallet to the **alias** address when applicable.
- Capability discovery should not rely only on `rpc_modules` or deprecated RPC namespace version metadata; it should use an explicit governance capability response that the dapp can evaluate before enabling controls.

## Development backlog

This backlog is scoped as **engineering work** split between `q-client`, HQ (or another dapp), and supporting documentation. Estimates use **person-days** (one developer working full-time on that item); ranges reflect uncertainty (integration, review cycles, regression testing).

Assumptions:

- One engineer familiar with Go/`q-client` governance and Ethereum signing.
- HQ codebase is separate; HQ estimates are indicative only (depends on framework and existing patterns).

### Estimate legend


| Range        | Typical meaning                                                  |
| ------------ | ---------------------------------------------------------------- |
| **0.5–1 pd** | Small change, tests localized, low risk                          |
| **2–4 pd**   | Medium feature, touches several files, needs integration tests   |
| **5–10 pd**  | Larger refactor or new protocol surface, cross-team coordination |
| **10+ pd**   | Epic-sized; split into milestones                                |


Priority: **P0** = blocking minimal viable flow, **P1** = production UX / safety, **P2** = polish and ops.

---

### Epic A — `q-client`: ingest externally signed governance (Phase 1)


| ID  | Item                                                                                                                                                                                        | Priority | Depends on | Estimate   |
| --- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | ---------- | ---------- |
| A1  | Extract shared validation/import path used by both `handler.handleRootSet` / `handleExclusionSet` and a new RPC entry point; avoid duplicating quorum and alias logic                       | P0       | —          | **5–8 pd** |
| A2  | Add public RPC methods to submit signed `RootList` / `ValidatorExclusionList` payloads (names TBD); wire to shared import path; unit tests against existing fixtures                        | P0       | A1         | **3–5 pd** |
| A3  | Register APIs in `governance.Governance.APIs()`, `internal/web3ext/web3ext.go`; optional `ethclient` wrappers                                                                               | P0       | A2         | **1–2 pd** |
| A4  | Policy decision + implementation: expose only submission methods on HTTP/WS (relax or specialize `filterProtectedAPIs` for new namespace or allowlisted methods) without opening full `gov` | P0       | A2         | **2–4 pd** |
| A5  | Spam/abuse limits on public submission path (payload size, rate limits, configurable disable)                                                                                               | P1       | A2         | **2–4 pd** |
| A6  | Regression tests: p2p + RPC paths produce equivalent state for same signed payload                                                                                                          | P1       | A1–A2      | **3–5 pd** |


**Epic A subtotal (rough):** **16–28 pd**

---

### Epic B — Typed / MetaMask-friendly signing (Phase 2)


| ID  | Item                                                                                                                                                                                           | Priority | Depends on      | Estimate                              |
| --- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | --------------- | ------------------------------------- |
| B1  | Specify versioned EIP-712 (or equivalent) domain and types for root-list and exclusion-list proposals; document encoding rules matching on-chain-off-chain identity (chain ID, purpose string) | P0       | Epic A baseline | **3–5 pd**                            |
| B2  | Implement verification in `q-client`: recover signer from typed data, map to active root / alias rules, merge into internal sets                                                               | P0       | B1, A1          | **5–8 pd**                            |
| B3  | RPC: `govPub`-style methods returning **canonical signing payload** for a given proposal (helps dapps build identical typed data)                                                              | P1       | B1              | **2–4 pd**                            |
| B4  | Backward compatibility tests: legacy console/raw-hash flow unchanged; new typed path coexists                                                                                                  | P1       | B2              | **2–4 pd**                            |
| B5  | Security review pass (replay, cross-chain, ambiguous alias signing)                                                                                                                            | P1       | B2              | **1–3 pd** (review only; fixes extra) |


**Epic B subtotal (rough):** **13–24 pd**

---

### Epic C — HQ (or target dapp): L0 governance UI


| ID  | Item                                                                                                                                               | Priority | Depends on                   | Estimate    |
| --- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------- | ---------------------------- | ----------- |
| C1  | Navigation + route for L0 governance subpage; feature flag; reuse or deep-link from existing root-node monitoring UX                               | P0       | Existing monitoring route    | **1–2 pd**  |
| C2  | Connect wallet; detect root-node eligibility using existing indexer/monitoring root data plus wallet state                                         | P0       | Existing monitoring data     | **2–4 pd**  |
| C3  | Action/review views for active / proposed root and exclusion lists; reuse monitoring context/helpers and avoid a duplicate read dashboard           | P0       | C1–C2                        | **3–6 pd**  |
| C4  | Flow: review proposal → sign (typed data via Epic B, or interim raw-hash prototype) → submit to configurable RPC URL                               | P0       | Epic A (+ B for MetaMask UX) | **5–10 pd** |
| C5  | Signature progress indicator using indexer signer data where sufficient and q-client status helpers only for action-critical gaps                  | P1       | C3–C4                        | **2–4 pd**  |
| C6  | Copy, warnings, links to official docs; optional telemetry (no sensitive data)                                                                     | P2       | C4                           | **2–4 pd**  |


**Epic C subtotal (rough):** **19–38 pd**

---

### Epic D — Documentation, operations, release


| ID  | Item                                                                                                                                             | Priority | Depends on                 | Estimate     |
| --- | ------------------------------------------------------------------------------------------------------------------------------------------------ | -------- | -------------------------- | ------------ |
| D1  | Update public docs ([L0 governance](https://docs.qgov.io/layer0-governance/#layer-0-governance-of-root-nodes)) with dapp path + RPC requirements | P1       | Epic A (+ B as applicable) | **1–2 pd**   |
| D2  | Runbook: which RPC endpoints operators may expose; firewall / rate limit recommendations                                                         | P1       | A4–A5                      | **1–2 pd**   |
| D3  | Release notes, migration note for root nodes (console still supported)                                                                           | P2       | D1                         | **0.5–1 pd** |


**Epic D subtotal (rough):** **2.5–5 pd**

---

### Milestone rollup (indicative)


| Milestone             | Scope                       | Overlap-friendly calendar note                                          | Rough total   |
| --------------------- | --------------------------- | ----------------------------------------------------------------------- | ------------- |
| **M1 — Backend MVP**  | Epic A (A1–A4, minimal A5)  | Enables scripted or wallet-external signing without full MetaMask story | **~12–20 pd** |
| **M2 — Wallet-ready** | Epic B + remainder of A5–A6 + `qgov/6` typed relay (Phase 1) | Typed signing, relay gating, hardening                               | **~18–32 pd** |
| **M3 — HQ GA**        | Epic C + D                  | Parallelizable once M1 read/submit contract is stable                   | **~21–43 pd** |


**End-to-end (all epics, sequential pessimistic sum):** roughly **50–95 person-days** (~~**2.5–5 months** for one full-time engineer). With **parallel tracks** (e.g. one backend + one frontend once RPC shapes are fixed), calendar time can compress toward **~~2–3 months** depending on review and QA.

### Suggested sequencing

1. Lock **RPC shapes** and **security policy** for public submission (short design spike: **1–2 pd**).
2. Deliver **M1** so integrations can start without waiting for EIP-712.
3. Parallelize **M3** UI shell + action/review layer using existing HQ/indexer monitoring reads while backend finishes ingest.
4. Land **M2** before promoting MetaMask-first UX to non-technical root nodes.

