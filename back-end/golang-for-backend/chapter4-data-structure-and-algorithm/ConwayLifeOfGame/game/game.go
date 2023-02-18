package game

import (
	"image/color"
	"math/rand"
)

var (
	liveColor = color.RGBA{255, 0, 0, 255}
	deadColor = color.RGBA{255, 255, 255, 255}
)

type Pixels struct {
	// RGBA pixel : flatten (x,y,color channel) to 1d array
	Colors []uint8
	Width  int
}

func NewPixels(width, height int) *Pixels {
	return &Pixels{Width: width, Colors: make([]uint8, width*height*4)}
}

func (pixels *Pixels) DrawRect(x, y, width, height int, rgba color.Color) {
	for idx := 0; idx < width; idx++ {
		for idy := 0; idy < height; idy++ {
			pixels.SetColor(x+idx, y+idy, rgba)
		}
	}
}

func (pixels *Pixels) SetColor(x, y int, rgba color.Color) {
	r, g, b, a := rgba.RGBA()
	idx := (y*pixels.Width + x) * 4
	pixels.Colors[idx] = uint8(r)
	pixels.Colors[idx+1] = uint8(g)
	pixels.Colors[idx+2] = uint8(b)
	pixels.Colors[idx+3] = uint8(a)
}

type ConwayGameOfLife struct {
	board       [][]int
	pixels      *Pixels
	sizeOfPixel int // 1 픽셀의 크기
}

func NewConwayGameOfLife(width, height, size int) *ConwayGameOfLife {
	board := make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}

	pixels := NewPixels(width*size, height*size)
	return &ConwayGameOfLife{board: board, pixels: pixels, sizeOfPixel: size}
}

func (gol *ConwayGameOfLife) InitWithRandom() {
	for y := range gol.board {
		for x := range gol.board[y] {
			gol.board[y][x] = rand.Intn(2)
		}
	}
}

// InitWithMethuselah 안정화 상태까지 오래 걸리는 초기 조건
func (gol *ConwayGameOfLife) InitWithMethuselah() {
	midX, midY := len(gol.board[0])/2, len(gol.board)/2

	gol.board[midY][midX] = 1
	gol.board[midY-1][midX] = 1
	gol.board[midY+1][midX] = 1
	gol.board[midY][midX-1] = 1
	gol.board[midY-1][midX+1] = 1

	for y := range gol.board {
		for x := range gol.board[y] {
			if gol.board[y][x] == 1 {
				gol.pixels.DrawRect(x*gol.sizeOfPixel, y*gol.sizeOfPixel, gol.sizeOfPixel, gol.sizeOfPixel, liveColor)
			} else {
				gol.pixels.DrawRect(x*gol.sizeOfPixel, y*gol.sizeOfPixel, gol.sizeOfPixel, gol.sizeOfPixel, deadColor)
			}
		}
	}

}

func (gol *ConwayGameOfLife) PlayRound() {
	neighbors := gol.countNeighbors()

	for y := range gol.board {
		for x := range gol.board[y] {
			numNeighbor := neighbors[y][x]

			// 규칙에 맞게 현재 셀의 생존 여부 및 픽셀 정보 업데이트
			if gol.board[y][x] == 1 && (numNeighbor == 2 || numNeighbor == 3) {
				// 살아있고, 주변에 살아있는 셀이 2~3개인 경우 => 생존
				continue
			} else if gol.board[y][x] == 0 && numNeighbor == 3 {
				// 죽어있으나, 주변에 살아있는 셀이 3개인 경우 => 부활
				gol.board[y][x] = 1
				gol.pixels.DrawRect(x*gol.sizeOfPixel, y*gol.sizeOfPixel, gol.sizeOfPixel, gol.sizeOfPixel, liveColor)
			} else {
				// 나머지는 모두 사망
				gol.board[y][x] = 0
				gol.pixels.DrawRect(x*gol.sizeOfPixel, y*gol.sizeOfPixel, gol.sizeOfPixel, gol.sizeOfPixel, deadColor)
			}
		}
	}
}

func (gol *ConwayGameOfLife) countNeighbors() [][]int {
	neighbors := make([][]int, len(gol.board))
	for idx, row := range gol.board {
		neighbors[idx] = make([]int, len(row))
	}

	for x := 0; x < len(gol.board); x++ {
		for y := 0; y < len(gol.board[x]); y++ {
			// 인접한 셀 검색
			for i := -1; i < 2; i++ {
				nextX := x + i

				// Out of Bound
				if nextX < 0 || nextX >= len(gol.board) {
					continue
				}

				for j := -1; j < 2; j++ {
					// Current Cell
					if i == 0 && j == 0 {
						continue
					}

					nextY := y + j

					// Out of Bound
					if nextY < 0 || nextY >= len(gol.board[x]) {
						continue
					}

					neighbors[x][y] += gol.board[nextX][nextY]
				}
			}
		}
	}

	return neighbors
}

func (gol *ConwayGameOfLife) GetPixels() *Pixels {
	return gol.pixels
}
