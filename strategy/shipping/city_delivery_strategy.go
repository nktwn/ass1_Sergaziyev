package shipping

type CityDeliveryStrategy struct{}

func (c *CityDeliveryStrategy) Calculate(weight float64) float64 {
	return 2.0 * weight
}
