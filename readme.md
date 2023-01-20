# Vesta Network ⚶
![banner](./assets/banner.png)
**Vesta** is a smart contract platform built with the Cosmos-SDK. Vesta allows developers to deploy and execute smart contracts written in Javascript on the blockchain. The contracts are stored in plain text and interpreted on runtime allowing for code auditability and composability. Unlike other smart contract platforms, vesta allows developers to assign names to smart contract instances and import functions from other contracts by name rather than address. View the [litepaper here](./assets/paper/litepaper.pdf).
## Install
```
go install ./...
```

## Smart Contracts
For examples of smart contracts built with vesta, see [examples](./examples/).

### Injections

```javascript

STD // An object holding the standard library functions & the injected libraries.

CTX // An object holding the context for the current execution including the sender of the message.

CONTRACT // An object holding both the information about the contract as well as a slot for the exported functions and queries.

```

### Deploying
```sh
vestad tx vm store {path_to_contract_file} --from {key}

vestad q vm list-contracts #note the code number of your contract

vestad tx vm instantiate {name} {code_num} {args} --from {key}
```

### Interacting
Execute function
```sh
vestad tx vm execute {name} {function} {args} --from {key}
```
Query function
```sh
vestad q vm query {name} {function} {args}
```


***
Powered by [Jackal Labs](https://www.jackallabs.io/)
