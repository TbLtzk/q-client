In order to connect to the Q MainNet network:

Start node:

`$ geth --datadir=<datadir> --syncmode=full`

In order to connect to the Q TestNet network:

Start node:

`$ geth --testnet --datadir=<datadir> --syncmode=full`
or
`$ geth --fischer --datadir=<datadir> --syncmode=full`

Optionally you can put list of peers for required network (for faster syncing) under `<datadir>/geth/static-nodes.json`
