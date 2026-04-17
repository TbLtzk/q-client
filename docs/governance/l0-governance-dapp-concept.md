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
- switch the signature format for dapp-based signing to an explicit typed message format,
- preserve compatibility with the current peer protocol and internal governance state.

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

1. The dapp fetches the current governance state from `govPub`.
2. The dapp builds a proposal object and shows a diff to the root node.
3. MetaMask signs a typed governance message.
4. The dapp submits the signed typed payload to a public submission endpoint.
5. The node verifies the typed signature, converts it into the internal governance set representation, stores it as `proposed` or merges it into `desired` / `active`, and rebroadcasts it via `qgov`.

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

## Proposed Architecture

## Components

### 1. HQ dapp L0 governance module

Responsibilities:

- discover whether the connected wallet is an active root node,
- show active / proposed / desired root and exclusion lists,
- show diffs against on-chain or active state,
- create proposals,
- collect signatures,
- submit signed payloads to a public node endpoint,
- show signature status and threshold progress.

### 2. Public governance submission API in `q-client`

New externally safe RPC methods should be introduced for signed submissions.

Suggested examples:

- `govPub_submitSignedRootList`
- `govPub_submitSignedExclusionList`
- `govPub_appendRootListSignature`
- `govPub_appendExclusionListSignature`
- `govPub_getGovernanceSigningPayload`
- `govPub_getGovernanceProposalStatus`

These methods should be limited to signed payload ingestion and read operations. They should not expose arbitrary internal `gov` mutators.

### 3. Existing `qgov` p2p propagation

No fundamental protocol redesign is required. The existing relay logic should remain the transport used after a node accepts a submission.

### 4. Governance state ingestion layer

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

## 6. Add governance-specific read APIs for dapp UX

Existing `govPub` methods already expose useful state, but the dapp would benefit from higher-level read endpoints:

- proposal diff preview,
- signature threshold progress,
- signer list with alias resolution,
- normalized active / proposed / desired status,
- "needs your signature" status for the connected root node.

These are not strictly required for the backend architecture, but they are highly valuable for user adoption.

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

There are two reasonable compatibility models:

### Model 1: RPC accepts typed signatures, p2p continues to exchange legacy internal sets

In this model:

- the public API verifies a typed signature,
- converts it into the internal set representation,
- stores only the recovered signer and signature result,
- and propagates the canonical internal set through the existing protocol.

This is lower impact on the p2p layer, but it requires a clear mapping from typed-signature verification back into the existing internal signer model.

### Model 2: Governance protocol evolves to carry signature metadata versioning

In this model:

- both RPC and p2p support multiple signature schemes explicitly.

This is cleaner long-term, but larger in scope.

## Recommendation

Start with Model 1.

## Suggested Rollout Plan

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
- show proposal diffs and signer progress,
- support propose, accept, and co-sign flows,
- guide users with warnings and confirmations.

Outcome:

- root nodes can perform L0 governance without console access.

## Security Considerations

- Public submission endpoints must not be equivalent to exposing the private `gov` namespace.
- Signed submissions should be the only mutating operations exposed publicly.
- Payloads must contain network/domain separation data to prevent replay across networks or contexts.
- Timestamp and staleness validation must remain strict.
- Signature counting must continue to respect active root membership and alias resolution.
- Rate limiting and payload size limits are recommended for public endpoints.

## Risks And Open Questions

- MetaMask signing semantics need to be pinned down early. The project should avoid designing around raw-hash signing if the target UX is HQ + MetaMask.
- Backward compatibility with legacy console-based governance flows must be defined explicitly.
- Public RPC nodes may be unwilling to expose any governance-related mutators unless the submission API is tightly scoped and abuse-resistant.
- The exact treatment of aliases in off-chain signatures must be documented clearly so root nodes understand which account should sign.
