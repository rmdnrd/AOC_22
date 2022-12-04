package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func get_input(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var results_list []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line string
		line = scanner.Text()
		results_list = append(results_list, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return results_list
}

func solve_part_1(lines []string) int {
	var score int
	
	for _, line := range lines{
		
		throws := strings.Split(line," ")
		opp_throw := throws[0]
		own_throw := throws[1]

		if opp_throw == "A"{
			if own_throw == "X"{
				score += 4
			} else if own_throw == "Y"{
				score += 8
			} else if own_throw == "Z"{
				score += 3
			}
		}else if opp_throw == "B"{
			if own_throw == "X"{
				score += 1
			} else if own_throw == "Y"{
				score += 5
			} else if own_throw == "Z"{
				score += 9
			}
		}else if opp_throw == "C"{
			if own_throw == "X"{
				score += 7
			} else if own_throw == "Y"{
				score += 2
			} else if own_throw == "Z"{
				score += 6
			}
		}
	}

	return score

}

func solve_part_2(lines []string) int {
	var score int
	
	for _, line := range lines{
		
		throws := strings.Split(line," ")
		opp_throw := throws[0]
		own_throw := throws[1]

		if opp_throw == "A"{
			if own_throw == "X"{
				score += 3
			} else if own_throw == "Y"{
				score += 4
			} else if own_throw == "Z"{
				score += 8
			}
		}else if opp_throw == "B"{
			if own_throw == "X"{
				score += 1
			} else if own_throw == "Y"{
				score += 5
			} else if own_throw == "Z"{
				score += 9
			}
		}else if opp_throw == "C"{
			if own_throw == "X"{
				score += 2
			} else if own_throw == "Y"{
				score += 6
			} else if own_throw == "Z"{
				score += 7
			}
		}
	}
	
	return score

}
func main(){
	file_path := "input.txt"
	lines := get_input(file_path)
	
	fmt.Println((solve_part_1(lines)))
	fmt.Println((solve_part_2(lines)))

}