package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	argms := os.Args[1:]
	if len(argms) != 2 {
		fmt.Println("Usage: program sample.txt result.txt")
		os.Exit(1)
	}
	data, err := os.ReadFile(argms[0])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	text := string(data)
	re := regexp.MustCompile(`\s+([.,!?:;])`)
	text = re.ReplaceAllString(text, "$1")

	detachPunct := regexp.MustCompile(`\)([^\s)])`)
	text = detachPunct.ReplaceAllString(text, ") $1")

	words := strings.Fields(text)

	fmt.Println(words)
	fmt.Println("-----------------------------------------------------------------")
	for i := 0; i < len(words); i++ {
		if words[i] == "a" || words[i] == "A" {
			//( words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'i' || words[i+1][0] == 'o' || words[i+1][0] == 'u' || words[i+1][0] == 'A' || words[i+1][0] == 'E' || words[i+1][0] == 'I' || words[i+1][0] == 'O' || words[i+1][0] == 'U' || words[i+1][0] == 'h' || words[i+1][0] == 'H')
			if len(words[i+1]) > 0 && strings.ContainsRune("aeiouAEIOUhH", rune(words[i+1][0])) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
		//////////////////////////uuuuuuupppppppp//////////////////
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""
		}
		if words[i] == "(up," {
			numStr := words[i+1][:len(words[i+1])-1]
			stepBack, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error up converting number:", err)
				os.Exit(1)
			}
			for j := 1; j <= stepBack && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
				words[i] = ""
				words[i+1] = ""
			}
		}
		//////////////////////////loooooooowwwww//////////////////
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = ""
		}
		if words[i] == "(low," {
			numStr := words[i+1][:len(words[i+1])-1]
			stepBack, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error low converting number:", err)
				os.Exit(1)
			}
			for j := 1; j <= stepBack && i-j >= 0; j++ {
				words[i-j] = strings.ToLower(words[i-j])
				words[i] = ""
				words[i+1] = ""
			}
		}
		//////////////////////////caaaappppppppp//////////////////
		if words[i] == "(cap)" {
			if len(words[i-1]) > 0 {
				words[i-1] = strings.ToUpper(words[i-1][:1]) + words[i-1][1:]
			}
			words[i] = ""
		}
		if words[i] == "(cap," {
			numStr := words[i+1][:len(words[i+1])-1]
			stepBack, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error cap converting number:", err)
				os.Exit(1)
			}
			for j := 1; j <= stepBack && i-j >= 0; j++ {
				if len(words[i-j]) > 0 {
					words[i-j] = strings.ToUpper(words[i-j][:1]) + words[i-j][1:]
				}
				words[i] = ""
				words[i+1] = ""
			}
		}
		if words[i] == "(hex)" {
			if len(words[i-1]) > 0 {
				num, err := strconv.ParseInt(words[i-1], 16, 64)
				if err != nil {
					fmt.Println("Error hex converting number:", err)
					os.Exit(1)
				}
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words[i] = ""
		}
		if words[i] == "(bin)" {
			if len(words[i-1]) > 0 {
				num, err := strconv.ParseInt(words[i-1], 2, 64)
				if err != nil {
					fmt.Println("Error bin converting number:", err)
					os.Exit(1)
				}
				words[i-1] = strconv.FormatInt(num, 10)
			}
			words[i] = ""
		}
	}
	var results []string
	for _, w := range words {
		if w != "" {
			results = append(results, w)
		}
	}
	final := strings.Join(results, " ")
	final = re.ReplaceAllString(final, "$1")

	punctAfter := regexp.MustCompile(`([.,!?:;])([^\s.,!?:;])`)
	final = punctAfter.ReplaceAllString(final, "$1 $2")

	quotePair := regexp.MustCompile(`'\s*(.*?)\s*'`)
	final = quotePair.ReplaceAllString(final, "'$1'")
	err = os.WriteFile(argms[1], []byte(final), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}
