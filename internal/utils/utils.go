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

func OrdinateExp(input [20]float64) (result [20]float64) {
	for i := 19; i > 0; i-- {
		result[19-i] = input[i]
	}
	return
}
