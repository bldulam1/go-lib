package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func textToGrid(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	contents := make([][]int, 0)

	for scanner.Scan() {
		str := scanner.Text()
		splits := strings.Split(str, " ")
		intValues := make([]int, 0)
		for _, v := range splits {
			intVal, _ := strconv.Atoi(v)
			intValues = append(intValues, intVal)
		}
		contents = append(contents, intValues)
	}
	return contents
}

type coordinate struct {
	row int
	col int
}

func getAdjacentCoordinates(position coordinate, gridSize coordinate) [][]coordinate {
	row := position.row
	col := position.col

	isRight := col+3 < gridSize.col
	isDown := row+3 < gridSize.row
	isLeft := col-3 >= 0

	adjacentCoordinates := make([][]coordinate, 0)
	if isRight {
		right := make([]coordinate, 0)
		right = append(right, coordinate{row, col})
		right = append(right, coordinate{row, col + 1})
		right = append(right, coordinate{row, col + 2})
		right = append(right, coordinate{row, col + 3})

		adjacentCoordinates = append(adjacentCoordinates, right)
	}
	if isDown {
		down := make([]coordinate, 0)
		down = append(down, coordinate{row, col})
		down = append(down, coordinate{row + 1, col})
		down = append(down, coordinate{row + 2, col})
		down = append(down, coordinate{row + 3, col})

		adjacentCoordinates = append(adjacentCoordinates, down)
	}

	if isDown && isRight {
		rightD := make([]coordinate, 0)
		rightD = append(rightD, coordinate{row, col})
		rightD = append(rightD, coordinate{row + 1, col + 1})
		rightD = append(rightD, coordinate{row + 2, col + 2})
		rightD = append(rightD, coordinate{row + 3, col + 3})
		adjacentCoordinates = append(adjacentCoordinates, rightD)
	}
	if isDown && isLeft {
		leftD := make([]coordinate, 0)
		leftD = append(leftD, coordinate{row, col})
		leftD = append(leftD, coordinate{row + 1, col - 1})
		leftD = append(leftD, coordinate{row + 2, col - 2})
		leftD = append(leftD, coordinate{row + 3, col - 3})
		adjacentCoordinates = append(adjacentCoordinates, leftD)
	}

	return adjacentCoordinates
}

func largestProductInGrid(fileName string) int {
	grid := textToGrid(fileName)
	gridSize := coordinate{len(grid), len(grid[0])}
	maxProd := 1
	// maxPosition := coordinate{0, 0}

	for row := range grid {
		for col := range grid[row] {
			position := coordinate{row, col}
			neighbors := getAdjacentCoordinates(position, gridSize)
			for _, group := range neighbors {
				prod := 1
				for _, c := range group {
					prod *= grid[c.row][c.col]
				}

				if prod > maxProd {
					maxProd = prod
					// maxPosition = coordinate{row, col}
					// fmt.Println(maxProd, grid[row][col], coordinate{row, col})
				}
			}
		}
	}

	return maxProd
}
