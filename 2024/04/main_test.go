package main

import "testing"

func TestIsPrevious(t *testing.T) {
	performTest(false, isPreviousLetter('X', 'X', 'f'), t)
	performTest(false, isPreviousLetter('X', 'M', 'f'), t)
	performTest(false, isPreviousLetter('X', 'A', 'f'), t)
	performTest(false, isPreviousLetter('X', 'S', 'f'), t)

	performTest(true, isPreviousLetter('M', 'X', 'f'), t)
	performTest(false, isPreviousLetter('M', 'M', 'f'), t)
	performTest(false, isPreviousLetter('M', 'A', 'f'), t)
	performTest(false, isPreviousLetter('M', 'S', 'f'), t)

	performTest(false, isPreviousLetter('A', 'X', 'f'), t)
	performTest(true, isPreviousLetter('A', 'M', 'f'), t)
	performTest(false, isPreviousLetter('A', 'A', 'f'), t)
	performTest(false, isPreviousLetter('A', 'S', 'f'), t)

	performTest(false, isPreviousLetter('S', 'X', 'f'), t)
	performTest(false, isPreviousLetter('S', 'M', 'f'), t)
	performTest(true, isPreviousLetter('S', 'A', 'f'), t)
	performTest(false, isPreviousLetter('S', 'S', 'f'), t)

	performTest(false, isPreviousLetter('S', 'S', 'r'), t)
	performTest(false, isPreviousLetter('S', 'A', 'r'), t)
	performTest(false, isPreviousLetter('S', 'M', 'r'), t)
	performTest(false, isPreviousLetter('S', 'X', 'r'), t)

	performTest(true, isPreviousLetter('A', 'S', 'r'), t)
	performTest(false, isPreviousLetter('A', 'A', 'r'), t)
	performTest(false, isPreviousLetter('A', 'M', 'r'), t)
	performTest(false, isPreviousLetter('A', 'X', 'r'), t)

	performTest(false, isPreviousLetter('M', 'S', 'r'), t)
	performTest(true, isPreviousLetter('M', 'A', 'r'), t)
	performTest(false, isPreviousLetter('M', 'M', 'r'), t)
	performTest(false, isPreviousLetter('M', 'X', 'r'), t)

	performTest(false, isPreviousLetter('X', 'S', 'r'), t)
	performTest(false, isPreviousLetter('X', 'A', 'r'), t)
	performTest(true, isPreviousLetter('X', 'M', 'r'), t)
	performTest(false, isPreviousLetter('X', 'X', 'r'), t)
}

func performTest(want, res bool, t *testing.T) {
	if want != res {
		t.Fatalf("SAF: %t != %t", want, res)
	}
}
