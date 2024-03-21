package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortFile(inputFile, outputFile string, column int, sortByNumeric, reverse, unique bool) error {
	// Open input file
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	// Read lines from input file
	scanner := bufio.NewScanner(input)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Custom sort function based on column and numeric flag
	sortFunc := func(i, j int) bool {
		if sortByNumeric {
			numI, _ := strconv.Atoi(strings.Fields(lines[i])[column-1])
			numJ, _ := strconv.Atoi(strings.Fields(lines[j])[column-1])
			if reverse {
				return numI > numJ
			}
			return numI < numJ
		}

		if reverse {
			return lines[i] > lines[j]
		}
		return lines[i] < lines[j]
	}

	sort.Slice(lines, sortFunc)

	// Remove duplicates if -u flag is set
	if unique {
		lines = removeDuplicates(lines)
	}

	// Write sorted lines to output file
	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	for _, line := range lines {
		fmt.Fprintln(output, line)
	}

	return nil
}

func removeDuplicates(lines []string) []string {
	encountered := map[string]bool{}
	result := make([]string, 0)

	for _, line := range lines {
		if !encountered[line] {
			encountered[line] = true
			result = append(result, line)
		}
	}

	return result
}

func main() {
	inputFile := flag.String("input", "input.txt", "Input file path")
	outputFile := flag.String("output", "output.txt", "Output file path")
	column := flag.Int("k", 0, "Column to sort on (default 0 - sort by whole line)")
	sortByNumeric := flag.Bool("n", false, "Sort by numeric value")
	reverse := flag.Bool("r", false, "Sort in reverse order")
	unique := flag.Bool("u", false, "Remove duplicate lines")

	flag.Parse()

	err := sortFile(*inputFile, *outputFile, *column, *sortByNumeric, *reverse, *unique)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("column:", *column)
	fmt.Println("sortByNumeric:", sortByNumeric)
	fmt.Println("reverse:", *reverse)
	fmt.Println("unique:", *unique)
}
