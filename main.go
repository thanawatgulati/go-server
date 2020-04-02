package main

func fibSlow(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibSlow(x-1) + fibSlow(x-2)
}

func fib(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	f := make([]int, x)
	f[0] = 0
	f[1] = 1
	for i := 2; i < x; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[x-1] + f[x-2]
}

func add(x, y int) int {
	return x + y
}

func main() {

}
