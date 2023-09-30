package shipping

type ShippingCalculator struct {
	Strategy ShippingStrategy
}

func (s *ShippingCalculator) Calculate(weight float64) float64 {
	return s.Strategy.Calculate(weight)
}
