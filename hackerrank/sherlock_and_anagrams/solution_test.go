package sherlockandanagrams

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestSherlockAndAnagrams(t *testing.T) {
	paths, err := filepath.Glob(filepath.Join("testdata", "*.input"))
	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		_, filename := filepath.Split(path)
		testname := filename[:len(filename)-len(filepath.Ext(path))]
		outname := fmt.Sprintf("%s.output", testname)
		outpath := filepath.Join(filepath.Dir(path), outname)

		// Each path turns into a test: the test name is the filename without the
		// extension.
		t.Run(testname, func(t *testing.T) {
			input, err := os.Open(path)
			if err != nil {
				t.Fatal("error opening input file:", err)
			}

			output, err := os.Open(outpath)
			if err != nil {
				t.Fatal("error opening output file:", err)
			}

			inputReader := bufio.NewReaderSize(input, 16*1024*1024)
			outputReader := bufio.NewReaderSize(output, 16*1024*1024)

			qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(inputReader)), 10, 64)
			if err != nil {
				t.Fatal("error parsing int:", err)
			}
			q := int32(qTemp)

			for qItr := 0; qItr < int(q); qItr++ {
				s := readLine(inputReader)
				want, err := strconv.ParseInt(readLine(outputReader), 10, 64)
				if err != nil {
					t.Fatal("error parsing int:", err)
				}

				// >>> This is the actual code under test.
				result := sherlockAndAnagrams(s)
				// <<<

				if result != int32(want) {
					t.Errorf("got %v, want %v", result, want)
				}
			}
		})
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
