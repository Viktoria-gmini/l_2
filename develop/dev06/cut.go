package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "выбрать поля")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	selectedFields := strings.Split(*fields, ",")
	selectedMap := make(map[int]bool)
	for _, field := range selectedFields {
		if val, err := strconv.Atoi(field); err == nil {
			selectedMap[val-1] = true
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		fields := strings.Split(line, *delimiter)
		output := []string{}
		for i, field := range fields {
			if selectedMap[i] {
				output = append(output, field)
			}
		}
		fmt.Println(strings.Join(output, *delimiter))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
	}
}
