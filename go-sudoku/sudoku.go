// Author: wkj

package sudoku

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 当前填充点的横座标,纵座标
type currentCoordinate struct {
	abscissa int
	ordinate int
}

func Sudoku(src [6][6]int) (dst [6][6]int) {
	dst = src
	if b := tryFill(&dst); b {
		log.Println("solved,got an answer")
	} else {
		log.Println("sudoku filling error")
	}
	return
}

func PrintSudoku(s *[6][6]int) {
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && i != 0 {
			fmt.Printf("-------------\n")
		}
		for j := range s[i] {
			if j%3 == 0 && j != 0 {
				fmt.Printf(" |")
			}
			if j != 0 {
				fmt.Printf(" %d", s[i][j])
			} else {
				fmt.Printf("%d", s[i][j])
			}

		}
		fmt.Printf("\n")
	}
}

// 遍历寻找下一个要填充的坐标,如果返回失败,则说明填充完成
func findNextEmpty(s *[6][6]int) (currentCoordinate, bool) {
	for i := 0; i < len(s); i++ {
		for j := range s[i] {
			if s[i][j] == 0 {
				return currentCoordinate{
					abscissa: i,
					ordinate: j,
				}, true
			}
		}
	}
	return currentCoordinate{}, false
}

func checkValid(s *[6][6]int, fillNum int, cc currentCoordinate) bool {
	// 同时检查当前填充的数字在当前行和当前列是否重复
	for i := 0; i < len(s); i++ {
		if (s[cc.abscissa][i] == fillNum && cc.ordinate != i) ||
			(s[i][cc.ordinate] == fillNum && cc.abscissa != i) {
			return false
		}
	}
	// 检查当前填充的数字在当前粗线宫内是否重复
	boxX := cc.abscissa / 2
	boxY := cc.ordinate / 3
	for i := boxX * 2; i < boxX*2+2; i++ {
		for j := boxY * 3; j < boxY*3+3; j++ {
			if s[i][j] == fillNum && i != cc.abscissa && j != cc.ordinate {
				return false
			}
		}
	}
	return true
}

func tryFill(s *[6][6]int) bool {
	cc, find := findNextEmpty(s)
	if !find {
		return true
	}
	/*
		如果当前填充有效,则递归填充下一个,如果填充失败,
		则继续尝试填充剩余的数字,如果所有的数字都填充失败,
		说明在这之前填充过的某一个或多个位置填充错误,
		则返回失败到上一个填充的位置,并且清空该位置,
		然后继续尝试填充剩余的数字,如果填充失败,则继续
		向前回溯
	*/
	// 伪随机,尝试生成不同的解
	pseudoR := generateRandomNum(6)
	for i := pseudoR; i <= 6; i++ {
		if checkValid(s, i, cc) {
			s[cc.abscissa][cc.ordinate] = i
			if tryFill(s) {
				return true
			}
			s[cc.abscissa][cc.ordinate] = 0
		}
	}
	for i := 1; i < pseudoR; i++ {
		if checkValid(s, i, cc) {
			s[cc.abscissa][cc.ordinate] = i
			if tryFill(s) {
				return true
			}
			s[cc.abscissa][cc.ordinate] = 0
		}
	}
	return false
}

// [1,6]
func generateRandomNum(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max) + 1
}
