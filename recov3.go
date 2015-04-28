package main

/*
POR HACER:

-Mejorar estilo (SetStyle es una puta mierda con el de go)
*/
import (
	"flag"
	"fmt"
	"github.com/sbinet/go-gnuplot"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	p = flag.String("p", "ninonino", "p:riemann,mandelbrot,saxpy,primes")
	numPruebas int
	tiempos []float64
)

func start() {
	f, _ := ioutil.ReadFile("/home/acastano/GoLand/scripts/resultados.txt")
	f1 := strings.Split(string(f), "\n")
	fmt.Println(f1)
	f2 := make([]float64, len(f1)-1)

	for i := 0; i < len(f2); i++ {
		f2[i], _ = strconv.ParseFloat(f1[i], 64)
	}
	fmt.Println("f2: ", f2, len(f2))
	for i := 0; i < numPruebas; i++ {
		tiempos[i] = tiempoMedio(f2[1*i : 1*i+1])
	}
	fmt.Println("tiempos", tiempos, "tiempos[0:4]", tiempos[0:4], "tiempos[4:8]", tiempos[4:8])
}

func plot() *gnuplot.Plotter {
	fname := ""
	persist := false
	debug := true

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	return p
}

func tiempoMedio(lista []float64) float64 {
	var res float64
	for i := 0; i < len(lista); i++ {
		res1, _ := lista[i], 64
		res += res1
	}
	res = res / float64(len(lista))
	return res
}

func main() {
	numPruebas = 8
	flag.Parse()
	tiempos = make([]float64, numPruebas)

	switch *p {
	case "peter":
		start()
		p1 := plot()
		p1.SetStyle("lines")
		p1.PlotNd("yeahbitch", []float64{1, 2, 3, 4, 5, 6, 7}, []float64{8, 9, 15, 20, 50})
		p1.PlotNd("yeahbitch", []float64{7, 6, 5, 4, 3, 2, 1}, []float64{50, 20, 20, 15, 8})
		p1.SetXLabel("operations")
		p1.SetYLabel("time")
		p1.SetZLabel("pet")
		p1.CheckedCmd("set terminal pdf")
		aux := "set output '" + *p + ".pdf" + "'"
		p1.CheckedCmd(aux)
		p1.CheckedCmd("replot")
		p1.CheckedCmd("q")
	case "riemann":
		start()
		p1 := plot()
		p1.SetStyle("lines")
		p1.PlotXY([]float64{1e8, 1e9, 2e9, 5e9, 1e10}, tiempos[0:4], "Secuencial")
		p1.PlotXY([]float64{1e8, 1e9, 2e9, 5e9, 1e10}, tiempos[4:8], "Paralelo")
		p1.SetXLabel("operations")
		p1.SetYLabel("time")
		p1.CheckedCmd("set terminal pdf")
		aux := "set output '" + *p + ".pdf" + "'"
		p1.CheckedCmd(aux)
		p1.CheckedCmd("replot")
		p1.CheckedCmd("q")

	case "mandelbrot":
		start()
		p1 := plot()
		p1.SetStyle("boxes")
		p1.PlotXY([]float64{1, 4, 8, 12, 20}, tiempos[0:2], "Secuencial")
		p1.PlotXY([]float64{0.25, 1, 2, 3, 5}, tiempos[3:5], "Paralelo")
		p1.SetXLabel("operations")
		p1.SetYLabel("time")
		p1.CheckedCmd("set terminal pdf")
		aux := "set output '" + *p + ".pdf" + "'"
		p1.CheckedCmd(aux)
		p1.CheckedCmd("replot")
		p1.CheckedCmd("q")

	case "saxpy":
		start()
		p1 := plot()
		p1.SetStyle("boxes")
		p1.PlotXY([]float64{1, 4, 8, 12, 20}, tiempos[0:2], "Secuencial")
		p1.PlotXY([]float64{0.25, 1, 2, 3, 5}, tiempos[3:5], "Paralelo")
		p1.SetXLabel("operations")
		p1.SetYLabel("time")
		p1.CheckedCmd("set terminal pdf")
		aux := "set output '" + *p + ".pdf" + "'"
		p1.CheckedCmd(aux)
		p1.CheckedCmd("replot")
		p1.CheckedCmd("q")

	case "primes":
		start()
		p1 := plot()
		p1.SetStyle("boxes")
		p1.PlotXY([]float64{1, 4, 8, 12, 20}, tiempos[0:2], "Secuencial")
		p1.PlotXY([]float64{0.25, 1, 2, 3, 5}, tiempos[3:5], "Paralelo")
		p1.SetXLabel("operations")
		p1.SetYLabel("time")
		p1.CheckedCmd("set terminal pdf")
		aux := "set output '" + *p + ".pdf" + "'"
		p1.CheckedCmd(aux)
		p1.CheckedCmd("replot")
		p1.CheckedCmd("q")

	default:
		fmt.Println("Fuck You!")
	}
}
