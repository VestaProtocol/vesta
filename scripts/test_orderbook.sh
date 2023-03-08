export VST_HOME="$HOME/vesta-test"

vestad tx vm store ./examples/token.js --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm instantiate a 0 "" --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm instantiate b 0 "" --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm execute a mint "" --from charlie --home=$VST_HOME -y
vestad tx vm execute b mint "" --from danny --home=$VST_HOME -y
sleep 6

vestad tx vm store ./examples/orderbook.js --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm instantiate book 1 "" --from danny --home=$VST_HOME -y
sleep 6

vestad tx vm execute book create "100vesta-a,100vesta-b" --from danny --home=$VST_HOME -y --gas=1000000
sleep 6
