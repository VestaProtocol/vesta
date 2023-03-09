const COUNTER = "COUNTER"
const MAX_COUNT = "MAX_COUNT"
const PRICE = "PRICE"

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

function getMaxTokens() {
    let s = STD.read(MAX_COUNT)
    if (s == null) {
        return 0
    }
    return parseInt(s)
}

function getPrice() {
    return STD.read(PRICE)
}

function SaveToken(token) {
    STD.write(token.id, JSON.stringify(token))
}


function LoadToken(id) {
    let sToken = STD.read(id)
    return JSON.parse(sToken)
}

function Token(id, owner, data) {
    return {
        id: id,
        owner: owner,
        data: data,
        approvals: [],
    }
}

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

CONTRACT.functions.add_approval = function(id, approval) {
    let token = LoadToken(id)

    if (token.owner !== CTX.sender) {
        STD.panic("cannot transfer token that isn't yours!")
    }

    token.approvals.push(approval)
    SaveToken(token)

    return true
}

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

CONTRACT.queries.get = function(id){
    let t = LoadToken(parseInt(id))
    return JSON.stringify(t)
}

CONTRACT.init = function(maxMint, price) {
    STD.write(MAX_COUNT, parseInt(maxMint))
    STD.write(PRICE, price)
}
