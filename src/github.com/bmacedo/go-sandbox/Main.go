package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

// Variable scopes:
//	- upper case: global
//	- lower case: package-internal or block scoped
// There's no private scope

// Naming conventions:
// 	- camelCase
// 	- Acronyms: uppercase eg. URL
// 	- Capitalized function name

// Casting: float32(i)

const (
	a = iota
	b = iota
	c = iota
)

const (
	_ = iota
	s1
	s2
	s3
)

var waitGroup = sync.WaitGroup{}

func main() {
	i := 27
	j := strconv.Itoa(i)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(c)
	fmt.Println(s3)

	// arrays
	arr := [3]int{1, 2, 3}
	arr2 := arr // assigning an array actually copies it to a new instance
	arr2[0] = 4
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("Size: %v\n", len(arr))

	// arrays have fixed size, but array slices operate like lists
	var slice []int
	fmt.Println(slice)
	slice = append(slice, 1)
	fmt.Println(slice)

	// assigning a slice does not create a new instance, it copies by reference
	slice2 := slice
	slice2[0] = 34
	fmt.Println(slice)
	fmt.Println(slice2)

	// maps
	// the order of the keys in the map is not guaranteed!
	pop := map[string]int{
		"California": 1,
		"Texas":      2,
	}
	fmt.Println(pop)
	fmt.Println(pop["Texas"])
	delete(pop, "Texas")
	fmt.Println(pop)
	// by default any non-existing key will return zero. To check if the key exists or not, you need the "comma-ok" pattern
	test, ok := pop["Test"]
	fmt.Println(test, ok) // ok is false because the key does not exists
	// you can add a key easily
	pop["Test"] = 3
	test2, ok := pop["Test"]
	fmt.Println(test2, ok) // ok is true because we added the key

	// maps are assigned without a copy
	pop2 := pop
	pop2["New"] = 12
	fmt.Println(pop)
	fmt.Println(pop2)

	// if - <initializer ; comparison>
	if p, ok := pop["Test"]; ok {
		fmt.Printf("p = %v\n", p)
	}
	switch 1 {
	case 1:
		fmt.Println("1!")
		fallthrough
	case 2:
		fmt.Println("2!")
	case 3:
		fmt.Println("3!")
	}

	// loops
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	for k, v := range pop {
		fmt.Println(k, v)
	}

	// defer keyword
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	// pointers
	// no pointer arithmetic is allowed (like sum 4 bytes to a mem address)
	var p1 = 123
	var p2 = &p1

	fmt.Println(&p1, p2)
	fmt.Println(p1, *p2)

	p1 = 234
	fmt.Println(&p1, p2)
	fmt.Println(p1, *p2)

	*p2 = 345
	fmt.Println(&p1, p2)
	fmt.Println(p1, *p2)

	// goroutines
	runtime.GOMAXPROCS(100)
	waitGroup.Add(1)
	go sayHello()
	// time.Sleep(100 * time.Millisecond)
	waitGroup.Wait()

	// channels
	ch := make(chan int, 50)
	waitGroup.Add(2)
	go receiver(ch)
	go sender(ch)
	waitGroup.Wait()

	fmt.Println("-- exit main --")
}

func sayHello() {
	fmt.Println("Hello")
	waitGroup.Done()
}

func sender(ch chan<- int) {
	ch <- 12345
	ch <- 23456
	close(ch)
	waitGroup.Done()
}

func receiver(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
	waitGroup.Done()
}
