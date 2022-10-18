This directory contains genesis block config (`genesis.json`) and default list of peers 
for existing q blockchain networks.

In order to connect to the network:
1. Init node:  
`$ geth --datadir=<datadir> init genesis.json`

2. Put list of peers for required network under `<datadir>/geth/static-nodes.json`

3. Start node with `networkid` from `genesis.json` (`chainId`):  
`$ geth --datadir=<datadir> --networkid=<networkid> --syncmode=full` 