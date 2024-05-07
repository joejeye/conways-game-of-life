package cells

import (
	"math/rand"
)

type Cell struct {
	Alive          bool
	RowIdx, ColIdx int
}

type Board struct {
	NumRows, NumCols int
	Cells            [][]Cell
}

func NewBoard(numRows, numCols int) *Board {
	cells := make([][]Cell, numRows)
	for i := 0; i < numRows; i++ {
		cells[i] = make([]Cell, numCols)
		for j := 0; j < numCols; j++ {
			cells[i][j] = Cell{false, i, j}
		}
	}
	return &Board{numRows, numCols, cells}
}

func (B *Board) RandomInit(n int) {
	for i := 0; i < n; {
		rowIdx := rand.Intn(B.NumRows)
		colIdx := rand.Intn(B.NumCols)
		if B.Cells[rowIdx][colIdx].Alive {
			continue
		}
		B.Cells[rowIdx][colIdx].Alive = true
		i++
	}
}

type Offset struct {
	RowOffset, ColOffset int
}

var NeighborOffsets = []Offset{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func WrapIndex(idx, size int) int {
	for idx < 0 {
		idx += size
	}

	return idx % size
}

func (B *Board) CountLiveNeighbors(cell Cell) int {
	count := 0
	for _, offset := range NeighborOffsets {
		rowIdx := WrapIndex(cell.RowIdx+offset.RowOffset, B.NumRows)
		colIdx := WrapIndex(cell.ColIdx+offset.ColOffset, B.NumCols)
		if rowIdx < 0 || rowIdx >= B.NumRows || colIdx < 0 || colIdx >= B.NumCols {
			continue
		}
		if B.Cells[rowIdx][colIdx].Alive {
			count++
		}
	}
	return count
}

func (B *Board) Next() {
	nextBoard := NewBoard(B.NumRows, B.NumCols)
	for i := 0; i < B.NumRows; i++ {
		for j := 0; j < B.NumCols; j++ {
			func(ii, jj int) {
				count := B.CountLiveNeighbors(B.Cells[ii][jj])
				if B.Cells[ii][jj].Alive {
					nextBoard.Cells[ii][jj].Alive = count == 2 || count == 3
				} else {
					nextBoard.Cells[ii][jj].Alive = count == 3
				}
			}(i, j)
		}
	}
	B.Cells = nextBoard.Cells
}
