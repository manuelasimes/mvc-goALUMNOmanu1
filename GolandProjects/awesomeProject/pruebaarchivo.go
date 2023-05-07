package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File { //lo crea y devuelve puntero a os file
	fmt.Println("creating")

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: &v\n", err)
		os.Exit(1)
	}

}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func main() {
	f := createFile("C:/Users/Manuela/Documents/TERCER AÑO INGENIERIA/ING SW I/archivo.txt")
	defer closeFile(f)
	writeFile(f)
}
