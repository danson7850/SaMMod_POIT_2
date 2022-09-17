package service

import (
	"log"
	"math"
	"sammod_2/internal/utils"
)

// k - const that we need to calculate histogram
const k = 20

// LehmerAlgorithm - function that generates fake random numbers
// Parameters:
// a, m - random int numbers
// r - start number
// n - count of iterations (new generated numbers)
// result - array of int numbers
func LehmerAlgorithm(a, m, n int, r float64) (result []float64) {

	for i := 0; i < n; i++ {
		r1 := r
		r = math.Mod(float64(a)*r1, float64(m))
		result = append(result, r/float64(m))
	}

	return
}

// EstimationCalculation - function which provides calculation of math.expectation, dispersion
// and rms(sqrt from dispersion)
func EstimationCalculation(data []float64) (mx, dx, sx float64) {

	mx, dx = 0, 0

	for _, i := range data {
		mx += i
	}
	mx /= float64(len(data))

	for _, i := range data {
		dx += math.Pow(i-mx, float64(2))
	}

	dx /= float64(len(data)) - 1

	sx = math.Sqrt(dx)

	return
}

// HistogramCalculation - function
func HistogramCalculation(data []float64) (ordinate [k]float64) {
	var numbers [k]int
	var values [k]float64

	min, max := utils.MinmaxElements(data)
	rVar := max - min
	delta := rVar / k
	left, right := min, min+delta

	for i := 0; i < k; i++ {
		for j := 0; j < len(data); j++ {
			if data[j] >= left && data[j] < right {
				numbers[i]++
			}
		}
		values[i] = (left + right) / 2
		left += delta
		right += delta

		ordinate[i] = float64(numbers[i]) / float64(len(data))
	}
	log.Println(numbers)
	log.Println(values)
	log.Println(ordinate)

	return
}

// UniformityChecker - function which checks that our numbers close to π/4 or not
func UniformityChecker(data []float64) float64 {

	count := 0

	for i := 0; i < len(data)/2; i++ {
		if math.Pow(data[2*i], 2)+math.Pow(data[2*i+1], 2) < 1 {
			count++
		}
	}

	return 2 * float64(count) / float64(len(data))
}

// AperiodicCalculation - function which calculates period and aperiod of our number sequence
func AperiodicCalculation(data []float64, n, m int) (period, aperiod int) {

	indexes := utils.Period(data)

	if len(indexes) > 1 {
		period = indexes[1] - indexes[0]
	} else {
		//TODO: add error
		period = indexes[0] + 1
	}

	if len(indexes) < 2 {
		aperiod = int(math.Min(float64(n), float64(m)))
	} else {
		res := 0
		for i := indexes[1]; i > period; i-- {
			if data[i] != data[i-period] {
				res = i
				break
			}
		}
		if res == 0 {
			res = indexes[0]
		}
		aperiod = res + period + 1
	}

	return
}
