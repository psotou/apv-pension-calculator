package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var inputAmount = flag.Float64("m", 50_000, "Cantidad a depositar mensualmente")
	var inputPeriodos = flag.Int("p", 408, "Períodos en meses") // 34 años
	var inputTasa = flag.Float64("t", 0.35, "Tasa de crecimiento mensual")
	flag.Parse()

	var tasa float64 = 1 + *inputTasa/100

	var sum float64 = 0
	for i := 1; i < *inputPeriodos+1; i++ {
		sum += math.Pow(tasa, float64(i))
	}

	total := *inputAmount * sum
	rentaMensual := total / (12 * 20)
	periodAnios := float64(*inputPeriodos) / 12

	fmt.Println("=======================================")
	fmt.Println("RENTA MENSUAL AHORRO VOLUNTARIO APV")
	fmt.Printf("Depósito mensual     %.f CLP\n", *inputAmount)
	fmt.Printf("Tasa crecimiento      %.2f %%\n", *inputTasa)
	fmt.Printf("Período               %.1f años\n", periodAnios)
	fmt.Println("=======================================")
	fmt.Printf("Monto total               %.f CLP\n", total)
	fmt.Printf("Renta mensual a 20 años     %.f CLP\n", rentaMensual)
	fmt.Println("---------------------------------------")
}
