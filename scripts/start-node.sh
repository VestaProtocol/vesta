export VST_HOME="$HOME/vesta-test"

rm -rf $VST_HOME
export CHAIN="pompeii-1"
export ALIAS="marston"
export MONIKER="vstnode"


vestad init $MONIKER --home=$VST_HOME --chain-id=$CHAIN
vestad config chain-id $CHAIN --home=$VST_HOME
vestad config keyring-backend test --home=$VST_HOME

sed -i.bak -e 's/chain-id = ""/chain-id = "pompeii-1"/' $VST_HOME/config/client.toml

echo "video pluck level diagram maximum grant make there clog tray enrich book hawk confirm spot you book vendor ensure theory sure jewel sort basket" | vestad keys add $ALIAS --keyring-backend=test --recover --home=$VST_HOME
echo "flock stereo dignity lawsuit mouse page faith exact mountain clinic hazard parent arrest face couch asset jump feed benefit upper hair scrap loud spirit" | vestad keys add charlie --keyring-backend=test --recover --home=$VST_HOME
echo "brief enhance flee chest rabbit matter chaos clever lady enable luggage arrange hint quarter change float embark canoe chalk husband legal dignity music web" | vestad keys add danny --keyring-backend=test --recover --home=$VST_HOME

vestad add-genesis-account $(vestad keys show -a $ALIAS --keyring-backend=test --home=$VST_HOME) 500000000uvst --home=$VST_HOME
vestad add-genesis-account $(vestad keys show -a charlie --keyring-backend=test --home=$VST_HOME) 500000000uvst --home=$VST_HOME
vestad add-genesis-account $(vestad keys show -a danny --keyring-backend=test --home=$VST_HOME) 500000000uvst --home=$VST_HOME
vestad add-genesis-account $1 500000000uvst --home=$VST_HOME

vestad gentx $ALIAS 200000000uvst \
--chain-id=$CHAIN \
--moniker="$MONIKER" \
--commission-max-change-rate=0.01 \
--commission-max-rate=0.20 \
--commission-rate=0.05 \
--fees=2500uvst \
--from=$ALIAS \
--keyring-backend=test \
--home=$VST_HOME

vestad collect-gentxs --home=$VST_HOME

sed -i.bak -e 's/stake/uvst/' $VST_HOME/config/genesis.json
sed -i.bak -e 's/^minimum-gas-prices =""/minimum-gas-prices = \"0.0025uvst\"/' $VST_HOME/config/app.toml
sed -i.bak -e 's/enable = false/enable=true/' $VST_HOME/config/app.toml
sed -i.bak -e 's/enable=false/enable=true/' $VST_HOME/config/app.toml
sed -i.bak -e 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/' $VST_HOME/config/app.toml
sed -i.bak -e 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/' $VST_HOME/config/config.toml
sed -i.bak -e 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/' $VST_HOME/config/config.toml
sed -i.bak -e 's/laddr = "tcp:\/\/127.0.0.1:26656"/laddr = "tcp:\/\/0.0.0.0:26656"/' $VST_HOME/config/config.toml
sed -i.bak -e 's/chain-id = ""/chain-id = "pompeii-1"/' $VST_HOME/config/client.toml

vestad start --home=$VST_HOME --log_level info