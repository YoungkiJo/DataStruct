package pointlist

import "fmt"

type Point struct {
	xpos int
	ypos int
}

func (ppos *Point) SetPointPos(xpos int, ypos int) {
	ppos.xpos = xpos
	ppos.ypos = ypos
}

func (ppos *Point) ShowPointPos() {
	fmt.Printf("[%d, %d] \n", ppos.xpos, ppos.ypos)
}

func PointComp(pos1 *Point, pos2 *Point) int {
	if pos1.xpos == pos2.xpos && pos1.ypos == pos2.ypos {
		return 0
	} else if pos1.xpos == pos2.xpos {
		return 1
	} else if pos1.ypos == pos2.ypos {
		return 2
	} else {
		return -1
	}
}
