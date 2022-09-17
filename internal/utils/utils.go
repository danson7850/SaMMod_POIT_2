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

func Period(data []float64) (result []int) {
	count := 0

	for i := 1; i < len(data); i++ {
		if count == 2 {
			break
		} else if data[i] == data[len(data)-1] {
			result = append(result, i)
			count++
		}
	}
	return result
}
