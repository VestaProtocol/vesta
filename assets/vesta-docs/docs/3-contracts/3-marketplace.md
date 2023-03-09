---
sidebar_position: 3
---
# NFT Marketplace
In this overview, we will be building both a better NFT contract as well as a marketplace to facilitate the listing and 
sale of these NFTs. The best part about this marketplace is that it will work for any NFT deployed within a specification
of contracts, but we'll get to that later.

:::note

Please do not use any of these contracts in a production environment, they are meant only as education tools and are 
for sure riddled with security vulnerabilities and exploits. You have been warned.

:::
## The NFT Standard
For this example, we will be using [this NFT contract](https://github.com/VestaProtocol/vesta/blob/master/examples/nft.js).
### Data Structures
In this contract, we will have a counter that increments the IDs of each minted NFT as well as the NFTs themselves. We 
store NFTs as a JSON object, this will allow us to save as much or as little information in each token as we want.
```js
function getCount() {
    let s = STD.read(COUNTER)
    if (s == null) {
        return 0
    }
    return parseInt(s)
}

function incrementCount() {
    let count = getCount()
    STD.write(COUNTER, count + 1)
}

function Token(id, owner, data) {
    return {
        id: id,
        owner: owner,
        data: data,
        approvals: [],
    }
}
```
### Minting
To be able to mint a token, it looks a lot like our last NFT example, we check if there are any left, and if the user has
the tokens to pay for the mint, if so, we mint a new token.
```js
CONTRACT.functions.mint = function() {
    let currentCount = getCount()
    if (currentCount >= getMaxTokens()) {
        STD.panic("can't mint anymore tokens")
    }

    let token = Token(currentCount, CTX.sender, "nft data")

    let cost = getPrice()
    let ok = STD.bank.sendTokens(CONTRACT.address, cost)
    if (!ok) {
        STD.panic("not enough balance of " + cost)
    }

    SaveToken(token)
    incrementCount()
}
```
### Transfers
Here we deviate from the previous NFT contract, in this case we don't only check if the user is the owner before
authorizing the transfer. We now also check the list of authorizations to see if the user has been authorized by the 
owner to act on their behalf.
```js
CONTRACT.functions.transfer = function(id, receiver) {
    let token = LoadToken(id)

    if (token.owner !== CTX.sender) {
        let found = false
        for (const apr of token.approvals) {
            if (apr === CTX.sender) {
                found = true
                break
            }
        }
        if (!found) {
            STD.panic("cannot transfer token that isn't yours!")
        }
    }
    token.approvals = []
    token.owner = receiver
    SaveToken(token)
    return true
}
```
### Approvals
And finally, we need to add a way for a user to add someone to the list of approvals.
```js
CONTRACT.functions.add_approval = function(id, approval) {
    let token = LoadToken(id)

    if (token.owner !== CTX.sender) {
        STD.panic("cannot transfer token that isn't yours!")
    }

    token.approvals.push(approval)
    SaveToken(token)

    return true
}
```
## The Marketplace
Now that we have an NFT with an approval list, we can build a marketplace dedicated to the buying and selling of these 
NFTs. Because every NFT created with this contract will have a `transfer` and an `add_approval` function, we can make
the marketplace generic and work with any NFT contract.

For this example, we will be using [this Marketplace contract](https://github.com/VestaProtocol/vesta/blob/master/examples/nft_marketplace.js).

### Listings
Analyzing what is needed to hold the information about a listing, we can see that we create an ID from the other data 
fields and some hashing. Focusing more on the data itself, we have a `creator` to hold who is selling the NFT, a 
`contract` field to keep track of which NFT contract we are creating the listing for, a `token` field to keep track of 
which token ID from the contract we are selling. Lastly, we have a price field to keep track of the asking price for the
sale.
```js
function Listing(creator, contract, id, price) {
    let listingId = STD.crypto.sha256(creator + contract + id + price)
    return {
        id: listingId,
        creator: creator,
        contract: contract,
        token: id,
        price: price,
    }
}
```
To create a listing, we use the newly configured `add_approval` function to give our marketplace contract the ability to 
transfer an NFT that the seller owns. To do so we call the `fetch` function from Roma's standard library with a type of
`FORWARD` to indicate we are transacting on the user's behalf. We create the listing and add an approval to the token, 
and voilà! A listing is born.
```js
CONTRACT.functions.list = function(contract, token, price) {
    let listing = Listing(CTX.sender, contract, parseInt(token), price)
    SaveListing(listing)
    let res = STD.fetch(contract, "add_approval", STD.FORWARD, token, CONTRACT.address)
    if (res !== "true") {
        STD.panic("failed to add approval for token " + token + " on " + contract + "(" + CONTRACT.address + ")\nFull Log: " + JSON.stringify(res))
    }
}
```
### Buying
Now that we have the ability to sell NFTs, we need to match that with the ability to buy them. To do so, we first attempt
to transfer the token specified in the listing to the user. If we are unable to do so it is because the marketplace is 
not approved. Then we send the price of the token to the seller, if all goes well, we have swapped an NFT for a fee all
without ever acting as an escrow account.
```js
CONTRACT.functions.buy = function(id) {
    let listing = LoadListing(id)

    let res = STD.fetch(listing.contract, "transfer", STD.POST, listing.token, CTX.sender)
    if (res !== "true") {
        STD.panic("failed to transfer token " + listing.token + " on " + listing.contract + "(" + CONTRACT.address + ")\nFull Log: " + JSON.stringify(res))
    }

    let ok = STD.bank.sendTokens(CONTRACT.address, listing.price)
    if (!ok) {
        STD.panic("not enough balance of " + listing.price)
    }

    ok = STD.bank.withdrawTokens(listing.creator, listing.price)
    if (!ok) {
        STD.panic("not enough balance of " + listing.price)
    }

    RemoveListing(listing.id)
}
```