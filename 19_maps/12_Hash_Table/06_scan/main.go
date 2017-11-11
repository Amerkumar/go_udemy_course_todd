package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const input = "Attitude is a choice. Happiness is a choice. Optimism is a choice. Kindness is a choice. Giving is a choice. Respect is a choice. Whatever choice you make makes you. Choose wisely."

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input :", err)
	}

}
