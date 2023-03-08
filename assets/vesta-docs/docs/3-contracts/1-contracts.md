---
sidebar_position: 1
---

# Building Smart Contracts

For this demonstration we will be building [721.js](https://github.com/VestaProtocol/vesta/blob/master/examples/721.js), a simplified NFT contract based on the ERC721 standard.

## Contract Structure

Smart contracts on Vesta are split into a few major parts, the first and most important being the `CONTRACT` object. This object is responsible for exposing functions and queries to both the user, and other contracts, as well as handling initialization.

For this NFT contract we want two main functions, the ability to Mint tokens, and send them to new owners. We also want to be able to view the contracts data in an organized way. We can handle the minting and transferring with two functions as so:
```js
CONTRACT.functions.transfer = function(token_id, to) {}

CONTRACT.functions.mint = function() {}
```
As you can see, we create two new functions as fields of the `functions` object on the `CONTRACT` object, this allows users/contracts to call these functions.

Because this is an NFT we want to store a list of tokens & respective owners, their metadata URL, the cost of minting a token, how many tokens are allowed to be minted/how many are already minted. We can do this with the `STD.read` & `STD.write` functions from the Roma Standard Library. All data saved in Vesta Smart Contracts is done through a Key-Value store. We will create a series of constant strings to be used as store keys.
```js
const KEY_URL = "url"
const KEY_NAME = "name"
const KEY_TOKEN_ID = "token_id"
const KEY_COST = "cost"
const KEY_MAX = "MAX"
const KEY_COUNTER = "COUNTER"

CONTRACT.functions.transfer = function(token_id, to) {}

CONTRACT.functions.mint = function() {}
```

Now that we have the keys mapped out, let's implement some logic to handle the minting of the tokens.
```js
CONTRACT.functions.mint = function() {
    let count = parseInt(STD.read(KEY_COUNTER))

    let max_tokens = parseInt(STD.read(KEY_MAX))

    if (max_tokens < count + 1) {
        STD.panic("too many minted already")
    }

    let cost = STD.read(KEY_COST)
    let ok = STD.bank.sendTokens(CONTRACT.address, cost)
    if (!ok) {
        STD.panic("not enough balance of " + cost)
    }

    STD.write(KEY_COUNTER, count + 1)

    STD.write(count, CTX.sender)
}
```
As you can see we are reading the counter value and parsing it into an integer value, this is because every value saved to the store is saved as a string, so numbers must be parsed to avoid potential implicit type conversions.

Once we verify that there are enough tokens left to mint, we attempt to take the price of the tokens from the user, this is done through the `STD.bank` object. This object is an injection created by [Roma's Bank Inject](https://github.com/VestaProtocol/roma/tree/master/vminjects/bank).

If the user can't send the tokens, the process will fail and the transaction will be reverted. However, if the user has the required balance, the new token ID is saved to the KV store with the users address as the value to indicate they are now the owner. Lastly, the counter is incremented.

We can then go ahead and do the same for the transfer function:
```js
CONTRACT.functions.transfer = function(token_id, to) {
    let count = parseInt(STD.read(KEY_COUNTER))

    if (parseInt(token_id) > count) {
        STD.panic("no token with that id exists")
    }

    let owner = STD.read(token_id)

    if (owner !== CTX.sender) {
        STD.panic("you do not own this token")
    }

    STD.write(parseInt(token_id), to)
}
```

This is great and all, but now how do we see any of the information? Sure we could check the raw KV store, but that would be awful. Instead, we can use the `queries` object from the `CONTRACT` object. Let's take a look at creating a query that returns the information about the contract.
```js
CONTRACT.queries.info = function() {
    let url = STD.read(KEY_URL)
    let name = STD.read(KEY_NAME)
    let tokenid = STD.read(KEY_TOKEN_ID)
    let max = STD.read(KEY_MAX)
    let cost = STD.read(KEY_COST)

    let count = STD.read(KEY_COUNTER)

    return JSON.stringify({
        url: url,
        name: name,
        token_id: tokenid,
        max: max,
        minted: count,
        price: cost,
    })
}
```
We can see that this query is reading the information form the keys we laid out earlier and presenting them in a JSON response.

We can also create a query to get similar information but for a single token.
```js
CONTRACT.queries.get = function(token_id) {
    let count = parseInt(STD.read(KEY_COUNTER))

    if (parseInt(token_id) >= count) {
        STD.panic("no token with that id exists")
    }

    let owner = STD.read(token_id)

    let url = STD.read(KEY_URL) + token_id

    return JSON.stringify({
        url: url,
        token_id: token_id,
        owner: owner,
    })
}
```
This time we read the token data from a token ID input, and return the owner or panic if the token can't be found.

Finally, you may be wondering how the URL field contains anything since we haven't written anything to it yet. For this we fall back to the initialization of the contract. Every contract has a field called `init` which contains a function that gets called when the contract is first instantiated.

```js
CONTRACT.init = function(name, tokenid, url, cost, max) {
    STD.write(KEY_URL, url)
    STD.write(KEY_NAME, name)
    STD.write(KEY_TOKEN_ID, tokenid)
    STD.write(KEY_COST, cost)
    STD.write(KEY_MAX, max)
    STD.write(KEY_COUNTER, "0")
}
```

Here you can see that we initialize the contract with fields send from the user upon initialization. This allows developers to instantiate multiple contracts with the same source code without needing to re-upload it.

And finally we can save this complete file as `721.js` and upload it to the chain like:
```shell
vestad tx vm upload 721.js --from {account} # this example we will consider this the first code uploaded, slot 0

vestad tx vm instantiate NFT {contract index} "NFT,nft,http://localhost/nft_metadata,10stake,20" --from {account}
```

This creates a contract called NFT with the ticker nft, a price of 10 Stake, and a max supply of 20.

You can then mint a token by running:
```shell
vestad tx vm execute NFT mint "" --from {account}
```

This should succeed given that you have 10 Stake on that account, you can check the token now with:
```shell
vestad q vm query NFT get 0
```
This should return the token information expressed in the `get` function we wrote earlier.

And that's it, you have now created a simple NFT contract on Vesta in Javascript!