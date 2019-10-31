package main

import "errors"

func isMutant(dna []string) (bool, error) {
	N := len(dna)
	for x := 0; x < N; x++ {
		if len(dna[x]) != N {
			return false, errors.New("Secuencia no váida")
		}
	}
	count := 0
	positions := [2][3]int{}
	/*Direcciones:
	0: Horizontal
	1: Oblicua
	2: Vertical
	3: Contra-Oblicua
	4: Final-Vertical
	5: Final-Horizontal
	*/
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
					if count > 0 && i == positions[count-1][0] && (j == positions[count-1][1]+1 || j == positions[count-1][1]+2 || j == positions[count-1][1]+3) && positions[count-1][2] == 0 {
						goto oblicuos
					}
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = 0
					count++
					if count > 1 {
						return true, nil
					}
				}
			}
		oblicuos:
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
					if count > 0 && ((i == positions[count-1][0]+1 && j == positions[count-1][1]+1) || (i == positions[count-1][0]+2 && j == positions[count-1][1]+2) || (i == positions[count-1][0]+3 && j == positions[count-1][1]+3)) && positions[count-1][2] == 1 {
						goto verticales
					}
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = 1
					count++
					if count > 1 {
						return true, nil
					}
				}
			}
		verticales:
			if dna[i][j] == dna[i+3][j] { // Elementos pivotantes Verticales
				coins := true
				for k := (i + 1); k <= (i + 2); k++ {
					if dna[i][j] != dna[k][j] {
						coins = false
						break
					}
				}
				if coins {
					if count > 0 && j == positions[count-1][1] && (i == positions[count-1][0]+1 || i == positions[count-1][0]+2 || i == positions[count-1][0]+3) && positions[count-1][2] == 2 {
						goto semiOblicuos
					}
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = 2
					count++
					if count > 1 {
						return true, nil
					}
				}
			}
		semiOblicuos:
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
					if count > 0 && ((i == positions[count-1][0]+1 && j+3 == positions[count-1][1]-1) || (i == positions[count-1][0]+2 && j+3 == positions[count-1][1]-2) || (i == positions[count-1][0]+3 && j+3 == positions[count-1][1]-3)) && positions[count-1][2] == 3 {
						continue
					}
					positions[count][0] = i
					positions[count][1] = j + 3
					positions[count][2] = 3
					count++
					if count > 1 {
						return true, nil
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
					if count > 0 && j == positions[count-1][1] && (i == positions[count-1][0]+1 || i == positions[count-1][0]+2 || i == positions[count-1][0]+3) && positions[count-1][2] == 4 {
						continue
					}
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = 4
					count++
					if count > 1 {
						return true, nil
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
					if count > 0 && i == positions[count-1][0] && (j == positions[count-1][1]+1 || j == positions[count-1][1]+2 || j == positions[count-1][1]+3) && positions[count-1][2] == 5 {
						continue
					}
					positions[count][0] = i
					positions[count][1] = j
					positions[count][2] = 5
					count++
					if count > 1 {
						return true, nil
					}
				}
			}
		}
	}
	return false, nil
}
