package main

func isMutant(dna []string) bool {
	N := len(dna)
	count := 0
	positions := [2][4]int{}
	direction := ""
	for i := 0; i <= (N - 1); i++ {
		for j := 0; j <= (N-4) && i <= (N-4); j++ {
			if dna[i][j] == dna[i][j+3] { // Elementos pivotantes Horizontal
				coins := true
				for k := (j + 1); k <= (j + 2); k++ {
					if dna[i][j] != dna[i][k] {
						coins = false
						break
					}
				}
				if coins {
					if count > 0 && i <= positions[count-1][2] && j <= positions[count-1][3] && i == positions[count-1][0] && direction == "H" {
						continue
					}
					direction = "H"
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = i
					positions[count][3] = j + 3
					count++
					if count > 1 {
						return true
					}
				}
			}
			if dna[i][j] == dna[i+3][j+3] { // Elementos pivotantes Oblicuos
				coins := true
				for k := (i + 1); k <= (i + 2); {
					for l := (j + 1); l <= (j + 2); l++ {
						if dna[i][j] != dna[k][l] {
							coins = false
							break
						}
						k++
					}
					if !coins {
						break
					}
				}
				if coins {
					if count > 0 && i <= positions[count-1][2] && j <= positions[count-1][3] && i != positions[count-1][0] && j != positions[count-1][1] && direction == "O" {
						continue
					}
					direction = "O"
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = i + 3
					positions[count][3] = j + 3
					count++
					if count > 1 {
						return true
					}
				}
			}
			if dna[i][j] == dna[i+3][j] { // Elementos pivotantes Verticales
				coins := true
				for k := (i + 1); k <= (i + 2); k++ {
					if dna[i][j] != dna[k][j] {
						coins = false
						break
					}
				}
				if coins {
					if count > 0 && i <= positions[count-1][2] && j <= positions[count-1][3] && j == positions[count-1][1] && direction == "V" {
						continue
					}
					direction = "V"
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = i + 3
					positions[count][3] = j
					count++
					if count > 1 {
						return true
					}
				}
			}
			if dna[i][j+3] == dna[i+3][j] { // Elementos semi-pivotantes Oblicuos
				coins := true
				for k := (i + 1); k <= (i + 2); {
					for l := (j + 2); l >= (j + 1); l-- {
						if dna[i][j+3] != dna[k][l] {
							coins = false
							break
						}
						k++
					}
					if !coins {
						break
					}
				}
				if coins {
					if count > 0 && i <= positions[count-1][2] && j <= positions[count-1][3] && i != positions[count-1][0] && j != positions[count-1][1] && direction == "%" {
						continue
					}
					direction = "%"
					positions[count][0] = i
					positions[count][1] = j + 3
					positions[count][2] = i + 3
					positions[count][3] = j
					count++
					if count > 1 {
						return true
					}
				}
			}
		}
		for j := (N - 4) + 1; j <= (N-1) && i <= (N-4); j++ { // Elementos finales Verticales
			if dna[i][j] == dna[i+3][j] {
				coins := true
				for k := (i + 1); k <= (i + 2); k++ {
					if dna[i][j] != dna[k][j] {
						coins = false
						break
					}
				}
				if coins {
					if count > 0 && i <= positions[count-1][2] && j <= positions[count-1][3] && j == positions[count-1][1] && direction == "v" {
						continue
					}
					direction = "v"
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = i + 3
					positions[count][3] = j
					count++
					if count > 1 {
						return true
					}
				}
			}
		}
		for j := 0; j <= (N-4) && i >= (N-4)+1; j++ { // Elementos finales Horizontales
			if dna[i][j] == dna[i][j+3] {
				coins := true
				for k := (j + 1); k <= (j + 2); k++ {
					if dna[i][j] != dna[i][k] {
						coins = false
						break
					}
				}
				if coins {
					if count > 0 && i <= positions[count-1][2] && j <= positions[count-1][3] && i == positions[count-1][0] && direction == "h" {
						continue
					}
					direction = "h"
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = i
					positions[count][3] = j + 3
					count++
					if count > 1 {
						return true
					}
				}
			}
		}
	}
	return false
}
