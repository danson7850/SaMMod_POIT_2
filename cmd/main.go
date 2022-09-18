package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"sammod_2/internal/service"
	"strconv"
)

func main() {
	gtk.Init(nil)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	err = b.AddFromFile("main.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	obj, err := b.GetObject("window_main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	obj, _ = b.GetObject("text_a")
	entry1, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_m")
	entry2, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_r0")
	entry3, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_n")
	entry4, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_a_other")
	entry5, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_b_other")
	entry6, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_mx_other")
	entry7, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_sigma")
	entry8, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_lambda")
	entry9, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("text_nu")
	entry10, _ := obj.(*gtk.Entry)

	obj, _ = b.GetObject("alg_type")
	entry11, _ := obj.(*gtk.ComboBoxText)

	obj, _ = b.GetObject("check_type")
	entry12, _ := obj.(*gtk.CheckButton)

	obj, _ = b.GetObject("calc")
	button1 := obj.(*gtk.Button)

	obj, _ = b.GetObject("Mx")
	label1 := obj.(*gtk.Label)

	obj, _ = b.GetObject("Dx")
	label2 := obj.(*gtk.Label)

	obj, _ = b.GetObject("Sx")
	label3 := obj.(*gtk.Label)

	obj, _ = b.GetObject("hist")
	hist := obj.(*gtk.DrawingArea)

	hist.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
		cr.SetSourceRGB(0, 0, 0)
		cr.Rectangle(0, 0, 700, 350)
		cr.Fill()
		cr.SetSourceRGB(255, 255, 255)
		cr.Rectangle(1, 1, 698, 348)
		cr.Fill()
	})

	button1.Connect("clicked", func() {
		A, _ := entry1.GetText()
		m, _ := entry2.GetText()
		r, _ := entry3.GetText()
		n, _ := entry4.GetText()
		a, _ := entry5.GetText()
		B, _ := entry6.GetText()
		Mx, _ := entry7.GetText()
		sigma, _ := entry8.GetText()
		lambda, _ := entry9.GetText()
		nu, _ := entry10.GetText()
		algType := entry11.GetActiveID()
		check := entry12.GetActive()
		fmt.Println(algType)

		if err == nil {
			A1, _ := strconv.Atoi(A)
			m1, _ := strconv.Atoi(m)
			r1, _ := strconv.ParseFloat(r, 8)
			n1, _ := strconv.Atoi(n)
			a1, _ := strconv.ParseFloat(a, 8)
			b1, _ := strconv.ParseFloat(B, 8)
			mx1, _ := strconv.ParseFloat(Mx, 8)
			sx1, _ := strconv.ParseFloat(sigma, 8)
			lambda1, _ := strconv.ParseFloat(lambda, 8)
			nu1, _ := strconv.Atoi(nu)

			mx, dx, sx, result := chooseAlg(algType, A1, m1, n1, nu1, a1, b1, r1, lambda1, mx1, sx1, check)

			ordinate := service.HistogramCalculation(result)

			win.QueueDraw()

			hist.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
				cr.SetSourceRGB(0, 0, 0)
				cr.Rectangle(0, 0, 700, 350)
				cr.Fill()
				cr.SetSourceRGB(255, 255, 255)
				cr.Rectangle(1, 1, 698, 348)
				cr.Fill()
			})

			hist.Connect("draw", func(da *gtk.DrawingArea, cr *cairo.Context) {
				num1 := 0.0

				for _, num := range ordinate {
					cr.SetSourceRGB(0, 0, 0)
					cr.Rectangle(0+num1, 350, 35, -2000*num)
					cr.Fill()
					cr.SetSourceRGB(255, 255, 255)
					cr.Rectangle(0+num1, 350, 1, -3500*num)
					cr.Fill()
					num1 += 35
				}
			})
			log.Println(mx, dx, sx)
			label1.SetText(fmt.Sprintf("Mx: %f", mx))
			label2.SetText(fmt.Sprintf("Dx: %f", dx))
			label3.SetText(fmt.Sprintf("Sx: %f", sx))
		}
	})

	win.ShowAll()

	gtk.Main()
}

func chooseAlg(algType string, A1, m1, n1, nu1 int, a1, b1, r1, lambda1, mx1, sx1 float64, check bool) (float64, float64, float64, []float64) {
	switch algType {
	case "tri":
		result := service.TriangleDistribution(A1, m1, n1, a1, b1, r1, check)
		mx, dx, sx := service.EstimationCalculation(result)
		return mx, dx, sx, result
	case "gamma":
		result := service.GammaDistribution(A1, m1, n1, nu1, r1, lambda1)
		mx, dx, sx := service.GammaEstimationCalculation(nu1, lambda1)
		return mx, dx, sx, result
	case "exp":
		result := service.ExponentialDistribution(A1, m1, n1, a1, b1, r1, lambda1)
		mx, dx, sx := service.ExpEstimationCalculation(lambda1)
		return mx, dx, sx, result
	case "gaus":
		result := service.GaussianDistribution(A1, m1, n1, r1, mx1, sx1)
		mx, dx, sx := service.EstimationCalculation(result)
		return mx, dx, sx, result
	case "uni":
		result := service.UniformDistribution(A1, m1, n1, a1, b1, r1)
		mx, dx, sx := service.UniEstimationCalculation(a1, b1)
		return mx, dx, sx, result
	case "sim":
		result := service.SimsonDistribution(A1, m1, n1, a1, b1, r1)
		mx, dx, sx := service.EstimationCalculation(result)
		return mx, dx, sx, result
	default:
		result := service.UniformDistribution(A1, m1, n1, a1, b1, r1)
		mx, dx, sx := service.UniEstimationCalculation(a1, b1)
		return mx, dx, sx, result
	}
}
