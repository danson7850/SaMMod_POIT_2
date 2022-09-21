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
func LehmerAlgorithm(A, m, n int, r float64) (result []float64) {

	for i := 0; i < n; i++ {
		r1 := r
		r = math.Mod(float64(A)*r1, float64(m))
		result = append(result, r/float64(m))
	}

	return
}

func TriangleDistribution(A, m, n int, a, b, r float64, check bool) (result []float64) {
	res := LehmerAlgorithm(A, m, n*2, r)

	for i := 0; i < len(res)-1; {
		if check {
			result = append(result, a+(b-a)*math.Max(res[i], res[i+1]))
		} else {
			result = append(result, a+(b-a)*math.Min(res[i], res[i+1]))
		}
		i += 2
	}

	return
}

func UniformDistribution(A, m, n int, a, b, r float64) (result []float64) {
	res := LehmerAlgorithm(A, m, n, r)

	for i := 0; i < n; i++ {
		result = append(result, a+(b-a)*res[i])
	}

	return
}

func GaussianDistribution(A, m, n int, r, mx, sx float64) (result []float64) {
	res := LehmerAlgorithm(A, m, n*12, r)
	for i := 0; i < n-13; {
		result = append(result, mx+sx*(utils.Sum(res[i:i+12])-6))
		i += 12
	}

	return
}

func ExponentialDistribution(A, m, n int, a, b, r, ly float64) (result []float64) {
	res := LehmerAlgorithm(A, m, n, r)

	for i := 0; i < n; i++ {
		result = append(result, math.Log(res[i])/ly)
	}

	return
}

func GammaDistribution(A, m, n, ny int, r, lambda float64) (result []float64) {
	res := LehmerAlgorithm(A, m, n*ny, r)
	for i := 0; i < n*ny-12; {
		sum := 0.0
		for j := 0; j < 12; {
			sum += math.Log(res[i+j])
			j++
		}
		result = append(result, (-1)*sum/lambda)
		i += 12
	}
	return
}

func SimsonDistribution(A, m, n int, a, b, r float64) (result []float64) {
	res := LehmerAlgorithm(A, m, n*2, r)

	for i := 0; i < len(res)-1; {
		result = append(result, (math.Max(a, b)-math.Min(a, b))*(res[i]+res[i+1]/2)+a)
		i += 2
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
func HistogramCalculation(data []float64) (maxNum int, ordinate [k]float64) {
	var values [k]float64
	var numbers [k]int
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

	maxNum = numbers[0]
	for _, i := range numbers {
		if i > maxNum {
			maxNum = i
		}
	}

	log.Println(numbers)
	log.Println(values)
	log.Println(ordinate)

	return
}

func UniEstimationCalculation(a, b float64) (mx, dx, sx float64) {
	mx = (a + b) / 2
	dx = math.Pow(b-a, 2) / 12
	sx = math.Sqrt(dx)
	return
}

func ExpEstimationCalculation(ly float64) (mx, dx, sx float64) {
	mx = 1 / ly
	dx = math.Pow(mx, 2)
	sx = mx

	return
}

func GammaEstimationCalculation(ny int, ly float64) (mx, dx, sx float64) {
	mx = float64(ny) / ly
	dx = float64(ny) / math.Pow(ly, 2)
	sx = math.Sqrt(dx)

	return
}
