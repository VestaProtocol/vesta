
function mint() {
    var count = read(SENDER)
    if (count == null) {
        save(SENDER, "10")
        return
    }

    var numcount = parseInt(count) + 10

    save(SENDER, numcount.toString())
}

contractFunctions = {
    mint: mint
}