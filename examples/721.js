const KEY_URL = "url"
const KEY_NAME = "name"
const KEY_TOKEN_ID = "token_id"
const KEY_COST = "cost"
const KEY_MAX = "MAX"
const KEY_COUNTER = "COUNTER"

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

CONTRACT.init = function(name, tokenid, url, cost, max) {
    STD.write(KEY_URL, url)
    STD.write(KEY_NAME, name)
    STD.write(KEY_TOKEN_ID, tokenid)
    STD.write(KEY_COST, cost)
    STD.write(KEY_MAX, max)
    STD.write(KEY_COUNTER, "0")
}

