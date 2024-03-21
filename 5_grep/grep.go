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
	// Флаги
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	// Получить поисковой запрос
	term := flag.Arg(0)

	// Валидировать запрос на предмет пустострочия
	if term == "" {
		fmt.Println("Please provide a search term")
		os.Exit(1)
	}

	// Открыть inputfile для чтения
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Инициализация переменной подсчёта подходящих строк
	lineCount := 0
	matched := false
	matches := []string{}

	// Сканер для построчного чтения unputfile
	scanner := bufio.NewScanner(file)
	savedLines := []string{}
	//здесь хранится количество последующих строк, которые должны быть
	//выведены
	temp := 0
	for scanner.Scan() {
		line := scanner.Text()
		savedLines = append(savedLines, line)
		lineCount++
		if temp > 0 {
			matches = append(matches, line)
			temp--
		}
		// Проверка вхождения поискового запроса в строку
		matchedLine := func() bool {
			if *ignoreCase {
				if *fixed {
					return strings.EqualFold(line, term)
				} else {
					return strings.Contains(strings.ToLower(line), strings.ToLower(term))
				}
			} else {
				if *fixed {
					return line == term
				} else {
					return strings.Contains(line, term)
				}
			}
		}()

		// Операции, продиктованные флагами
		switch {
		case matchedLine && !*invert:
			matched = true
			matches = append(matches, line)
		case !matchedLine && *invert:
			matched = true
			matches = append(matches, line)
		case matched && *before > 0:
			matches = append(matches, line)
			for i := lineCount - *before - 1; i < lineCount-1; i++ {
				if i >= 0 {
					str := strconv.Itoa(i) + ": " + savedLines[i]
					matches = append(matches, str)
				}
			}
		case matched && *context > 0:
			matches = append(matches, line)
			for i := lineCount - 1 - *context/2; i < lineCount-1; i++ {
				if i >= 0 {
					str := strconv.Itoa(i) + ": " + savedLines[i]
					matches = append(matches, str)
				}
			}
			temp = *context / 2
		case matched && *after > 0:
			matches = append(matches, line)
			temp = *after
		case *count:
			if matchedLine {
				matches = append(matches, line)
			}
		}

		// Печатание запрашиваемых строк
		if len(matches) > 0 {
			if *count {
				continue
			}

			// номер строки
			if *lineNum {
				fmt.Printf("%d: ", lineCount)
			}

			// полученные строки
			for _, match := range matches {
				fmt.Println(match)
			}

			// очистка совпадений
			matches = []string{}
		}

		// Обновление индикатора совпадения для следующей строки
		if matchedLine && !*invert {
			matched = true
		} else if !matchedLine && *invert {
			matched = true
		} else {
			matched = false
		}
	}

	// Напечатать количество строк если был нужный флаг
	if *count {
		fmt.Println(len(matches))
	}

}
