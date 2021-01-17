package myatoi

import "testing"

func TestMyAtoi(t *testing.T) {

	if myAtoi("42") != 42 {
		t.Error(`myAtoi("42") != 42`)
	}

	if myAtoi("      -42") != -42 {
		t.Error(`myAtoi("      -42") != -42`)
	}

	if myAtoi("4193 with words") != 4193 {
		t.Error(`myAtoi("4193 with words") != 4193`)
	}

	if myAtoi("-91283472332") != -2147483648 {
		t.Error(`myAtoi("-91283472332") != -2147483648`)
	}

	if myAtoi("+1") != 1 {
		t.Error(`myAtoi("+1") != 1`)
	}

	if myAtoi("+-12") != 0 {
		t.Error(`myAtoi("+-12") != 0`)
	}

	if myAtoi("21474836460") != 2147483647 {
		t.Error(`myAtoi("21474836460")  != 2147483647`)
	}

	if myAtoi("00000-42a1234") != 0 {
		t.Error(`myAtoi("00000-42a1234") != 0`)
	}

	if myAtoi("3.14159") != 3 {
		t.Error(`myAtoi("3.14159") != 3`)
	}

	if myAtoi(".1") != 0 {
		t.Error(`myAtoi(".1") != 0`)
	}

	if myAtoi("  -0012a42") != -12 {
		t.Error(`myAtoi("  -0012a42") != -12`)
	}
}
