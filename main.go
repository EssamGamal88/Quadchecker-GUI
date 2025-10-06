package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	// Read all of stdin
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	input := strings.Join(lines, "\n")
	input = strings.TrimRight(input, "\n")

	width, height := getDimensions(input)
	if width == 0 || height == 0 {
		fmt.Println("Not a quad function")
		return
	}

	quadNames := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	var matches []string
	for _, qn := range quadNames {
		output, err := quadFromBinary(qn, width, height)
		if err == nil && output == input {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", qn, width, height))
		}
	}
	sort.Strings(matches)
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	fmt.Println(strings.Join(matches, " || "))
}
func getDimensions(input string) (int, int) {
	lines := strings.Split(strings.TrimRight(input, "\n"), "\n")
	if len(lines) == 0 || lines[0] == "" {
		return 0, 0
	}
	width := len(lines[0])
	for _, line := range lines {
		if len(line) != width {
			return 0, 0
		}
	}
	return width, len(lines)
}

func quadFromBinary(quadName string, x, y int) (string, error) {
	cmd := exec.Command("./"+quadName, fmt.Sprint(x), fmt.Sprint(y))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(output), "\n"), nil
}
