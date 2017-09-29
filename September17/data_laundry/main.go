package main

import "fmt"

type iface struct {
	name string
	inet string
	status string
}

func readFile(f string, i *iface) error {
	i.name = "lo0"
	i.inet = "127.0.0.1"

	return nil
}

func main() {
	i := iface{"test", "test", "up"}
	fmt.Println(i.status)
}