package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Player struct {
	pos   Point
	steps []Point
}

func makeGrid(height int, width int) [][]rune {
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	return grid
}

func printGrid(grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}

func findNextPlayerSteps(grid [][]rune, player Player) []Point {
	height := len(grid)
	width := len(grid[0])
	point := player.pos

	north := Point{point.x, point.y - 1}
	east := Point{point.x + 1, point.y}
	south := Point{point.x, point.y + 1}
	west := Point{point.x - 1, point.y}

	result := []Point{}
	if north.y >= 0 && grid[north.y][north.x] == '.' && !slices.Contains(player.steps, north) {
		result = append(result, north)
	}
	if east.x < width && grid[east.y][east.x] == '.' && !slices.Contains(player.steps, east) {
		result = append(result, east)
	}
	if south.y < height && grid[south.y][south.x] == '.' && !slices.Contains(player.steps, south) {
		result = append(result, south)
	}
	if west.x >= 0 && grid[west.y][west.x] == '.' && !slices.Contains(player.steps, west) {
		result = append(result, west)
	}
	return result
}

var height = 71
var width = 71
var blocks = 1024

func advent18_1() {
	grid := makeGrid(height, width)

	bytes, _ := os.ReadFile("18.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	for _, line := range lines[:blocks] {
		coords := strings.Split(line, ",")
		x, _ := strconv.ParseInt(coords[0], 10, 8)
		y, _ := strconv.ParseInt(coords[1], 10, 8)
		grid[y][x] = '#'
	}

	var visited = make(map[Point]int)
	player := Player{Point{0, 0}, []Point{Point{0, 0}}}
	goal := Point{width - 1, height - 1}
	players := []Player{player}
	routes := [][]Point{}
	for {
		newPlayers := []Player{}
		for _, player := range players {
			steps := findNextPlayerSteps(grid, player)
			for _, step := range steps {
				newSteps := append(player.steps, step)
				if visited[step] >= len(newSteps) {
					continue
				} else {
					visited[step] = len(newSteps)
				}
				newPlayers = append(newPlayers, Player{step, newSteps})
				if step == goal {
					routes = append(routes, player.steps)
				}
			}
		}
		if len(newPlayers) == 0 {
			break
		} else {
			players = newPlayers
		}
	}

	min := height * width
	for _, route := range routes {
		if len(route) < min {
			min = len(route)
		}
	}
	fmt.Println(min)
}

// Just find any route
func findARoute(grid [][]rune) []Point {
	var visited = make(map[Point]int)
	player := Player{Point{0, 0}, []Point{Point{0, 0}}}
	goal := Point{width - 1, height - 1}
	players := []Player{player}
	for {
		newPlayers := []Player{}
		for _, player := range players {
			steps := findNextPlayerSteps(grid, player)
			for _, step := range steps {
				newSteps := append(player.steps, step)
				if visited[step] >= len(newSteps) {
					continue
				} else {
					visited[step] = len(newSteps)
				}
				newPlayers = append(newPlayers, Player{step, newSteps})
				if step == goal {
					return player.steps
				}
			}
		}
		if len(newPlayers) == 0 {
			break
		} else {
			players = newPlayers
		}
	}
	return nil
}

func advent18_2() {
	grid := makeGrid(height, width)
	bytes, _ := os.ReadFile("18.txt")
	data := strings.Split(string(bytes), "\n")
	lines := data[:len(data)-1]
	route := findARoute(grid)

	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		p := Point{x, y}
		grid[y][x] = '#'

		if slices.Contains(route, p) {
			route = findARoute(grid)
		}

		fmt.Println("Checked", i+1, "blocks")

		if route == nil {
			fmt.Println(x, y)
			break
		}
	}

}

func main() {
	advent18_1()
	advent18_2()
}
