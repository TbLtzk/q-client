#! /bin/sh

# you can use this script to reset the blockchain data to the latest stakecraft snapshot
# This can be useful for faster catchup (instead of syncing from genesis block)
# or to rewind to a stable state when your own chain data is corrupt or your node is stuck otherwise.

# usage:
#    cp ./reset-to-snapshot.sh ./data
#    docker-compose run --rm --entrypoint /data/reset-to-snapshot.sh q-client


rm -rf /data/geth/chaindata
mkdir -p /data/geth/chaindata
cd /data/geth/chaindata
wget -O - https://snapshots.stakecraft.com/q-testnet-latest.tar | tar xf -
