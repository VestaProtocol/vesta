function CreatePool(tokenA, tokenB, amountA, amountB) {
    let pool = {
        denomA: tokenA,
        denomB: tokenB,
        amountA: amountA,
        amountB: amountB
    }

    let sent = STD.bank.sendTokens(CONTRACT.address, amountA + tokenA)
    if (!sent) {
        return false
    }
    sent = STD.bank.sendTokens(CONTRACT.address, amountB + tokenB)
    if (!sent) {
        return false
    }

    let count = STD.read("counter")
    if (count) {
        count = parseInt(count) + 1
    } else {
        count = 0
    }

    STD.write("counter", count)

    STD.write(count, JSON.stringify(pool))

    return true
}

function Swap(poolId, tokenIn, amountIn) {
    let p = STD.read(poolId)
    if (p == null || p == undefined) {
        return false
    }

    let pool = JSON.parse(p)

    let k = Number(pool.amountA) * Number(pool.amountB)

    if (tokenIn == pool.denomA) {

        let m = Math.floor(k / (Number(pool.amountA) + Number(amountIn)))

        let refund = Number(pool.amountB) - m

        let pulled = STD.bank.withdrawTokens(CTX.sender, refund + pool.denomB)
        if (!pulled) {
            return false
        }

        pool.amountA = Number(pool.amountA) + Number(amountIn)
        pool.amountB = m
    }

    if (tokenIn == pool.denomB) {

        let m = Math.floor(k / (Number(pool.amountB) + Number(amountIn)))

        let refund = Number(pool.amountA) - m

        let pulled = STD.bank.withdrawTokens(CTX.sender, refund + pool.denomA)
        if (!pulled) {
            return false
        }

        pool.amountB = Number(pool.amountB) + Number(amountIn)
        pool.amountA = m
    }

    STD.write(poolId, JSON.stringify(pool))

    return true
}

CONTRACT.functions.swap = Swap
CONTRACT.functions.create = CreatePool

