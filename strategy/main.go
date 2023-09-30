package main

import (
	"SDP/strategy/shipping"
	"fmt"
)

func main() {

	shippingCalculator := shipping.ShippingCalculator{
		Strategy: &shipping.CourierShippingStrategy{},
	}

	cost := shippingCalculator.Calculate(10.0)

	fmt.Println("Стоимость доставки:", cost)
}
