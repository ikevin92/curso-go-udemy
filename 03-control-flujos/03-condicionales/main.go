package main

import "fmt"

func main() {
	/** App para restaurante
	* Descuentos de 10% hasta 100 de consumo
	* Descuentos de 20% hasta 200 de consumo
	* Descuentos de 30% mayor a 200 de consumo
	* iva 19%
	 */

	var consumo, descuento float64
	var datosDescuento string

	//Entrada de datos
	fmt.Println("Ingrese el consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			// descuento de 10%
			datosDescuento = "10%"
			descuento = consumo * 0.10
		} else if consumo > 100 && consumo <= 200 {
			// descuento de 20%
			datosDescuento = "20%"
			descuento = consumo * 0.20
		} else if consumo > 200 {
			// descuento de 30%
			datosDescuento = "30%"
			descuento = consumo * 0.30
		}
	} else {
		fmt.Println("Error al ingresar el consumo")
	}

	// Operaciones
	montoDescuento := consumo - descuento
	iva := montoDescuento * 0.19
	totalPagar := montoDescuento + iva

	// Salida de datos
	fmt.Println("------FACTURA DE CONSUMO----")
	fmt.Println("Descuento que aplica:", datosDescuento)
	fmt.Println("Consumo:", consumo)
	fmt.Println("Descuento:", descuento)
	fmt.Println("Monto con descuento: ", montoDescuento)
	fmt.Println("IVA:", iva)
	fmt.Println("Total a pagar:", totalPagar)

}
