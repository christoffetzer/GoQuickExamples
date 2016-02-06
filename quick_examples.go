/* (C) Christof Fetzer, 2016
*
* simple examples to demonstrate advantages / disadvantages of random testing
*
* test by executing: "go test -v"
*/

/*
package quick_examples introduces some simple examples to explain the random testing
*/
package quick_examples


// Abs1 is a buggy absolute value function
func Abs1(v int) int {
    if v < 0 {
        return v
    }
    return v
}

// Abs2 is also buggy absolute value function - maybe less obvious
func Abs2(v int) int {
    if v < 0 {
        return -v
    }
    return v
}

// Abs3 is also buggy absolute value function - restricted value range to help random testing
func Abs3(v int8) int8 {
    if v < 0 {
        return -v
    }
    return v
}
