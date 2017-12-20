package main

import (
	"io"
	"io/ioutil"
	"mime/quotedprintable"
	"os"

	"github.com/skratchdot/open-golang/open"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convert(input *os.File) {
	QP := quotedprintable.NewReader(input)

  output, err := ioutil.TempFile("", "quotedprintableviewer")
  check(err)
  newName := output.Name() + ".html"
  err = output.Close()
  check(err)
  err = os.Remove(output.Name())
  check(err)
  output, err = os.Create(newName)
  check(err)
	defer output.Close()

	_, err = io.Copy(output, QP)
	check(err)

	err = open.RunWith(output.Name(), "chrome")
	if err != nil {
		err = open.Run(output.Name())
		check(err)
	}
}

func main() {
  for _, inputName := range os.Args[1:] {
    input, err := os.Open(inputName)
    check(err)
    convert(input)
  }
}
