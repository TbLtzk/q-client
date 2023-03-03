# Changelog

All notable changes to this project will be documented in this file.

## [1.3.0] - 2023-03

### Changed
This version contains the core updates - now it contains changes from Go-ethereum up to
[version 1.10.26](https://github.com/ethereum/go-ethereum/releases/tag/v1.10.26).
Please note that some command-line flags became deprecated.

### Added
- multi-architecture build. Since this version ARM builds are also available in the registry
- L0 console methods for more human-readable representation of the exclusion lists
- API extension for getting transition block signatures
- API extension for the transaction indexer purposes
- API extension for getting voting events from the blockchain
- constitution storage. The part of the L0 which allows you to store and request the required constitution files. See [more detailed description](governance/README.adoc)

### Fixed
- L0 timestamp issues
