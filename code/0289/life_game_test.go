package main

func gameOfLife(board [][]int) {
	newBoard := make([][]int, len(board))
	for k, v := range board {
		newBoard[k] = make([]int, len(v))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			num := 0
			//左上
			if i > 0 && j > 0 && board[i-1][j-1] > 0 {
				num++
			}
			//正上
			if i > 0 && board[i-1][j] > 0 {
				num++
			}
			//右上
			if i > 0 && j < len(board[i])-1 && board[i-1][j+1] > 0 {
				num++
			}
			//正右
			if j < len(board[i])-1 && board[i][j+1] > 0 {
				num++
			}
			//右下
			if i < len(board)-1 && j < len(board[i])-1 && board[i+1][j+1] > 0 {
				num++
			}
			//正下
			if i < len(board)-1 && board[i+1][j] > 0 {
				num++
			}
			//左下
			if i < len(board)-1 && j > 0 && board[i+1][j-1] > 0 {
				num++
			}
			//正左
			if j > 0 && board[i][j-1] > 0 {
				num++
			}
			newBoard[i][j] = board[i][j]
			switch {
			case num < 2:
				newBoard[i][j] = 0
			case num == 3 && board[i][j] == 0:
				newBoard[i][j] = 1
			case num > 3:
				newBoard[i][j] = 0
			}
		}
	}
	copy(board, newBoard)
}
