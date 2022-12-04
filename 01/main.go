package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func solve_part_1(lines []string) int{
	max_calories := 0
	elf_calories := 0
	for _, line := range lines{
		if len(line) == 0{
			if elf_calories > max_calories{
				max_calories = elf_calories
			}
			elf_calories = 0
		}

		calories, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			elf_calories += int(calories)
		}
	}
	return max_calories
	
}

func solve_part_2(lines []string) int{
	var top_three [3] int
	//one,two,three := 0,0,0

	var elf_calories int
	for _, line := range lines{
		if len(line) == 0{
			for i, value := range top_three{
				if elf_calories > value{
					if i == 2{
						top_three[i] = elf_calories
						break
					}else if i ==1{
						top_three[i+1] = top_three[i]
						top_three[i] = elf_calories
						break
					}else{
						top_three[i+2] = top_three[i+1]
						top_three[i+1] = top_three[i]
						top_three[i] = elf_calories
						break
					}
				}

			}
		
			elf_calories = 0
		}
		calories, err := strconv.ParseInt(line, 10, 64)
		if err == nil {
			elf_calories += int(calories)
		}
	}
	var total int
	for _,calories := range top_three{
		total += calories
	}
	return total
}
func main() {

	file_path := "input.txt"
	lines := get_input(file_path)
	fmt.Println(solve_part_1(lines))
	fmt.Println(solve_part_2(lines))

 }
 
