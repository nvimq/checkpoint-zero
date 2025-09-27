package piscine

import "strconv"

func FromTo(from int, to int) string {
	if from < 0 || to < 0 || from > 99 || to > 99 {
		return "Invalid\n"
	}

	var res string
	step := 1
	if from > to {
		step = -1
	}

	for i := from; ; i += step {
		if i < 10 {
			res += "0" + strconv.Itoa(i)
		} else {
			res += strconv.Itoa(i)
		}
		if i == to {
			break
		}
		res += ", "
	}
	return res + "\n"
}
