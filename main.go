package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var softTarget float64
var hardTarget float64
var ignore []string

// get and parse env vars
func init() {
	t := os.Getenv("INPUT_HARD_TARGET")
	if t == "" {
		panic("Environment variable '$INPUT_HARD_TARGET' not found")
	}
	var err error
	hardTarget, err = strconv.ParseFloat(t, 64)
	if err != nil {
		panic(err)
	}
	t = os.Getenv("INPUT_SOFT_TARGET")
	if t == "" {
		panic("Environment variable '$INPUT_SOFT_TARGET' not found")
	}
	softTarget, err = strconv.ParseFloat(t, 64)
	if err != nil {
		panic(err)
	}
	t = os.Getenv("IGNORE")
	ignore = strings.Split(t, ",")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		row, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		file, line, fn, pct := getLineRefAndPercent(row)
		if file == "" {
			continue
		}

		// "total" is the last line and the percent is the average of all the function-level coverage
		if file == "total" {
			fmt.Printf("::set-output name=coverage::%.1f%%\n", pct)
			if pct < softTarget {
				fmt.Printf("::warning::Coverage %.1f%% below soft target %.1f%%\n", pct, softTarget)
			}
			if pct < hardTarget {
				fmt.Printf("::error::Coverage %.1f%% below hard target %.1f%%\n", pct, hardTarget)
				os.Exit(1)
			}
			continue
		}
		skip := false
		for _, e := range ignore {
			if a, _ := path.Match(e+"*", file); a {
				skip = true
			}
		}
		if skip == true {
			continue
		}
		// if the function is not covered by tests at all, write a warning
		if pct == 0 {
			fmt.Printf("::warning file=%s,line=%s::'%s' not covered by tests\n", file, line, fn)
		}
	}
}

// matches the output from "go tool cover -func=cover.out"
var rx = regexp.MustCompile("(.*?)(:\\d+)?:\\s+(.*?)\\s+(\\d?\\d?\\d.\\d)%")

func getLineRefAndPercent(row string) (file string, line string, fn string, pct float64) {
	m := rx.FindStringSubmatch(row)
	if len(m) == 0 {
		return "", "", "", 0
	}
	file = strings.ReplaceAll(m[1], "github.com/"+os.Getenv("GITHUB_REPOSITORY"), "")
	if len(m[2]) > 0 {
		line = m[2][1:]
	}
	fn = m[3]
	pct, err := strconv.ParseFloat(m[4], 64)
	if err != nil {
		panic(err)
	}
	return
}
