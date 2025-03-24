package main

import "time"

func main() {
	i := 2
	println("Write ", i, "as")
	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")
	default:println("It's a weekday")
	}
	t:=time.Now()
	switch{
	case i>2:println("i is crazy")
	case t.Hour() < 12:
		println("before noon")
	default:println("ater noon")
	}
	whatAmI := func(i interface{}){
		switch t:=i.(type){
		case bool:
			println("i am bool")
		case int:println("Im an int")
		default:
			println("I dont know shit about", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	j := 4
	switch{
	case j>1:println("i > 1")
	case j>2:println("i > 2")
	case j>3:println("i > 3")
	}
	switch k:=3{
	case 3:println(k)
	}

}