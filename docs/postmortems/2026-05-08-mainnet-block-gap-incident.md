# Mainnet Block Gap Incident Report (Preliminary)

**Date:** 2026-05-08  
**Network:** QGOV mainnet  
**Status:** Investigation ongoing  
**Prepared by:** Incident analysis (automated data collection)

## Executive Summary

A prolonged block production gap was observed between block `25979111` and block `25979112`.  
The measured time gap is approximately **10h 40m 47s** (from `2026-05-07T18:34:51Z` to `2026-05-08T05:15:38Z`).

Pre-gap chain behavior appears stable:
- Block interval in the 100 blocks before the incident is exactly **5 seconds** (min/avg/max = 5s).
- The parent hash linkage around the gap is continuous (no obvious reorg at the handover).

These signals are more consistent with a **liveness halt** than with a fork event.

## Incident Window

- Last block before gap: [`25979111`](https://explorer.qgov.io/block/25979111)
- First block after gap: [`25979112`](https://explorer.qgov.io/block/25979112)
- Gap duration: `38447s` (~10h 40m 47s)

## Data Sources

- Block details from Blockscout API (`/api/v2/blocks/<height>`)
- Block lists from Blockscout API (`/api/v2/blocks?block_number=...&items_count=...`)
- Mainnet explorer pages for anchor blocks:
  - [Block 25979111](https://explorer.qgov.io/block/25979111)
  - [Block 25979112](https://explorer.qgov.io/block/25979112)

## Observations

### 1) Last 10 block producers before the incident

Blocks `25979111` down to `25979102` were produced by 10 distinct validators:

1. `0x3BC7a69cb97bacA92520b581AdB9B747d5a97a3b`
2. `0xc38fA1d5e19a5dB2628E1323c30c784C923F3d1C`
3. `0x207157E38B380AF3890D88E78F851dfFb15419AA`
4. `0x004bAd8901038cC780F331de3B597Fc4a81f4357`
5. `0x748F7FE62bC5CFd8A6a9600161D098f6c08def1f`
6. `0xA4a4607B05e337871e4014CD60F15433FBCff384`
7. `0x101Bf5CefC653A663311368d68cB79FAFB76A25A`
8. `0x318F0fb4aAc8945232323f57721D930715D5c4D3`
9. `0xCa8109cc86371964dAd125Ed3915b45403c35e61`
10. `0x1FC4b7c3cdaf7C7391dB70c170Bf6f9Fad3AE25F`

### 2) Miner histogram over 100 blocks before incident

Window: `25979110` down to `25979011`

- 14 unique producers observed.
- Distribution is balanced; no immediate single-producer spike.

### 3) Miner histogram over 1000 blocks before incident

Window: `25979110` down to `25978111`

- 14 unique producers observed.
- Producers present in the 1000-block histogram but **not in the final-10 pre-gap window**:
  - `0x45b268C5Ca2b35cc740fe67BAF46f1b6ed3a2627` (72 blocks)
  - `0xb9b2ab8c4BE6deC0e1eD397DDe7DA149b0e75afC` (72 blocks)
  - `0x9D60e04FAf20bF9C34b131D27b4F279da4367327` (70 blocks)
  - `0xD4cF1Bb665aFf0Ee7F4849d1b7190385F1384394` (70 blocks)

## Consensus-Oriented Interpretation (Preliminary)

Given the stated rule context (Clique-like recent-signers restriction), the final pre-gap sequence suggests:

- The 10 most recent signers were ineligible for immediate resealing.
- The next eligible cohort did not produce for an extended period.

Potential causes to validate:
- Partial or broad validator offline event
- Network partition / peer connectivity collapse
- Consensus timeout escalation without quorum recovery
- Operational incident (coordinated restart, infra outage, clock drift, key/service issue)

## Recommended Next Investigation Steps

1. **Validator liveness logs** for all top-set validators covering the full gap window.
2. **Consensus round telemetry** (proposal/prevote/precommit failures, timeout progression).
3. **P2P connectivity evidence** (peer count history, disconnect spikes, regional reachability).
4. **Voting power and active-set changes** immediately before/during the gap.
5. **Host-level health checks** (NTP drift, disk IO stalls, resource saturation, process restarts).
6. **Change audit** (deployments/config/key management/firewall changes in preceding hours).

## Caveats

- Validator ranking page data was client-rendered and not fully extractable via simple HTTP fetch in this run.
- This report is based on observable chain data and does not yet include node-internal logs.
- This document is intentionally preliminary and should be updated with root-cause evidence.
