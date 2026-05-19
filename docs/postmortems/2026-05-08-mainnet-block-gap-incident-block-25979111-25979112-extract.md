# Extract: block 25979111/25979112 and reorg signals

- Source: `docs/postmortems/2026-05-08-mainnet-block-gap-incident-evidence-v2202503172553322813.md`
- Matched lines: 15

```
[validator L256] INFO [05-07|18:34:47.641] Commit new sealing work                  number=25,979,111 sealhash=fe8d01..239383 txs=0    gas=0         fees=0             elapsed=5.964ms
[validator L258] INFO [05-07|18:34:52.496] Imported new chain segment               number=25,979,111 hash=475c25..9d7334 blocks=1       txs=0    mgas=0.000  elapsed=5.853ms   mgasps=0.000   triedirty=1.30MiB
[validator L259] INFO [05-07|18:34:52.503] Commit new sealing work                  number=25,979,112 sealhash=cd4e6f..212967 txs=0    gas=0         fees=0             elapsed=6.805ms
[validator L261] INFO [05-07|18:34:52.513] Imported new chain segment               number=25,979,111 hash=9cc9f3..64e350 blocks=1       txs=0    mgas=0.000  elapsed=17.078ms  mgasps=0.000   triedirty=1.30MiB
[validator L263] INFO [05-07|18:34:52.778] Imported new chain segment               number=25,979,111 hash=0cbbd3..35c9e5 blocks=1       txs=0    mgas=0.000  elapsed=17.372ms  mgasps=0.000   triedirty=1.30MiB
[validator L264] INFO [05-07|18:34:52.856] Imported new chain segment               number=25,979,111 hash=7b495a..b2148d blocks=1       txs=0    mgas=0.000  elapsed=13.557ms  mgasps=0.000   triedirty=1.30MiB
[validator L7462] INFO [05-08|05:15:42.150] Chain reorg detected                     number=25,979,110 hash=12426c..2639fd drop=1 dropfrom=475c25..9d7334 add=2 addfrom=328a21..d2bcdc
[validator L7463] INFO [05-08|05:15:42.151] Imported new chain segment               number=25,979,112 hash=328a21..d2bcdc blocks=1       txs=0    mgas=0.000  elapsed=13.883ms  mgasps=0.000   triedirty=1.30MiB
[validator L7551] INFO [05-08|05:17:38.792] Chain reorg detected                     number=25,979,134 hash=aec758..5f7da6 drop=1 dropfrom=a8b6b3..738d58 add=2 addfrom=e31326..d9de0a
[rootnode L7781] INFO [05-07|18:34:52.492] Imported new chain segment               number=25,979,111 hash=9cc9f3..64e350 blocks=1     txs=0   mgas=0.000 elapsed=10.031ms  mgasps=0.000   triedirty=881.39KiB
[rootnode L7782] INFO [05-07|18:34:52.514] Imported new chain segment               number=25,979,111 hash=475c25..9d7334 blocks=1     txs=0   mgas=0.000 elapsed=14.682ms  mgasps=0.000   triedirty=881.39KiB
[rootnode L7784] INFO [05-07|18:34:52.768] Imported new chain segment               number=25,979,111 hash=0cbbd3..35c9e5 blocks=1     txs=0   mgas=0.000 elapsed=15.547ms  mgasps=0.000   triedirty=881.39KiB
[rootnode L7785] INFO [05-07|18:34:52.873] Imported new chain segment               number=25,979,111 hash=7b495a..b2148d blocks=1     txs=0   mgas=0.000 elapsed=14.016ms  mgasps=0.000   triedirty=881.39KiB
[rootnode L14888] INFO [05-08|05:15:42.151] Imported new chain segment               number=25,979,112 hash=328a21..d2bcdc blocks=1     txs=0   mgas=0.000 elapsed=13.864ms  mgasps=0.000   triedirty=881.39KiB
[rootnode L14940] INFO [05-08|05:17:38.794] Chain reorg detected                     number=25,979,134 hash=aec758..5f7da6 drop=1 dropfrom=a8b6b3..738d58 add=2 addfrom=e31326..d9de0a
```
