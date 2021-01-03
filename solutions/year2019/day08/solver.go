package day08

import "fmt"

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

var WIDTH = 25
var HEIGHT = 6

func (s *Solver) SolveA(input string) string {
	return ""
}

func (s *Solver) SolveB(input string) string {
	layerSize := WIDTH * HEIGHT
	numLayers := len(input) / layerSize

	layers := make([]string, numLayers)
	for i := 0; i < numLayers; i++ {
		layers[i] = input[i*layerSize : (i+1)*layerSize]
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			idx := y*WIDTH + x
			for _, layer := range layers {
				pixel := layer[idx]
				if pixel == '1' {
					fmt.Print("X")
					break
				} else if pixel == '0' {
					fmt.Print(" ")
					break
				}
			}
		}
		fmt.Println("")
	}

	return ""
}
