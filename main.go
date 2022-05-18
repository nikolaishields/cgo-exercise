
//main.go
package main

/*
#include <stdio.h>

int add(int a, int b) {
  return a + b;
}

int substract(int a, int b) {
  return a - b;
}

int multiply(int a, int b) {
  return a * b;
}

int divide(int a, int b) {
  return a / b;
}

int power(int a, int b) {
  int result = 1;
    while(b) { result *= a; b--; }
    return result;
}

int mod(int a, int b) {
    return a % b;
}
*/
import "C"

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	add := flag.Bool("add", false, "Add")
	flag.Parse()

	if len(os.Args) != 4 {
		log.Fatalf("Missing operands\nUsage: %s {-add|-sub|-mul|-div|-mod|-pow} NUM NUM\n", os.Args[0])
	}

	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	if *add {
		//Call to int function with two params
		res, err := doAdd(a, b)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Sum of %d + %d is %d\n", a, b, res)
	}
}

// add also is a C binding to make a sum. As before it returns a result and
// an error. Look that we had to pass the Int values to C.int values before using
// the function and cast the result back to a Go int value
func doAdd(a, b int) (int, error) {
	//Convert Go ints to C ints
	aC := C.int(a)
	bC := C.int(b)

	result, err := C.add(aC, bC)
	if err != nil {
		return 0, errors.New("error calling Sum function: " + err.Error())
	}

	//Convert C.int result to Go int
	res := int(result)

	return res, nil
}
