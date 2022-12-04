package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func any_overlap(first string, second string) int {
	one_start := strings.Split(first, "-")[0]
	one_start_int, err := strconv.ParseInt(one_start, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	one_end := strings.Split(first, "-")[1]
	one_end_int, err := strconv.ParseInt(one_end, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	two_start := strings.Split(second, "-")[0]
	two_start_int, err := strconv.ParseInt(two_start, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	two_end := strings.Split(second, "-")[1]
	two_end_int, err := strconv.ParseInt(two_end, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	if one_start_int <= two_start_int && one_end_int >= two_start_int {
		return 1
	} else if two_start_int <= one_start_int && two_end_int >= one_start_int {
		return 1
	}

	return 0
}
func total_overlap(first string, second string) int {

	one_start := strings.Split(first, "-")[0]
	one_start_int, err := strconv.ParseInt(one_start, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	one_end := strings.Split(first, "-")[1]
	one_end_int, err := strconv.ParseInt(one_end, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	two_start := strings.Split(second, "-")[0]
	two_start_int, err := strconv.ParseInt(two_start, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	two_end := strings.Split(second, "-")[1]
	two_end_int, err := strconv.ParseInt(two_end, 10, 0)
	if err != nil {
		fmt.Println("uh oh")
	}

	if one_start_int <= two_start_int && one_end_int >= two_end_int {
		return 1
	} else if two_start_int <= one_start_int && two_end_int >= one_end_int {
		return 1
	}

	return 0
}

func solve_part_1(lines []string) int {
	var assignment_duplicates int

	for _, line := range lines {
		sections := strings.Split(line, ",")
		first := sections[0]
		second := sections[1]
		assignment_duplicates += total_overlap(first, second)

	}

	return assignment_duplicates
}

func solve_part_2(lines []string) int {
	var assignment_duplicates int

	for _, line := range lines {
		sections := strings.Split(line, ",")
		first := sections[0]
		second := sections[1]
		assignment_duplicates += any_overlap(first, second)

	}

	return assignment_duplicates
}

func main() {

	puzzle_input := get_input("input-4.txt")
	fmt.Println(solve_part_1(puzzle_input))
	fmt.Println(solve_part_2(puzzle_input))

}
