package shipping

type ShippingStrategy interface {
	Calculate(weight float64) float64
}
