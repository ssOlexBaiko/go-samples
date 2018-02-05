package main

import (
	"fmt"
	"regexp"
	"flag"
	"bytes"
)

var (
	expr = flag.String("expr", `foo.?`, "set regular expression")
	text  = flag.String("text", "seafood fool", "set text for parsing")
)

func main() {
	flag.Parse()
	// When creating constants with regular expressions you can use the MustCompile variation of Compile.
	// A plain Compile won’t work for constants because it has 2 return values.
	// MustCompile is like Compile but <PANICs> if the expression cannot be parsed
	re, err := regexp.Compile(*expr)
	if err != nil {
		fmt.Println(err)
	}

	boolMatchString, err := regexp.MatchString(*expr, *text)
	fmt.Println("package MatchString: ",boolMatchString)
		fmt.Println("compiled MatchString: ", re.MatchString(*text))
	fmt.Println("FindAllString: ", re.FindAllString(*text, -1))
	fmt.Println("FindAllStringSubmatch: ", re.FindAllStringSubmatch(*text, -1))
	fmt.Println("FindAllStringSubmatchIndex: ", re.FindAllStringSubmatchIndex(*text, -1))
	fmt.Println("FindString: ", re.FindString(*text))
	fmt.Println("ReplaceAllLiteralString: ", re.ReplaceAllLiteralString(*text, `$\o`)) // The replacement repl is substituted directly, without using Expand.
	fmt.Println("ReplaceAllString: ", re.ReplaceAllString(*text, `$1`)) // $ signs are interpreted as in Expand, so for instance $1 represents the text of the first submatch.
	fmt.Println("Split: ", re.Split(*text, -1))
	fmt.Println("ReplaceAllFunc: ", re.ReplaceAllFunc([]byte(*text), bytes.ToUpper))
}
