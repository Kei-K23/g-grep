# `G-Grep` - A Custom grep -E Implementation in Go

## Overview

`G-Grep` is a custom implementation of the grep -E command in `Go`. It reads input from stdin and matches it against a specified pattern, returning success (exit code 0) if a match is found, and failure (exit code 1) otherwise. This program supports some basic regular expression features, including `^, $, ?, +, ., character classes ([]), and backreferences`.

## Features

`Pattern Matching`: Matches lines based on the provided pattern.
`Anchors`: Supports ^ (beginning of line) and $ (end of line).
`Character Classes`: Supports character classes with negation ([^]).
`Quantifiers`: Supports ? (zero or one), + (one or more), and . (any single character).
`Backreferences`: Supports backreferences like \1.
`Escape Sequences`: Supports \d (digits) and \w (word characters).

## Usage

```bash
echo <input_text> | ./g-grep -E <pattern>
```

### Examples

1. Match any line containing the word "Go":

```bash
echo "I love Go programming" | ./g-grep -E "Go"
```

2. Match lines that start with "Hello":

```bash
echo "Hello, World!" | ./g-grep -E "^Hello"
```

3. Match lines that end with "end":

```bash
echo "This is the end" | ./g-grep -E "end$"
```

4. Match lines with a digit:

```bash
echo "The year is 2024" | ./g-grep -E "\d"
```

5. Match lines with a word character:```

```bash
echo "This_is_a_word" | ./g-grep -E "\w"
```

6. Match lines with an optional character (zero or one occurrence):

```bash
echo "color" | ./g-grep -E "colou?r"
```

7. Match lines with one or more occurrences of a character:

```bash
echo "aaa" | ./g-grep -E "a+"
```

8. Match lines using a backreference:

```bash
echo "abcabc" | ./g-grep -E "(abc)\1"
```

### Exit Codes

- 0: A match was found.
- 1: No match was found.
- 2: An error occurred (e.g., incorrect usage or pattern syntax).

## Installation

Ensure you have Go installed on your system.

1. Clone this repository:

```bash
git clone https://github.com/yourusername/g-grep.git
```

2. Build the program:

```bash
cd g-grep
go build -o g-grep
```

## Contributing

Feel free to contribute to this project by opening issues, suggesting features, or submitting pull requests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

This project is inspired by the grep -E command and aims to provide a basic understanding of pattern matching and regular expressions in `Go`.
