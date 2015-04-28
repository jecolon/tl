package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	ud "github.com/fiam/gounidecode/unidecode"
	flag "github.com/ogier/pflag"
)

type config struct {
	infile string
	input  io.Reader
	output io.Writer
}

var infile = flag.StringP("file", "f", "stdin", "Input file. Defaults to stdin")

func main() {
	flag.Parse()
	conf := &config{
		infile: *infile,
		output: os.Stdout,
	}
	process(conf)
}

func process(c *config) {
	scanner := bufio.NewScanner(os.Stdin)

	if c.infile != "stdin" {
		if c.infile == "test" {
			scanner = bufio.NewScanner(c.input)
		} else {
			f, err := os.Open(c.infile)
			check(err)
			defer f.Close()
			scanner = bufio.NewScanner(f)
		}
	}

	for scanner.Scan() {
		fmt.Fprintf(c.output, "%s\n", ud.Unidecode(scanner.Text()))
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
