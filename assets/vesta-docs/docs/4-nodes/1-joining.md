---
sidebar_position: 1
---
# Creating Validator


:::tip

This guide assumes you are using the same machine as the full node.

Perform the following steps as your `vesta` user.

:::

## Creating A Wallet

We need to create a wallet and set the keyring password.

```sh
vestad keys add WALLET_NAME --keyring-backend os
```

This wallet is used to claim rewards, commission and to vote as your validator.

You will see a similar output once created.

```
- address: vst1xr2j5l0d9nmj9w5av3ll58jcg98k2c8uz0jpml
  name: WALLET_NAME
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AuGnw1Eoe7gK/h3UTI4G5RQL+vvVZ7IQV63jkyZDNrAc"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

cruise embark render title certain start twin describe ankle joy topple math code typical under lunch easily duck humor rural skill laugh skirt trend

```

Be sure to back up the seed phrase of your validator wallet.  It's also recommended to keep an offline copy along with your key files.  Remember, your key files cannot be restored and ***must*** be backed up.  See the installation page for instructions.

You should also back up your keyring files.

Change `WALLET_NAME` to the name of your wallet.
```sh
mkdir ~/keyring_backup
cp ~/.vesta/WALLET_NAME.info ~/keyring_backup
cp ~/.vesta/keyhash ~/keyring_backup
```

## Setting Up

### Configure Gas Prices

As a validator, you'll need to set a minimum gas price like so:
```sh
GAS="0.02uvst"
sed -i.bak -e "s/^minimum-gas-prices *=.*/minimum-gas-prices = \"$GAS\"/" $HOME/.vesta/config/app.toml
```

### Create Your Validator

Before continuing, please note that `commission-max-change` and `commission-max-rate` cannot be changed once you set them.  Your `commission-rate` may be changed once per day.

There are a few things you will need to alter in this command.  `amount` needs to be changed to what you are starting your self bond as.  `from` needs to be the name of your wallet you created earlier.  The `moniker`, `details`, `identity`, `website`, and `security-contact` should all be filled with the appropiate information.
```sh
vestad tx staking create-validator \
    --amount 1000000uvst \
    --commission-max-change-rate 0.10 \
    --commission-max-rate 0.2 \
    --commission-rate 0.1 \
    --from WALLET_NAME \
    --min-self-delegation 1 \
    --moniker "YOUR_MONIKER" \
    --details="YOUR DETAILS" \
    --identity "PGP IDENTITY" \
    --website="https://example.com" \
    --security-contact="your-email@email.com" \
    --pubkey $(vestad tendermint show-validator) \
    --chain-id vesta-1 \
    --gas-prices 0.02uvst
```
