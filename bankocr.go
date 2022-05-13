package bankocr

import "strings"

// Parse returns a numeric string for the input.
func Parse(input string) string {
	lines := strings.Split(input, "\n")
	var digit_lines [][]string
	for i := 0; i < len(lines); i += 3 {
		digit_lines = append(digit_lines, []string{lines[i], lines[i+1], lines[i+2]})
	}
	// For each entry
	result := ""
	for _, digit_line := range digit_lines {
		for i := 0; i < 9; i++ {
			result += getDigitString(digit_line, i)
		}
		result += "\n"
	}
}

func getDigitString(input []string, idx int) string {
	result := ""
	for _, line := range input {
		for i := 0; i < 3; i++ {
			result += string(line[i+idx*3])
		}
		result += "\n"
	}
	return result
}
