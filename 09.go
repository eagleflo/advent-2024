package main

import (
	"fmt"
	"os"
	"strconv"
)

type Block struct {
	id   int
	free bool
}

type Disk struct {
	blocks []Block
	id_seq int
}

func readDiskMap(diskMap string) Disk {
	blocks := []Block{}
	id_seq := 0
	for i, s := range diskMap {
		n, _ := strconv.Atoi(string(s))
		if i%2 != 1 {
			// File blocks
			for range n {
				blocks = append(blocks, Block{id_seq, false})
			}
			id_seq += 1
		} else {
			// Free blocks
			for range n {
				blocks = append(blocks, Block{-1, true})
			}
		}
	}

	return Disk{blocks, id_seq}
}

func compactDisk(disk Disk) Disk {
	for {
		firstEmpty := -1
		for i, r := range disk.blocks {
			if r.free {
				firstEmpty = i
				break
			}
		}
		if firstEmpty == -1 {
			return disk
		}

		lastBlock := disk.blocks[len(disk.blocks)-1]
		disk.blocks[firstEmpty] = lastBlock
		disk.blocks[len(disk.blocks)-1].free = true
		disk.blocks = disk.blocks[:len(disk.blocks)-1]
	}
}

func checksum(disk Disk) int {
	sum := 0
	for i, b := range disk.blocks {
		if !b.free {
			sum += i * b.id
		}
	}
	return sum
}

func printDisk(disk Disk) {
	for _, b := range disk.blocks {
		if !b.free {
			fmt.Print(string(b.id + '0'))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func advent09_1() {
	bytes, _ := os.ReadFile("09.txt")
	data := string(bytes)
	disk := readDiskMap(data)
	compacted := compactDisk(disk)
	fmt.Println(checksum(compacted))
}

func defragmentDisk(disk Disk) Disk {
	for i := disk.id_seq - 1; i > 0; i-- {
		// printDisk(disk)

		position := -1
		length := 0
		for j, r := range disk.blocks {
			if r.id == i {
				if position == -1 {
					position = j
				}
				length += 1
			}
		}
		// fmt.Println("Found file with seq_id", i, "at position", position, "with length of", length)

		// Find continuous free block to the left of position
		freePosition := -1
		freeLength := 0
		for k, r := range disk.blocks[:position] {
			if r.free {
				if freePosition == -1 {
					freePosition = k
				}
				freeLength += 1

				if freeLength == length {
					// fmt.Println("Found free block of length", freeLength, "at position", freePosition)
					// Move block
					for l, b := range disk.blocks[position : position+length] {
						disk.blocks[freePosition+l].id = b.id
						disk.blocks[freePosition+l].free = false
						disk.blocks[position+l].id = -1
						disk.blocks[position+l].free = true
					}
					break
				}
			} else {
				freePosition = -1
				freeLength = 0
			}
		}

	}
	return disk
}

func advent09_2() {
	bytes, _ := os.ReadFile("09.txt")
	data := string(bytes)
	disk := readDiskMap(data)
	defragged := defragmentDisk(disk)
	fmt.Println(checksum(defragged))
}

func main() {
	advent09_1()
	advent09_2()
}
