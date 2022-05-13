package bankocr

import (
	"testing"
)

func TestParseZeros(t *testing.T) {
	in := `
 _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|`
	want := "000000000"
	out := Parse(in)
	if out != want {
		t.Errorf("Parse(%v) = %v; want %v", in, out, want)
	}
}

func TestParseNumbers(t *testing.T) {
	in := `
    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|`
	want := "123456789"
	out := Parse(in)
	if out != want {
		t.Errorf("Parse(%v) = %v; want %v", in, out, want)
	}
}
