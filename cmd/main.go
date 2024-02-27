package main

import (
	"fmt"

	"github.com/syhbt/fynd/pkg/provider" // Ganti dengan path package Anda
)

func main() {
	listDomain, err := provider.ListAllDomain()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", listDomain)
}
