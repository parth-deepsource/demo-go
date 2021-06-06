package main

func HexLiteral() bool {
	x := 0xFff
	y := 0xFFF
	z := 0Xfff

	return (x == y) && (y == z)
}
