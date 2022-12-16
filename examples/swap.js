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

    let a = STD.NewDec(pool.amountA)
    let b = STD.NewDec(pool.amountB)

    let k = a.mul(b)

    if (tokenIn == pool.denomA) {

        let i = STD.NewDec(amountIn)
        let m = k.div(a.add(i)).floor()

        let refund = b.sub(m)

        let pulled = STD.bank.withdrawTokens(CTX.sender, refund.toInt() + pool.denomB)
        if (!pulled) {
            return false
        }

        pool.amountA = a.add(i).toString()
        pool.amountB = m.toString()
    }

    if (tokenIn == pool.denomB) {

        let i = STD.NewDec(amountIn)
        let m = k.div(b.add(i)).floor()

        let refund = a.sub(m)

        let pulled = STD.bank.withdrawTokens(CTX.sender, refund.toInt() + pool.denomA)
        if (!pulled) {
            return false
        }

        pool.amountB = b.add(i).toString()
        pool.amountA = m.toString()
    }

    STD.write(poolId, JSON.stringify(pool))

    return true
}

CONTRACT.functions.swap = Swap
CONTRACT.functions.create = CreatePool

