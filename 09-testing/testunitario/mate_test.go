package testunitario

import "testing"

/* func TestSuma(t *testing.T) {
	total := Suma(10, 15)

	if total != 25 {
		t.Errorf("Suma Incorrecta, de 10 + 15 = %d; esperaba 25", total)
	}
} */

func TestSuma(t *testing.T) {
	tabla := []struct {
		a, b, total int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.total {
			t.Errorf("Suma Incorrecta, de %d + %d = %d; esperaba %d", item.a, item.b, total, item.total)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a, b, max int
	}{
		{1, 2, 2},
		{3, 2, 3},
		{-1, 1, 1},
		{100, 200, 200},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)
		if max != item.max {
			t.Errorf("GetMax Incorrecto, de %d y %d = %d; esperaba %d", item.a, item.b, max, item.max)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		n, fib int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, item := range tabla {
		fib := Fibonacci(item.n)
		if fib != item.fib {
			t.Errorf("Fibonacci Incorrecto, de %d = %d; esperaba %d", item.n, fib, item.fib)
		}
	}

}
