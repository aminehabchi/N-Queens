func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{[]string{"Q"}}
	}
	if n <= 3 {
		return nil
	}
	board := make([]string, n)
	s := "Q"
	for i := 0; i < n-1; i++ {
		s = s + "."
	}
	for i := 0; i < n; i++ {
		board[i] = s
		s = "." + s[0:n-1]
	}
	fmt.Println(board)
	arr := [][]string{}
	for i := 0; i < n; i++ {
		arr = backtrack(arr, board, []string{board[i]}, n,1)
	}
	return arr
}
func backtrack(arr [][]string, board []string, newBoard []string, n int,l int) [][]string {
	if n == l {
		if checkboard(newBoard, n) {
			a := make([]string, n)
			copy(a, newBoard)
			arr = append(arr, a)
		}
		return arr
	} else {
		if len(newBoard)>3 && !checkboard(newBoard, l) {
			return arr
		}
	}
	for i := 0; i < n; i++ {
		if !is(newBoard, board[i]) {
			newBoard = append(newBoard, board[i])
			arr = backtrack(arr, board, newBoard, n,l+1)
			newBoard = newBoard[0 : len(newBoard)-1]
		}
	}
	return arr
}
func is(arr []string, s string) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			return true
		}
	}
	return false
}
func checkboard(board []string, n int) bool {
	for i := 0; i < len(board); i++ {
		j := strings.Index(board[i], "Q")
		if !check(board, len(board), i, j) {
			return false
		}
	}
	return true
}
func check(board []string, n, i, j int) bool {
	for k := 1; k <= n; k++ {
		if i+k < n && j+k < n && board[i+k][j+k] == 'Q' {
			return false
		}
		if i-k >= 0 && j-k >= 0 && board[i-k][j-k] == 'Q' {
			return false
		}
		if i-k >= 0 && j+k < n && board[i-k][j+k] == 'Q' {
			return false
		}
		if i+k < n && j-k >= 0 && board[i+k][j-k] == 'Q' {
			return false
		}
	}
	return true
}

