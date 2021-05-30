package contest

import "strconv"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	dic := make(map[string]int, 0)
	dic["a"] = 0
	dic["b"] = 1
	dic["c"] = 2
	dic["d"] = 3
	dic["e"] = 4
	dic["f"] = 5
	dic["g"] = 6
	dic["h"] = 7
	dic["i"] = 8
	dic["j"] = 9
	fi, se, ta := 0, 0, 0
	for i := 0; i < len(firstWord); i++ {
		fi = fi*10 + dic[firstWord[i:i+1]]
	}
	for i := 0; i < len(secondWord); i++ {
		se = se*10 + dic[secondWord[i:i+1]]
	}
	for i := 0; i < len(targetWord); i++ {
		ta = ta*10 + dic[targetWord[i:i+1]]
	}
	return fi+se == ta
}

func maxValue(n string, x int) string {
	if n[0:1] == "-" {
		for i := 1; i < len(n); i++ {
			in, _ := strconv.Atoi(n[i : i+1])
			if x < in {
				return n[0:i] + strconv.Itoa(x) + n[i:]
			}
		}
		return n + strconv.Itoa(x)
	} else {
		for i := 0; i < len(n); i++ {
			in, _ := strconv.Atoi(n[i : i+1])
			if x > in {
				return n[0:i] + strconv.Itoa(x) + n[i:]
			}
		}
		return n + strconv.Itoa(x)
	}
}
