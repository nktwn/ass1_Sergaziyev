package shipping

type PostalShippingStrategy struct{}

func (p *PostalShippingStrategy) Calculate(weight float64) float64 {
	return 5.0 * weight
}
