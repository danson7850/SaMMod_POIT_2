package main

import (
	"fmt"
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"math"
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

	obj, _ = b.GetObject("calc")
	button1 := obj.(*gtk.Button)

	obj, _ = b.GetObject("Mx")
	label1 := obj.(*gtk.Label)

	obj, _ = b.GetObject("Dx")
	label2 := obj.(*gtk.Label)

	obj, _ = b.GetObject("Sx")
	label3 := obj.(*gtk.Label)

	obj, _ = b.GetObject("check")
	label4 := obj.(*gtk.Label)

	obj, _ = b.GetObject("period")
	label5 := obj.(*gtk.Label)

	obj, _ = b.GetObject("aperiod")
	label6 := obj.(*gtk.Label)

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
		a, _ := entry1.GetText()
		m, _ := entry2.GetText()
		r, _ := entry3.GetText()
		n, _ := entry4.GetText()

		if err == nil {
			a1, _ := strconv.Atoi(a)
			m1, _ := strconv.Atoi(m)
			r1, _ := strconv.ParseFloat(r, 8)
			n1, _ := strconv.Atoi(n)

			result := service.LehmerAlgorithm(a1, m1, n1, r1)
			mx, dx, sx := service.EstimationCalculation(result)
			un := service.UniformityChecker(result)
			per, aper := service.AperiodicCalculation(result, n1, m1)
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
					cr.Rectangle(0+num1, 350, 35, -3500*num)
					cr.Fill()
					cr.SetSourceRGB(255, 255, 255)
					cr.Rectangle(0+num1, 350, 1, -3500*num)
					cr.Fill()
					num1 += 35
				}
			})

			label1.SetText(fmt.Sprintf("Mx: %f", mx))
			label2.SetText(fmt.Sprintf("Dx: %f", dx))
			label3.SetText(fmt.Sprintf("Sx: %f", sx))
			label4.SetText(fmt.Sprintf("π/4 check: %.7f ---> %f\n", un, math.Pi/4))
			label5.SetText(fmt.Sprintf("Period: %d", per))
			label6.SetText(fmt.Sprintf("Aperiod: %d", aper))
		}
	})

	win.ShowAll()

	gtk.Main()
}
