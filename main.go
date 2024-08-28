package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	var ok bool
	var err error
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: g-grep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}
	if strings.Contains(pattern, "\\1") {
		ok, err = matchLine(strings.TrimSpace(string(line)), pattern)
	} else {
		ok, err = matchLine(strings.TrimSpace(string(line)), pattern)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}
	os.Exit(0)
}

func matchLine(line string, pattern string) (bool, error) {
	for i := 0; i <= len(line); i++ {
		if matchPattern(line, pattern, i) {
			return true, nil
		}
	}
	return false, nil
}

func matchPattern(line string, pattern string, pos int) bool {
	n := len(pattern)
	m := len(line)
	j := pos

	for i := 0; i < n; i++ {
		if j >= m {
			// Handle optional characters at the end of the pattern
			if i+1 < n && pattern[i+1] == '?' {
				i++
				continue
			}
			return false
		}

		if pattern[i] == '^' {
			sizeOfPattern := len(pattern[1:])
			sizeOfLine := len(line[:sizeOfPattern])
			if sizeOfLine == sizeOfPattern {
				if strings.EqualFold(line[:sizeOfPattern], pattern[1:]) {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		} else if strings.Contains(pattern, "\\1") {
			// Single Backreference regex
			input := replaceReference(pattern, line)
			if input == line {
				return true
			} else {
				return false
			}
		} else if pattern[i] == '.' {
			if line[:j] == pattern[:i] && line[j+1:] == pattern[i+1:] {
				return true
			}
			return false
		} else if string(pattern[i]) == "(" && string(pattern[len(pattern)-1]) == ")" {
			input := pattern[i+1 : len(pattern)-1]
			parts := strings.Split(input, "|")
			if contains(parts, line) {
				return true
			} else {
				fmt.Println("HER")
				return false
			}
		} else if pattern[len(pattern)-1] == '$' {
			sizeOfPattern := len(pattern[:len(pattern)-1])

			if strings.EqualFold(pattern[:len(pattern)-1], line[len(line)-(sizeOfPattern):]) {
				return true
			} else {
				return false
			}
		} else if pattern[i] == '[' && i+1 < n && pattern[i+1] == '^' {
			endPos := strings.Index(pattern[i:], "]")
			matchAnyPattern := pattern[i+1 : endPos]
			if strings.Contains(matchAnyPattern, string(line[j])) {
				return false
			}
			i = endPos
		} else if pattern[i] == '[' {
			endPos := strings.Index(pattern[i:], "]")
			matchAnyPattern := pattern[i+1 : endPos]
			if !strings.Contains(matchAnyPattern, string(line[j])) {
				return false
			}
			i = endPos
		} else if pattern[i] == '\\' && i+1 < n {
			// Handle escape sequences like \d and \w
			if pattern[i+1] == 'd' && !unicode.IsDigit(rune(line[j])) {
				return false
			} else if pattern[i+1] == 'w' && !(unicode.IsLetter(rune(line[j])) || unicode.IsDigit(rune(line[j])) || line[j] == '_') {
				return false
			} else {
				i++
				j++
			}
		} else if i+1 < n && pattern[i+1] == '?' {
			// Handle the ? quantifier
			if line[j] == pattern[i] {
				j++
			}
			// Skip the '?' in the pattern
			i++
		} else if i+1 < n && pattern[i+1] == '+' {
			// Handle the + quantifier
			if line[j] != pattern[i] {
				return false
			}
			// Move forward in the line as long as the characters match
			for j < m && line[j] == pattern[i] {
				j++
			}
			// Skip the '+' in the pattern and continue with the rest of the pattern
			i++
		} else if line[j] != pattern[i] {
			return false
		} else {
			j++
		}
	}

	// Ensure we've processed all characters in the input line
	return true
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}

func replaceReference(input string, line string) string {
	// Step 1: Extract the content inside parentheses
	var content string
	var startIndex, endIndex int
	found := false

	for i, char := range input {
		if char == '(' {
			startIndex = i + 1
		} else if char == ')' {
			endIndex = i
			content = input[startIndex:endIndex]
			found = true
			break
		}
	}

	if !found {
		return input // No parentheses found, return original string
	}

	// For pattern case
	if string(content[0]) != "\\" {
		// here
		result := strings.Replace(input, "\\1", content, -1)
		finalResult := strings.Replace(result, fmt.Sprintf("(%s)", content), content, -1)
		return finalResult
	} else if string(content[0]) == "\\" {

		originalString := strings.Split(line, " ")[0]
		fmt.Println(originalString)
		result := strings.Replace(input, "\\1", originalString, -1)
		finalResult := strings.Replace(result, fmt.Sprintf("(%s)", content), originalString, -1)
		return finalResult
	} else {
		return input
	}
}
