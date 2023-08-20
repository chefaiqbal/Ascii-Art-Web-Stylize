package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetAsciiLine(filename string, num int) string {
	file, e := os.Open(filename)
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	return line
}

func AsciiArt(input, filename string) string {

	banner := "banners/" + filename + ".txt"
	line := ""
	result := "\n"

	args := strings.Split(input, "\n")
	for _, word := range args {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				result += GetAsciiLine(banner, 1+int(letter-' ')*9+i)
			}
			line += "\n"
			result += line
			line = ""
		}
	}
	return result
}
