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
	ColLoop:
		for numCol := range spreadsheet[row] {
			numValue := spreadsheet[row][numCol]

			for denomCol := range spreadsheet[row] {
				if numCol == denomCol {
					continue
				}
				denomValue := spreadsheet[row][denomCol]

				if numValue%denomValue == 0 {
					checksum += numValue / denomValue
					break ColLoop
				}
			}
		}
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
