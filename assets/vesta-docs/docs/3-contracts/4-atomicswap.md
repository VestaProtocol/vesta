---
sidebar_position: 4
---
# Building an Atomic Swap
## What is an Atomic Swap?
An atomic swap is a token swapping mechanism that allows users to create swap orders that can then be fulfilled by 
other users. For example, if I have 10 Atom tokens, and I want 15 Vesta tokens, I can create an order for 15000000uvst
and 10000000uatom. Then, if someone decides that want 10 Atom tokens and has 15 Vesta tokens, they can fulfil this order
and the Vesta tokens will be sent to me, and they'll get my Atom tokens.

## The Smart Contract
### Defining Data-Structures
We first need to think about what each order will be composed of. Let's think about this, each order will need a 
*token wanted* and a *token offered*, let's call these `tokenIn` and `tokenOut`. We will also need to keep track of who
made the order, let's call this `creator`. And finally, we'll need to keep track of each order with an identifier, let's
use the hash of all other details and call it `id`.

```js
function Order(tokenIn, tokenOut, creator) {
    let id = STD.crypto.sha256(tokenIn + tokenOut + creator)
    return {
        id: id,
        tokenIn: tokenIn,
        tokenOut: tokenOut,
        creator: creator,
    }
}
```

The code above will allow use to create an order object using the standard libraries `crypto` library.

### Data Permanence
We can't expect each order to happen within the same block as each-other, and as such we'll need a way to save the 
orders onto the chain. We can do this with a few helper functions:
```js
function SaveOrder(order) {
    STD.write(order.id, JSON.stringify(order))
}

function RemoveOrder(id) {
    STD.delete(id)
}

function LoadOrder(id) {
    let sOrder = STD.read(id)
    return JSON.parse(sOrder)
}
```
### Creating Orders
When creating orders, we need to allow users to specify the tokens wanted/offered as well as keep their offer in the 
contracts account as escrow. We then create the order and save it to the chain, nothing crazy here.
```js
CONTRACT.functions.create = function(tIn, tOut) {
    let order = Order(tIn, tOut, CTX.sender)

    let ok = STD.bank.sendTokens(CONTRACT.address, tOut)
    if (!ok) {
        STD.panic("not enough balance of " + cost)
    }

    SaveOrder(order)
}
```
### Fulfilling Orders
In order to fulfil orders, we ask the user for an order ID, then we make sure they have enough tokens to support the 
swap. We then take the tokens from them, send them to the order `creator` and send the tokens they gave the contract 
earlier to the fulfiller. Finally, we delete the order to ensure nobody can try to fulfil it a second time.
```js
CONTRACT.functions.fulfil = function(orderId) {
    let order = LoadOrder(orderId)

    let ok = STD.bank.sendTokens(CONTRACT.address, order.tokenIn)
    if (!ok) {
        STD.panic("not enough balance of " + order.tokenIn)
    }

    ok = STD.bank.withdrawTokens(order.creator, order.tokenIn)
    if (!ok) {
        STD.panic("not enough balance of " + order.tokenIn)
    }

    ok = STD.bank.withdrawTokens(CTX.sender, order.tokenOut)
    if (!ok) {
        STD.panic("not enough balance of " + order.tokenOut)
    }

    RemoveOrder(orderId)
}
```
### Querying Data
If we want users to be able to view all the info about an order before fulfilling it, we can create a query like this.
```js
CONTRACT.queries.show = function(orderId) {
    let order = LoadOrder(orderId)
    return JSON.stringify(order)
}
```
## Interacting With The Contract
Now that you've built an atomic swap, you want to be able to interact with it. To do this you'll need two accounts with different tokens on each of them. (See [Testing AtomSwap Script](https://github.com/VestaProtocol/vesta/blob/master/scripts/test_atomswap.sh)).
### Creating Swap
```shell
vestad tx vm store ./examples/orderbook.js --from {account}
vestad tx vm instantiate book {code} "" --from {account}
```
### Creating Order
```shell
vestad tx vm execute book create "{token_a},{token_b}" --from {account}
```
After executing this command you can find the ID by checking this commands output for your contract's storage.
```shell
vestad q tx list-romdata
```
### Fulfilling Order
```shell
vestad tx vm execute book fulfil "{id}" --from {account}
```
## Conclusion
And there you have it, a full atomic swap system build with JavaScript deployed on Vesta! Happy coding!