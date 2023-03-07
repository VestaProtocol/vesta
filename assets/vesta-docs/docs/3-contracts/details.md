---
sidebar_position: 2
---
# Contract Details
## Injections

```javascript

STD // An object holding the standard library functions & the injected libraries.

CTX // An object holding the context for the current execution including the sender of the message.

CONTRACT // An object holding both the information about the contract as well as a slot for the exported functions and queries.

```

## Deploying
```sh
vestad tx vm store {path_to_contract_file} --from {key}

vestad q vm list-contracts #note the code number of your contract

vestad tx vm instantiate {name} {code_num} {args} --from {key}
```

## Interacting
Execute function
```sh
vestad tx vm execute {name} {function} {args} --from {key}
```
Query function
```sh
vestad q vm query {name} {function} {args}
```