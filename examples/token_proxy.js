
function proxyMint(proxyContract) {
    STD.fetch(proxyContract, "mint", "POST")
    var sent = STD.bank.withdrawTokens(CTX.sender, "1000000vesta-" + proxyContract)
    return sent
}

CONTRACT.functions.mint = proxyMint
