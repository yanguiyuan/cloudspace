package math

// GCD 计算最大公约数
func GCD(a, b int) int {
	if a < 0 || b < 0 {
		panic("GCD only accepts non-negative integers")
	}
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
