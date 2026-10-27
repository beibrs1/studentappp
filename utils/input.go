package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func ReadString(prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func ReadInt(prompt string) int {
	for {
		value := ReadString(prompt)
		n, err := strconv.Atoi(value)
		if err == nil {
			return n
		}
		if value == "" {
			return 0
		}
		fmt.Println("Бүтін сан енгізіңіз.")
	}
}
