package utils

func MinmaxElements(data []float64) (min, max float64) {
	min, max = data[0], data[0]
	for _, i := range data {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}
	return
}

func Sum(array []float64) (result float64) {
	for _, v := range array {
		result += v
	}
	return
}
