package main;

import (
	"fmt"
	"learn-1/lib"
)

var globalVar int = 3;

const constVar int = 4;

func main() {

	fmt.Println(lib.Version);
	fmt.Println(lib.GetEnv());
	var a int;

	a = 1;
	
	fmt.Println(a);

	var b = 2;

	fmt.Println(b);

	c := 3;

	fmt.Println(c);

	var (
		d int = 4;
		e int = 5;
	)

	fmt.Println(d);
	fmt.Println(e);

	f, g, h := 6, 7, 8;

	fmt.Println(f);
	fmt.Println(g);
	fmt.Println(h);

	fmt.Println(hello)
}