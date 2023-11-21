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

## [1.3.1] - 2023-06
### Added
- Testnet trusted checkpoint

### Fixed
- L0 exlusion set issues
- Mainnet trusted checkpoint

## [1.3.2] - 2023-06
### Added
- Duplicate root/exclusion list requires a confirmation
- Fallback to the list of signers from the genesis block if all validators are banned
- Added *gasBuffer* flag. If set - estimated gas will be multiplied by the given value
- Quarantine mechanism for exclusion lists. If the exclusion list can cause a huge rewind of the blockchain - this list
  will be quarantined until you decide to apply it. See [more detailed description](governance/README.adoc)
- Quota for proposed exclusion lists. It implies that if some of root nodes propose root/exclusion lists too
  often - all lists above the quota will be ignored
- Implemented rules for EIP-3436 (fork choice rules)
- Trace filters

### Notes
Default quota values are set to 3 per Root node/24 hours. If you want to change it, please use the following flags:
*gov.proposalQuotaMax* (max new list from root node) and *gov.proposalQuotaTimeWindow* (in hours)

## [1.3.3] - 2023-08
### Added
- S3 backup mananger and command line flags to enable it.
- Additional governance unit tests

### Changed 
- Make public root node approvals endpoint in governance API

## [1.3.4] - 2023-09
### Fixed
- Co-signing transition blocks by root nodes

## [1.3.5] - 2023-10
### Fixed
- DB freezing and graceful shutdown 

### Changed
- Improve sync speed
- Update mainnet root node list

### Added
- Ability to get account aliases proxy address at given block height

## Unreleased
### Changed
- Transition block approval now fails after defined unsuccessful attempts in a row (default is 10)
- The transition block is approved only after the transition block is in a canonical state

### Added
- Max approval failures can be changed using *gov.approvalMaxFailures* flag
- Number of verified blocks after the transition block needed for the approval can be changed using *gov.transitionBlockVerifiedBlocks* flag
