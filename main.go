package main

import (
	"encoding/binary"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"os"
)

var (
	filenames     = kingpin.Arg("file", "baldr file to decode").Strings()
	nullDelimited = kingpin.Flag("null", "delimit output with null").Short('0').Bool()
)

func readBaldr(reader io.Reader) ([]byte, error) {
	var length uint64
	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	data := make([]byte, length)
	_, err := io.ReadFull(reader, data)
	return data, err
}

func process(reader io.Reader, nullDelimited bool) error {
	for {
		data, err := readBaldr(reader)
		if data != nil {
			fmt.Print(string(data[:]))
			if nullDelimited {
				fmt.Print("\x00")
			} else {
				fmt.Println()
			}
		}
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func processFilesOrStdin(filenames []string, nullDelimited bool, process func(io.Reader, bool) error) error {
	if len(filenames) == 0 {
		return process(os.Stdin, nullDelimited)
	}

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		if err := process(file, nullDelimited); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	kingpin.Parse()

	if err := processFilesOrStdin(*filenames, *nullDelimited, process); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
