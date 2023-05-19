package main

func main() {

}

func primo(slice *[]int, n int) {
	for i := 2; i <= n; i++ {
		isPrimo := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrimo = false
			}
		}
	}
}
