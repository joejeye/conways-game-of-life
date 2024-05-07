package cells

func (B *Board) Equal(C *Board) bool {
	if B.NumRows != C.NumRows || B.NumCols != C.NumCols {
		return false
	}

	for i := 0; i < B.NumRows; i++ {
		for j := 0; j < B.NumCols; j++ {
			if B.Cells[i][j].Alive != C.Cells[i][j].Alive {
				return false
			}
		}
	}

	return true
}
