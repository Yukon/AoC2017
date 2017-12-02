package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter spreadsheet lines:\n")

	var spreadsheet [][]int
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			break
		} else {
			row := strings.Split(input, "\t")
			spreadsheet = addRow(row, spreadsheet)
		}
	}

	var checksum int
	for row := range spreadsheet {
		low, high := -1, -1
		for col := range spreadsheet[row] {
			value := spreadsheet[row][col]
			if value > high {
				high = value
			}
			if value < low || low == -1 {
				low = value
			}
		}

		checksum += high - low
	}

	fmt.Printf("\n%d\n", checksum)
}

func addRow(row []string, spreadsheet [][]int) [][]int {
	var newSpreadsheet [][]int
	if spreadsheet != nil {
		newSpreadsheet = make([][]int, len(spreadsheet)+1)
		newSpreadsheet[len(newSpreadsheet)-1] = make([]int, len(row))
		for i := range spreadsheet {
			newSpreadsheet[i] = spreadsheet[i]
		}
	} else {
		newSpreadsheet = make([][]int, 1)
		newSpreadsheet[0] = make([]int, len(row))
	}

	for i := range row {
		num, _ := strconv.ParseInt(row[i], 10, 32)
		newSpreadsheet[len(newSpreadsheet)-1][i] = int(num)
	}

	return newSpreadsheet
}
