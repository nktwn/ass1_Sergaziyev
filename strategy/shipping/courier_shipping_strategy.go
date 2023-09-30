package shipping

type CourierShippingStrategy struct{}

func (c *CourierShippingStrategy) Calculate(weight float64) float64 {
	return 10.0 * weight
}
