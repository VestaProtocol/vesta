# Vesta Network ⚶
**Vesta** is a smart contract platform built with the Cosmos-SDK. Vesta allows developers to deploy and execute smart contracts written in Javascript on the blockchain. The contracts are stored in plain text and interpreted on runtime allowing for code auditability and composability. Unlike other smart contract platforms, vesta allows developers to assign names to smart contract instances and import functions from other contracts by name rather than address.
## Install
```
go install ./...
```

## Smart Contracts
For examples of smart contracts built with vesta, see [examples](./testdata/).

### Deploying
```sh
vestad tx vm store {path_to_contract_file} --from {key}

vestad q vm list-contracts #note the code number of your contract

vestad tx vm instantiate {name} {code_num} --from {key}
```

### Interacting
```sh
vestad tx vm execute {name} {function} {args} --from {key}
```