package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var inputAmount = flag.Float64("m", 50_000, "Cantidad a depositar mensualmente")
	var inputPeriodos = flag.Int("p", 408, "Períodos en meses") // 34 años
	var inputTasa = flag.Float64("t", 1.0035, "Tasa de crecimiento mensual")
	flag.Parse()

	var sum float64 = 0
	for i := 0; i < *inputPeriodos; i++ {
		sum += math.Pow(*inputTasa, float64(i))
	}

	total := *inputAmount * sum
	rentaMensual := total / (12 * 20)

	fmt.Println("=======================================")
	fmt.Println("RENTA MENSUAL AHORRO VOLUNTARIO APV\nDepósito por defecto: 50.000 CLP\nPeríodos por defecto: 34 años")
	fmt.Println("=======================================")
	fmt.Printf("Monto total               %.f CLP\n", total)
	fmt.Printf("Renta mensual a 20 años     %.f CLP\n", rentaMensual)
	fmt.Println("---------------------------------------")
}
