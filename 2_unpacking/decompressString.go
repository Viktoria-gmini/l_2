package main

import (
	"bytes"
	"unicode"
)

func DecompressString(s string) string {
	var buf bytes.Buffer
	var count int
	var rl rune

	for i, r := range s {
		if unicode.IsDigit(r) {
			count = count*10 + int(r-'0')
		} else {
			if rl != 0 {
				if count <= 1 {
					buf.WriteRune(rl)
				} else {
					for j := 0; j < count; j++ {
						buf.WriteRune(rl)
					}
					count = 0
				}
			}
			rl = r
		}
		if (i == len(s)-1) && (rl != 0) {
			if count == 0 {
				count = 1
			}
			for j := 0; j < count; j++ {
				buf.WriteRune(rl)
			}
		}
	}

	return buf.String()
}
