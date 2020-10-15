package main

import "fmt"

func main()  {
	fmt.Println("Hola Mundo!") // eliminar esta línea
}import "fmt"

func main() {
	var tam int
	var num int
	var aux int
	
	//fmt.Println("Tamaño del slice: ")
	fmt.Scan(&tam)

	s := make([]int, 0, tam)

	for i := 0 ; i < tam ; i++ {
		if i < tam {
			//fmt.Println("ingrese un numero: ")
			fmt.Scan(&num)
			s = append(s,num);
			aux += num
		}		
	}
	fmt.Println(aux)
	fmt.Println(len(s), cap(s))

	for i, v := range s {
		fmt.Println(i,v)
	}

	
}