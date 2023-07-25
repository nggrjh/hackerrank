package solution

import (
	"strconv"
	"strings"
)

func getMax(operations []string) []int32 {
	var (
		prints   []int32
		elements []int32

		maxIdx int
	)

	for _, op := range operations {
		cmds := strings.Split(op, " ")

		cmd, _ := strconv.Atoi(cmds[0])
		switch cmd {
		case 1:
			elem, _ := strconv.Atoi(cmds[1])
			elements = append(elements, int32(elem))

			if lastIdx := len(elements) - 1; elements[lastIdx] > elements[maxIdx] {
				maxIdx = lastIdx
			}

		case 2:
			lastIdx := len(elements) - 1
			elements = elements[:lastIdx]

			if lastIdx == maxIdx {
				maxIdx = 0
				for i, e := range elements {
					if e > elements[maxIdx] {
						maxIdx = i
					}
				}
			}

		case 3:
			prints = append(prints, elements[maxIdx])

		default:

		}
	}

	return prints
}
