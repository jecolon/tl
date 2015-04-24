package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"translit"

	flag "github.com/ogier/pflag"
)

var unspace = flag.BoolP("unspace", "s", false, "Replace whitespace")
var spacesub = flag.StringP("unspace-sub", "r", "_", "Whitespace replacement")
var spacecol = flag.BoolP("collapse", "c", false, "Collapse repeated whitespace")
var infile = flag.StringP("file", "f", "", "Input file. If none, stdin")

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	if *infile != "" {
		f, err := os.Open(*infile)
		check(err)
		defer f.Close()
		scanner = bufio.NewScanner(f)
	}

	for scanner.Scan() {
		s := scanner.Text()
		if *spacecol {
			s = translit.CollapseSpaces(s)
		}
		if *unspace {
			s = translit.Unspace(s, *spacesub)
		}

		fmt.Println(translit.Transliterate(s))
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
