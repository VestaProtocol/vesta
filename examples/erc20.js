
function mint() {
    var count = read(SENDER)
    if (count == null) {
        save(SENDER, "10")
        return
    }

    var numcount = parseInt(count) + 10

    save(SENDER, numcount.toString())
}

function balance(sender) {
    var count = read(sender)
    if (count == null) {
        return "0"
    }
    return count
}

contractFunctions = {
    mint: mint
}

contractQueries = {
    balance: balance
}