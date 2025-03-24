package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i += 1
	}
	for j := 0; j < 3; j++ /*j+=1*/ {
		println(j)
	}
	for i := range 3 {
		println("range", i)
	}
	t := 0
	for {
		println("loop")
		t += 1
		if t == 3 {
			break
		}
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}