package main

import (
	"bytes"
	"testing"
)

func TestTransliterate(t *testing.T) {
	s := `-thí$- is áéíóúñ :/ bad~  .txt`

	conf := &config{
		infile: "test",
		input:  bytes.NewReader([]byte(s)),
		output: bytes.NewBuffer(make([]byte, 0)),
	}

	process(conf)

	w := "-thi$- is aeioun :/ bad~  .txt\n"
	o := conf.output.(*bytes.Buffer).String()
	if o != w {
		t.Fatalf("got %s wanted %s", o, w)
	}
}

func TestInfile(t *testing.T) {
	conf := &config{
		infile: "data.txt",
		output: bytes.NewBuffer(make([]byte, 0)),
	}

	process(conf)

	w := "-thi$- is aeioun :/ bad~  .txt\n"
	o := conf.output.(*bytes.Buffer).String()
	if o != w {
		t.Fatalf("got %s wanted %s", o, w)
	}
}
