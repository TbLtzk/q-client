In order to connect to the QMainnet network:

Start node: 

`$ geth --datadir=<datadir> --syncmode=full`

In order to connect to the QTestnet network:

Start node:

`$ geth --testnet --datadir=<datadir> --syncmode=full` 

Optionally you can put list of peers for required network (for faster syncing) under `<datadir>/geth/static-nodes.json`
