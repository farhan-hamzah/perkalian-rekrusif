package main
import "fmt"

func main(){
	var a, b, hasil int
	fmt.Scan(&a, &b)
	hasil = perkalian(a, b)
	fmt.Print(hasil)
}
func perkalian(a, b int)int{
	if b == 0{
		return 0
	}else if b == 1{
		return a
	}else{
		return a + perkalian(a, b-1)
	}
}