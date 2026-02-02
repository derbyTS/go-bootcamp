package tax

type Calculator struct {
	rate             float64
	serviceFee       float64
	cachedMultiplier float64
}

func NewCalculator(rate float64, fee float64) *Calculator {
	combinedMultiplier := 1 + (rate / 100)

	return &Calculator{
		rate:             rate,
		serviceFee:       fee,
		cachedMultiplier: combinedMultiplier,
	}
}

func (c *Calculator) Total(amount float64) float64 {
	return (amount * c.cachedMultiplier) + c.serviceFee
}
