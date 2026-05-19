# Incident Log Triage Summary

- Source: `docs/postmortems/2026-05-08-mainnet-block-gap-incident-evidence-v2202503172553322813.md`
- Host: `v2202503172553322813.megasrv.de`
- Incident window (UTC): `2026-05-07T18:30:00Z` to `2026-05-08T05:20:00Z`

## Key Signal Counts

| Signal | Count | First seen (UTC) | Last seen (UTC) |
|---|---:|---|---|
| `imported_segment` | 240 | 2026-05-07T18:30:01.842Z | 2026-05-08T05:19:58.737Z |
| `sealed_block` | 6 | 2026-05-07T18:30:32.404Z | 2026-05-08T05:18:45.868Z |
| `signed_recently` | 63 | 2026-05-07T18:30:01.872Z | 2026-05-08T05:19:28.634Z |
| `exclusion_hash_mismatch` | 13199 | 2026-05-07T18:30:15.657Z | 2026-05-08T05:19:56.834Z |
| `quarantine_notice` | 1300 | 2026-05-07T18:30:32.782Z | 2026-05-08T05:19:32.908Z |
| `peer_count_zero` | 0 | - | - |
| `conn_issue` | 12 | 2026-05-07T18:39:15.477Z | 2026-05-07T18:39:15.477Z |

## Gap Boundary Markers

- Last imported-segment log at/before halt start: `2026-05-07T18:34:47.635000Z`
  - `INFO [05-07|18:34:47.635] Imported new chain segment               number=25,979,110 hash=12426c..2639fd blocks=1       txs=0    mgas=0.000  elapsed=5.447ms   mgasps=0.000   triedirty=1.30MiB`
- First imported-segment log at/after recovery: `2026-05-08T05:15:42.151000Z`
  - `INFO [05-08|05:15:42.151] Imported new chain segment               number=25,979,112 hash=328a21..d2bcdc blocks=1       txs=0    mgas=0.000  elapsed=13.883ms  mgasps=0.000   triedirty=1.30MiB`

## Representative Samples

- `signed_recently`: WARN [05-07|18:30:01.872] Block sealing failed                     err="signed recently, must wait for others"
- `exclusion_hash_mismatch`: WARN [05-07|18:30:31.898] Exclusion list hash mismatch             provided=0x37809e789038d75ab188f9e2112495b088e732282c8343741c5902540dd01507 calculated=0x61b7caed3930d652bf98942ab4dfaaa5a948736e87de11a013556e5895a09a06 timestamp=1,647,860,003
- `quarantine_notice`: WARN [05-07|18:30:32.782] You have exclusion lists in the quarantine. You can see them with command: gov.quarantinedExclusionLists()
- `conn_issue`: WARN [05-07|18:39:15.477] Failed to retrieve stats server message  err="read tcp 172.30.0.2:54362->172.67.151.198:443: read: connection reset by peer"

## Preliminary Triage Notes

- Validator/rootnode containers show `restart=0` over the period (no obvious container restart during incident).
- Repeated `Exclusion list hash mismatch` warnings are persistent before and through the gap window.
- Repeated `signed recently, must wait for others` appears as expected under clique-like sealing constraints.
- Correlate these warnings with governance state / exclusion-list propagation across other validators to determine if this contributed to liveness failure.
