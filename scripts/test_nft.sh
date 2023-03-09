export VST_HOME="$HOME/vesta-test"
vestad tx vm store ./examples/nft.js --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm instantiate nft 0 "10,1uvst" --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm execute nft mint "" --from danny --home=$VST_HOME -y --gas 1000000
sleep 6

vestad tx vm store ./examples/nft_marketplace.js --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm instantiate market 1 "" --from danny --home=$VST_HOME -y
sleep 6
vestad tx vm execute market list "nft,0,10uvst" --from danny --home=$VST_HOME -y --gas 1000000
sleep 6
#vestad tx vm execute nft add_approval "0,vst10k05lmc88q5ft3lm00q30qkd9x6654h3gct5dt" --from danny --home=$VST_HOME -y --gas 1000000
#sleep 6
#vestad tx vm execute nft transfer "0,vst10k05lmc88q5ft3lm00q30qkd9x6654h3gct5dt" --from charlie --home=$VST_HOME --gas 1000000 -y