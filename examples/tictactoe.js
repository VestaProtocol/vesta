function makeBoard() {
    return [
        [0,0,0],
        [0,0,0],
        [0,0,0]
    ]
}

function t(token) {
    let tokens = [" ", "X", "O"]

    return tokens[token]
}

function printBoard(board) {
    let tb = board[0]
    let mb = board[1]
    let bb = board[2]

    let divider = "-------\n"

    let topRow = "|" + t(tb[0]) + "|" + t(tb[1]) + "|" + t(tb[2]) + "|\n"
    let middleRow = "|" + t(mb[0]) + "|" + t(mb[1]) + "|" + t(mb[2]) + "|\n"
    let bottomRow = "|" + t(bb[0]) + "|" + t(bb[1]) + "|" + t(bb[2]) + "|\n"

    return divider + topRow + divider + middleRow + divider + bottomRow + divider
}

function play(board, player, x, y) {
    if (x > 2 || x < 0 || y > 2 || y < 0 || player > 1 || player < 0) {
        return false
    }

    if (board[y][x] !== 0) {
        return false
    }

    board[y][x] = player + 1
}

function checkWinner(board) {
    if (board[0][0] === board[0][1] && board[0][0] === board[0][2]) {
        return board[0][0]
    }

    if (board[1][0] === board[1][1] && board[2][0] === board[1][2]) {
        return board[1][0]
    }

    if (board[2][0] === board[2][1] && board[2][0] === board[2][2]) {
        return board[2][0]
    }

    if (board[0][0] === board[1][0] && board[0][0] === board[2][0]) {
        return board[0][0]
    }

    if (board[0][1] === board[1][1] && board[0][1] === board[2][1]) {
        return board[0][1]
    }

    if (board[0][2] === board[1][2] && board[0][2] === board[2][2]) {
        return board[0][2]
    }

    if (board[0][0] === board[1][1] && board[0][0] === board[2][2]) {
        return board[0][0]
    }

    if (board[2][0] === board[1][1] && board[2][0] === board[0][2]) {
        return board[2][0]
    }

    return 0
}

function debug() {
    let b = makeBoard()

    play(b, 1, 0 ,0)
    play(b, 0, 1 ,0)
    play(b, 1, 1 ,1)

    console.log(printBoard(b))
    console.log(checkWinner(b))

    play(b, 0, 2 ,0)
    play(b, 1, 2 ,2)

    console.log(printBoard(b))
    console.log(checkWinner(b))

}

debug()