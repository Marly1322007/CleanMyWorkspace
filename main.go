package main

import (
	"CleanMyWorkspace/Clean"
	"fmt"
)

func GenerateWorkSpace() *[][]*string {
	workspace := [][]*string{
		{ptr("Canette de soda"), nil, ptr("Mouchoir")},
		{nil, ptr("Paquet de confiserie"), nil},
		{ptr("Emballage sandwich"), nil, nil},
	}
	return &workspace
}

func ptr(s string) *string {
	return &s
}

func main() {
	workspace := GenerateWorkSpace()

	fmt.Println("Workspace avant nettoyage :")
	for _, row := range *workspace {
		fmt.Println(row)
	}

	Clean.CleanWorkSpace(workspace)

	fmt.Println("\nWorkspace après nettoyage :")
	for _, row := range *workspace {
		fmt.Println(row)
	}
}
