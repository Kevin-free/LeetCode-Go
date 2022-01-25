package _130_surrounded_regions

func solve(board [][]byte) {
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		if board[i][0] == 0 {
			board[i][0] = 2
		}
		if board[i][len(board[0])-1] == 0 {
			board[i][len(board[0])-1] = 2
		}
	}
	for i := 0; i < col; i++ {
		if board[0][i] == 0 {
			board[0][i] = 2
		}
		if board[len(board)-1][i] == 0 {
			board[len(board)-1][i] = 2
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i-1][j] != 2 && board[i][j-1] != 2 && board[i][j+1] != 2 && board[i+1][j] != 2 {
				board[i][j] = 1
			}
		}
	}

	for i := 0; i < row; i++ {
		if board[i][0] == 2 {
			board[i][0] = 0
		}
		if board[i][len(board[0])-1] == 2 {
			board[i][len(board[0])-1] = 0
		}
	}
	for i := 0; i < col; i++ {
		if board[0][i] == 2 {
			board[0][i] = 0
		}
		if board[len(board)-1][i] == 2 {
			board[len(board)-1][i] = 0
		}
	}

}
