package main

import (
	"fmt"
	"os"
	"strings"
)

type Path struct {
	x []int
	y []int
}

func (p *Path) append(x, y int) {
	p.x = append(p.x, x)
	p.y = append(p.y, y)
}

var (
	moves = map[uint8][]int{
		'-': {1, 0},
		'|': {0, 1},
		'J': {-1, -1},
		'L': {1, -1},
		'7': {-1, 1},
		'F': {1, 1},
	}
)

func main() {
	var path Path
	input := os.Args[1]
	splittedInput := strings.Split(input, "\n")

	path.findStart(&splittedInput)
	path.findFirstStep(&splittedInput)
	fmt.Println(path.x)
	fmt.Println(path.y)
	curSymb := splittedInput[path.lastY()][path.lastX()]
	for curSymb != 'S' {
		x, y := path.step(curSymb)
		curSymb = splittedInput[y][x]
	}
	fmt.Println("Farthest dist:", path.dist())

}

func (p *Path) dist() int {
	return len(p.x) / 2
}

func (p *Path) step(curSymb uint8) (int, int) {
	fmt.Println(string(curSymb))
	mov := moves[curSymb]

	curPos := []int{p.lastX(), p.lastY()}
	prevPos := []int{p.prevX(), p.prevY()}
	if mov[0] == 0 {
		if prevPos[1]+mov[1] == curPos[1] {
			x, y := curPos[0]+mov[0], curPos[1]+mov[1]
			p.append(x, y)
			return x, y
		} else {
			x, y := curPos[0]-mov[0], curPos[1]-mov[1]
			p.append(x, y)
			return x, y
		}
	} else if mov[1] == 0 {
		if prevPos[0]+mov[0] == curPos[0] {
			x, y := curPos[0]+mov[0], curPos[1]+mov[1]
			p.append(x, y)
			return x, y
		} else {
			x, y := curPos[0]-mov[0], curPos[1]-mov[1]
			p.append(x, y)
			return x, y
		}
	} else {
		if prevPos[0] == curPos[0] {
			// going: up or down
			x, y := curPos[0]+mov[0], curPos[1]
			p.append(x, y)
			return x, y
		} else {
			x, y := curPos[0], curPos[1]+mov[1]
			// going: back or forward
			p.append(x, y)
			return x, y
		}
	}
}

func (p *Path) lastX() int {
	if len(p.x) > 0 {
		return p.x[len(p.x)-1]
	}
	return -1
}

func (p *Path) prevX() int {
	if len(p.x) > 1 {
		return p.x[len(p.x)-2]
	}
	return -1
}

func (p *Path) lastY() int {
	if len(p.x) > 0 {
		return p.y[len(p.y)-1]
	}
	return -1
}

func (p *Path) prevY() int {
	if len(p.x) > 1 {
		return p.y[len(p.y)-2]
	}
	return -1
}

func (p *Path) findStart(fields *[]string) {
	for yPos, line := range *fields {
		xPos := strings.Index(line, "S")
		if xPos != -1 {
			p.append(xPos, yPos)
			return
		}
	}
}

func (p *Path) findFirstStep(fields *[]string) {
	height := len(*fields)
	width := len((*fields)[0])
	for y := p.y[0] - 1; y <= p.y[0]+1; y++ {
		if y < 0 || y >= height {
			continue
		}
		for x := p.x[0] - 1; x <= p.x[0]+1; x++ {
			if x < 0 || x >= width || (x != p.x[0] && y != p.y[0]) {
				continue
			}
			if (*fields)[y][x] != '.' && (*fields)[y][x] != 'S' {
				//fmt.Println(string((*fields)[y][x]))
				p.append(x, y)
				return
			}
		}
	}
}
