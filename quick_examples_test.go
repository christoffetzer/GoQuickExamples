/* (C) Christof Fetzer, 2016
*
* simple examples to demonstrate advantages / disadvantages of random testing
*
* run by executing "go test -v"
*/

package quick_examples

import (
    "math"
    "testing"
    "testing/quick"
    "github.com/christoffetzer/extremeValues"
)

func TestAbs1(t *testing.T) {
	f := func(x int) bool {
		return Abs1(x) >= 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestAbs2 out of the box, random testing typically does not find the bug
func TestAbs2(t *testing.T) {
	f := func(x int) bool {
		return Abs2(x) >= 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}


// TestAbs3 out of the box, random testing typically does not find the bug even for int8
func TestAbs3(t *testing.T) {
	f := func(x int8) bool {
		return Abs3(x) >= 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// TestAbs3b performs a complete test - here we find the bug
func TestAbs3b(t *testing.T) {
	f := func(x int) bool {
		return Abs3(int8(x)) >= 0
	}
    for i := math.MinInt8 ; i <= math.MinInt8 ; i++ {
	   if !f(i) {
		t.Error("test failed")
       }
	}
}

// TestAbs3c adjusts the number of random test cases to the mean number of test cases to find the bug
func TestAbs3c(t *testing.T) {
	f := func(x int8) bool {
		return Abs3(x) >= 0
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 256}); err != nil {
		t.Error(err)
	}
}

// TestAbs3e uses a combination of extreme and random values
func TestAbs3e(t *testing.T) {
	f := func(x int8) bool {
		return Abs3(x) >= 0
	}
	if err := quick.Check(f, &quick.Config{Values: extremeValues.ExtremeValues(f)}); err != nil {
		t.Error(err)
	}
}

// TestAbs2e uses a combination of extreme and random values
func TestAbs2e(t *testing.T) {
	f := func(x int) bool {
		return Abs2(x) >= 0
	}
	if err := quick.Check(f, &quick.Config{Values: extremeValues.ExtremeValues(f)}); err != nil {
		t.Error(err)
	}
}
