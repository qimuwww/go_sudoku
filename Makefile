sudoku: $(wildcard *.go)
	go build -x -o sudoku

clean:
	rm -f sudoku