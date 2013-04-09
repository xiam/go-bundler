/*
  Copyright (c) 2012-2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"menteslibres.net/gosexy/cli"
	"os"
)

func init() {
	cli.Register("encode", cli.Entry{
		Description: "Converts a file into a []byte sequence.",
		Usage:       "encode /path/tofile",
		Command:     &initCommand{},
	})
}

type initCommand struct {
}

func (self *initCommand) Execute() error {

	var input string

	if flag.NArg() > 1 {
		input = flag.Arg(1)
	} else {
		return cli.Usage("encode")
	}

	stat, err := os.Stat(input)

	if err != nil {
		return err
	}

	if stat == nil {
		return fmt.Errorf("Could not open file.")
	} else {
		if stat.IsDir() == true {
			return fmt.Errorf("Cannot open file: %s, it's a directory.", input)
		}
	}

	fp, err := os.Open(input)

	if err != nil {
		return err
	}

	defer fp.Close()

	buf := make([]byte, 12)

	for {
		count, err := fp.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		for i := 0; i < count; i++ {
			if i > 0 {
				fmt.Printf(` `)
			}
			fmt.Printf(`0x%02x,`, buf[i])
		}
		fmt.Printf("\n")
	}

	return nil
}
