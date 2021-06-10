package arithmetic

// 두 숫자의 사칙연산 계산 제공 패키지

// x, y 2개의 Integer 구조체

type Numbers struct {
	X int
	Y int
}

// x, y 합을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

// x ,y 차를 계산해서 반환
func (o *Numbers) Minus() int {
	return o.X - o.Y
}

func (o *Numbers) Product() int {
	return o.X * o.Y
}

func (o *Numbers) Devide() int {
	return o.X / o.Y
}
