package spell

import "fmt"

// котелок))
type Pot struct {

	// матрица
	M [25][25]string
	// Отображение решетки
	ViewS string
	// Отображение точки
	ViewO string
	// Размерность
	D int
}

func NewPot() Pot {

	var m [25][25]string

	d := 25
	o := " . "
	s := " # "

	for i := 0; i < d; i++ {

		for j := 0; j < d; j++ {
			m[i][j] = o

		}
	}

	return Pot{M: m, ViewS: s, ViewO: o, D: d}
}

func (p *Pot) Print() {

	for i := 0; i < p.D; i++ {
		for j := 0; j < p.D; j++ {
			fmt.Print(p.M[i][j])
		}

		fmt.Println("\n")
	}
}

// 1 spell 01
func (p *Pot) S01() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i <= j {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// 2 spell 02
func (p *Pot) S02() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i == j {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// 3 spell 03
func (p *Pot) S03() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i == j {
				p.M[i][p.D-j-1] = p.ViewS
			}

		}
	}
}

// 4 spell 04
func (p *Pot) S04() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i < 5 || j < 5 {
				p.M[i][j] = p.ViewS
			}

			if i < p.D-5 {
				if i < j {
					p.M[i+5][p.D-j] = p.ViewS
				}
			}
		}
	}
}

// 5 spell 05
func (p *Pot) S05() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if j/2 == i {

				p.M[i][j] = p.ViewS

			}

		}
	}
}

// spell 06
func (p *Pot) S06() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i < 10 {
				p.M[i][j] = p.ViewS
			}
			if j < 10 {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// 6 spell 07
func (p *Pot) S07() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i > 15 && j > 15 {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// 1 spell 08
func (p *Pot) S08() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i < 1 || j < 1 {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// 2 spell 09
func (p *Pot) S09() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i <= j-3 || j <= i-11 {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// 3 spell 10
func (p *Pot) S10() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if j/2-1 < i && j > i {

				p.M[i][j] = p.ViewS

			}
		}
	}
}

// 4 spell 11
func (p *Pot) S11() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i == 1 || i == 23 || j == 1 || j == 23 {

				p.M[i][j] = p.ViewS

			}
		}
	}
}

// 5 spell 13
func (p *Pot) S13() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if j >= i-3 && i >= j-3 {
				p.M[p.D-i-1][j] = p.ViewS
			}

		}
	}
}

// 6 spell 18
func (p *Pot) S18() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i < 2 && j > 0 || j < 2 && i > 0 {

				p.M[i][j] = p.ViewS

			}
		}
	}
}

// 7 spell 19
func (p *Pot) S19() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i == 0 || i == 24 || j == 0 || j == 24 {

				p.M[i][j] = p.ViewS

			}
		}
	}
}

// 8 spell 20
func (p *Pot) S20() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i < 2 && j > 0 || j < 2 && i > 0 {

				p.M[i][j] = p.ViewS

			}
		}
	}
}

// 9 spell 23
func (p *Pot) S23() {

	for i := 0; i < p.D; i = i + 3 {

		for j := 0; j < p.D; j = j + 2 {

			p.M[i][j] = p.ViewS

		}
	}
}

// 10 spell 24
func (p *Pot) S24() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i == j {
				p.M[i][p.D-j-1] = p.ViewS
				p.M[i][j] = p.ViewS
			}

		}
	}
}
