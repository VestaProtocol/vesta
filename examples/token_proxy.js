
function proxyMint(proxyContract) {
    STD.fetch(proxyContract, "mint", STD.post)
    return STD.bank.withdrawTokens(CTX.sender, "1000000vesta-" + proxyContract)
}

CONTRACT.functions.mint = proxyMint
