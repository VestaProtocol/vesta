
function proxyMint(proxyContract) {
    STD.fetch(proxyContract, "mint", "POST")
}

CONTRACT.functions.mint = proxyMint
