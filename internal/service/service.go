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
func LehmerAlgorithm(a, m int, r float64) (result float64) {

	r1 := r
	r = math.Mod(float64(a)*r1, float64(m))

	return r / float64(m)
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
