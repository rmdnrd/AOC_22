package main

import (
	"bufio"
	"fmt"
	"os"
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
	var priorities int
	for _, line := range lines {
		solved := false
		str1 := line[0 : len(line)/2]
		str2 := line[len(line)/2:]
		// a-z = 97-122 (-96)
		// A-Z = 65-90 if x < 90 (-38)
		for _, item1 := range str1 {
			if solved {
				break
			} else {
				for _, item2 := range str2 {
					if solved {
						break
					} else {
						if item1 == item2 {
							priority := int(item1)
							if priority > 96 {
								priority -= 96
							} else {
								priority -= 38
							}
							priorities += priority
							solved = true
						}
					}
				}
			}
		}

	}

	return priorities
}

func solve_part_2(lines []string) int {
	var priorities int
	for i := 0; i < len(lines); i += 3 {
		solved := false
		str1 := lines[i]
		str2 := lines[i+1]
		str3 := lines[i+2]
		// a-z = 97-122 (-96)
		// A-Z = 65-90 if x < 90 (-38)
		for _, item1 := range str1 {
			if solved {
				break
			} else {
				for _, item2 := range str2 {
					if solved {
						break
					} else {
						if item1 == item2 {
							for _, item3 := range str3 {
								if item1 == item3 {
									priority := int(item1)
									if priority > 96 {
										priority -= 96
									} else {
										priority -= 38
									}
									priorities += priority
									solved = true
									break
								}
							}

						}
					}
				}
			}
		}

	}

	return priorities
}

func main() {

	puzzle_input := get_input("input.txt")
	fmt.Println(solve_part_1(puzzle_input))
	fmt.Println(solve_part_2(puzzle_input))

}
