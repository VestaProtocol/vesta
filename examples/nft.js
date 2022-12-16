
function save(address, id) {
    STD.write(id, address)
}

CONTRACT.functions.mint = function() {
    let count = STD.read("counter")
    if (count) {
        count = parseInt(count) + 1
    } else {
        count = 0
    }

    if (parseInt(STD.read("max")) <= count) {
        STD.panic("too many minted already")
    }

    let cost = STD.read("cost")
    let ok = STD.bank.sendTokens(CONTRACT.address, cost)
    if (!ok) {
        STD.panic("not enough balance of " + cost)
    }

    STD.write("counter", count)

    save(CTX.sender, count)
}

CONTRACT.queries.get = function(id) {
    let count = STD.read("counter")
    if (count) {
        count = parseInt(count) + 1
    } else {
        count = 0
    }
    if (parseInt(id) >= count) {
        return null
    }


    url = STD.read("url")
    owner = STD.read(id)

    return JSON.stringify({
        url: url + id,
        owner: owner,
        token_id: id,
    })
}

CONTRACT.queries.info = function() {
    url = STD.read("url")
    name = STD.read("name")
    tokenid = STD.read("token_id")
    max = STD.read("max")
    cost = STD.read("cost")
    return JSON.stringify({
        url: url,
        name: name,
        token_id: tokenid,
        max: max,
        cost: cost,
    })
}

CONTRACT.init = function(name, tokenid, url, cost, max) {
    STD.write("url", url)
    STD.write("name", name)
    STD.write("token_id", tokenid)
    STD.write("cost", cost)
    STD.write("max", max)
}