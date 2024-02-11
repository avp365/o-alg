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

	for j := 0; j < p.D; j++ {

		for i := 0; i < p.D; i++ {

			fmt.Print(p.M[j][i])
		}

		fmt.Println("\n")
	}
}

// spell 01
func (p *Pot) S01() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i <= j {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// spell 02
func (p *Pot) S02() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i <= j {
				p.M[i][j] = p.ViewS
			}

		}
	}
}

// spell 03
func (p *Pot) S03() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i == j {
				p.M[i][p.D-j-1] = p.ViewS
			}

		}
	}
}

// spell 04
func (p *Pot) S04() {

	for i := 0; i < p.D; i++ {

		for j := 0; j < p.D; j++ {

			if i < 5 {
				p.M[i][j] = p.ViewS
			}
			if j < 5 {
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

// spell 05
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
