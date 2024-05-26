package main

import "fmt"

func puntero() {
	x := 10 // Declarar una variable de tipo int
	// Similar a var p *int = &x
	p := &x // Obtener la dirección de memoria de x y asignarla a un puntero p
	var z *int = &x

	// Imprimir las direcciones de memoria de x y p
	fmt.Printf("Dirección de x: %p\n", &x) // Dirección de x
	fmt.Printf("Dirección de p: %p\n", &p) // Dirección de p
	fmt.Printf("Dirección de z: %p\n", &z) // Dirección de z

	// Imprimiar las direcciones de memoria a las que apuntan p y z
	fmt.Println("Dirección a la que apunta x (&x):", &x) // Dirección a la que apunta x
	fmt.Println("Dirección a la que apunta p (p):", p)   // Dirección a la que apunta p
	fmt.Println("Dirección a la que apunta z (z):", z)   // Dirección a la que apunta z

	// Imprimir el valor de x y el valor al que apunta el puntero p
	fmt.Println("Valor de x:", x)  // Imprime: Valor de x: 10
	fmt.Println("Valor de p:", *p) // Imprime: Valor de p: 10
	fmt.Println("Valor de z:", *z) // Imprime: Valor de z: 10

	// Cambiar el valor de x a través del puntero p
	chValue(p, 99)

	// Cambiar el puntero p usando chValuePtr
	newX := 50
	chValuePtr(&p, &newX)

	// Imprimir el nuevo valor de x y el valor al que apunta el puntero p
	fmt.Println("Nuevo valor de x:", x)  // Imprime: Nuevo valor de x: 20
	fmt.Println("Nuevo valor de p:", *p) // Imprime: Nuevo valor de p: 20
	fmt.Println("Nuevo valor de z:", *z) // Imprime: Nuevo valor de z: 20

	chValue(&x, 33)

	fmt.Println("Nuevo valor de x:", x) // Imprime: Nuevo valor de x: 33

	p = changePointer(10)
	fmt.Println("Nuevo valor de p:", *p) // Imprime: Nuevo valor de p: 10

}

func chValue(xptr *int, val int) {
	*xptr = val
}

func chValuePtr(pptr **int, newPtr *int) {
	*pptr = newPtr
}

func changePointer(newPtr int) *int {
	return &newPtr
}
