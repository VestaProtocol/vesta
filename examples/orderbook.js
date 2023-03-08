function Order(tokenIn, tokenOut, creator) {
    let id = STD.crypto.sha256(tokenIn + tokenOut + creator)
    return {
        id: id,
        tokenIn: tokenIn,
        tokenOut: tokenOut,
        creator: creator,
    }
}

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

CONTRACT.queries.show = function(orderId) {
    let order = LoadOrder(orderId)
    return JSON.stringify(order)
}

CONTRACT.functions.create = function(tIn, tOut) {
    let order = Order(tIn, tOut, CTX.sender)

    let ok = STD.bank.sendTokens(CONTRACT.address, tOut)
    if (!ok) {
        STD.panic("not enough balance of " + cost)
    }

    SaveOrder(order)
}

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