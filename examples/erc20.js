
function mint() {
    let minted = std.mint("20")
    if (!minted) {
        return "failed mint"
    }
    minted = std.withdrawTokens(SENDER, "20vesta-" + NAME)
    if (!minted) {
        return "failed withdraw"
    }
}

contractFunctions = {
    mint: mint,
}

contractQueries = {
    // balance: balance
}