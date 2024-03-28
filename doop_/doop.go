package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	res, a, b := 0, 0, 0

	args := os.Args[1:]

	negative := false

	if args[0][0] == '-' {
		negative = true
		args[0] = args[0][1:]
	}

	for i := 0; i < len(args[0]); i++ {
		if args[0][i] < '0' || args[0][i] > '9' {
			return
		}
		a = a*10 + (int(args[0][i]) - 48)
		if a < 0 {
			return
		}
	}
	if negative {
		a = -a
		if a > 0 {
			return
		}
	}

	negative = false

	if args[2][0] == '-' {
		negative = true
		args[2] = args[2][1:]
	}

	for i := 0; i < len(args[2]); i++ {
		if args[2][i] < '0' || args[2][i] > '9' {
			return
		}
		b = b*10 + (int(args[2][i]) - 48)
		if b < 0 {
			return
		}
	}
	if negative {
		b = -b
		if b > 0 {
			return
		}
	}
	switch args[1] {
	case "+":
		res = a + b
		if a >= 0 && b >= 0 && res < 0 {
			return
		}
	case "-":
		res = a - b
		if a < b && res > 0 {
			return
		}
	case "*":
		for i := 0; i < Abs(b); i++ {
			res += a
			if a > 0 && b > 0 && res < 0 {
				return
			}
		}
		if b < 0 {
			res = -res
		}
	case "/":
		if b == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		res = a / b
	case "%":
		if b == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		res = a % b
	default:
		return
	}

	if res == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	result := ""

	tmp := res
	if res < 0 {
		tmp = -res
	}
	for ; tmp != 0; tmp /= 10 {
		result = string(rune(tmp%10+48)) + result
	}
	if res < 0 {
		result = "-" + result
	}

	os.Stdout.WriteString(result + "\n")

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
