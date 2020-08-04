package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	sudoku "go-sudoku/sudoku/go-sudoku"
)

var (
	questionFIle string
	errData      error = errors.New("input data error")
)

func init() {
	flag.StringVar(&questionFIle, "f", "./cases/case_1.txt", "read a question from a file")
	flag.Parse()
}

func main() {
	question, err := readQuestion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sudoku question:")
	sudoku.PrintSudoku(&question)
	dst := sudoku.Sudoku(question)
	sudoku.PrintSudoku(&dst)
}

func readQuestion() ([6][6]int, error) {
	var (
		n int
		q [6][6]int
	)
	f, err := os.Open(questionFIle)
	if err != nil {
		return q, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		s := strings.Fields(line)
		if len(s) != 6 {
			return q, errData
		}
		for i := 0; i < 6; i++ {
			num, err := strconv.Atoi(s[i])
			if num < 0 || num > 6 || err != nil {
				return q, errData
			}
			q[n][i] = num
		}
		n++
		if n == 6 {
			break
		}
	}
	return q, nil
}
