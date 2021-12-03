package game

func PercentageChange(old, new int) (delta float64) {
	diff := float64(new - old)
	delta = (diff / float64(old)) * 100
	return
}

func FindMinAndMax(scores []int) (min int, max int, index int) {
	min = scores[0]
	max = scores[0]
	var j = 0
	for i, value := range scores {
		if value < min {
			min = value
		}
		if value > max {
			max = value
			j = i
		}
	}
	return min, max, j
}

func FindAverage(scores []int) (avg float64) {
	sum := 0
	len := len(scores)
	for i, _ := range scores {
		sum += (scores[i])
	}
	avg = (float64(sum)) / (float64(len))
	return avg
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
