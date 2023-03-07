---
sidebar_position: 3
---
# Joining Testnet

After installing `vestad`. You can join the testnet by following these steps:

```sh
vestad init <alias> --chain-id=<chain-id>
```

:::note

`chain-id` for testnet is currently `pompeii-1`.

:::

Then we want to replace our generated genesis file with the one used to start the network. We also need to set our peers and seeds.

```sh
wget -O ~/.vesta/config/genesis.json {COMING SOON}

export SEEDS="{COMING SOON}"
sed -i.bak -e "s/^seeds *=.*/seeds = \"$SEEDS\"/" ~/.vesta/config/config.toml
```


As a validator, you'll need to set a minimum gas price like so:
```sh
GAS="0.02uvst"
sed -i.bak -e "s/^minimum-gas-prices *=.*/minimum-gas-prices = \"$GAS\"/" $HOME/.vesta/config/app.toml
```