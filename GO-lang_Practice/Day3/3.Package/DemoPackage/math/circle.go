package math

func CircleProperties(radius float64) (area, cir float64) {
	area = pI * (radius * radius)
	cir = 2 * pI * radius
	return
}
