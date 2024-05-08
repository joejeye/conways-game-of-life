package cells

func (B *Board) TerminalVisualize() (visual string) {
	visual = ""
	for i := 0; i < B.NumRows; i++ {
		line := []rune{}
		for j := 0; j < B.NumCols; j++ {
			if B.Cells[i][j].Alive {
				line = append(line, '■', ' ')
			} else {
				line = append(line, ' ', ' ')
			}
		}
		line = append([]rune{'|', ' '}, line...)
		line = append(line, '|')
		visual += string(line) + "\n"
	}
	return
}
