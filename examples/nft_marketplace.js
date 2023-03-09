
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

function SaveListing(listing) {
    STD.write(listing.id, JSON.stringify(listing))
}

function RemoveListing(id) {
    STD.delete(id)
}

function LoadListing(id) {
    let sListing = STD.read(id)
    return JSON.parse(sListing)
}

CONTRACT.functions.list = function(contract, token, price) {
    let listing = Listing(CTX.sender, contract, parseInt(token), price)
    SaveListing(listing)
    let res = STD.fetch(contract, "add_approval", STD.FORWARD, token, CONTRACT.address)
    if (res !== "true") {
        STD.panic("failed to add approval for token " + token + " on " + contract + "(" + CONTRACT.address + ")\nFull Log: " + JSON.stringify(res))
    }
}

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