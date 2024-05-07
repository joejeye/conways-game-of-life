package cells

func (B *Board) Clone() *Board {
	cells := make([][]Cell, B.NumRows)
	for i := 0; i < B.NumRows; i++ {
		cells[i] = make([]Cell, B.NumCols)
		for j := 0; j < B.NumCols; j++ {
			cells[i][j] = Cell{B.Cells[i][j].Alive, i, j}
		}
	}
	return &Board{B.NumRows, B.NumCols, cells}
}
