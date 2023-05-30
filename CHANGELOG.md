# Changelog
All notable changes to this project will be documented in this file.

## [1.2.4] - 2023-04
### Changed
- Updated initial root nodes for Mainnet

## [1.3.0] - 2023-05
### Changed
This version contains the core updates - now it contains changes from Go-ethereum up to
[version 1.10.26](https://github.com/ethereum/go-ethereum/releases/tag/v1.10.26).
Please note that some command-line flags became deprecated.

### Added
- Multi-architecture build. Since this version ARM builds are also available in the registry
- Sentry.io integration. If configured, q-client will send alerts about critical/fatal errors/set of blockchain-related error (bad blocks, etc.)
- Root list flexibility. Root nodes can propose root list without themselves, active root list is no longer compared with on-chain list.
- L0 console methods for more human-readable representation of the exclusion lists
- New function in gov API extension for extracting transition block signatures
- Indexer API extension for the transaction indexer purposes (moved some functionality from Clique)
- New functions in indexer API extension for getting voting events from the blockchain
- Constitution storage. The part of the L0 that allows you to store and request the required constitution files. See [more detailed description](governance/README.adoc)
- Updated initial root nodes for Mainnet
- Updated DNS settings for Mainnet/Tesnet.

### Fixed
- L0 timestamp issues
- Multiple fixes

### Notes
Pay attention, that `--whitelist` flag became deprecated, if such flag is set, you should consider removing it
