package arithmetic

// x,y 제곱의 합을 리턴
func (o *Numbers) SqaurePlus() int {
	return (o.X * o.X) + (o.Y * o.Y)
}

// x,y 제곱의 차를 리턴
func (o *Numbers) SquareMinus() int {
	return (o.X * o.X) - (o.Y * o.Y)
}
