package main

import (
	"CleanMyWorkspace/Clean"
	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {
	// Générer le workspace avec leur fonction
	workspace := CleanMyWorkspace.GenererateWorkSpace()

	fmt.Println("Workspace avant nettoyage :")
	displayWorkspace(workspace)

	// Nettoyer le workspace
	cleanedWorkspace := Clean.CleanWorkSpace(workspace)

	fmt.Println("\nWorkspace après nettoyage :")
	displayWorkspace(cleanedWorkspace)
}

func displayWorkspace(workspace *[][]*string) {
	for _, row := range *workspace {
		fmt.Print("|")
		for _, item := range row {
			if item == nil {
				fmt.Print("nil|")
			} else {
				fmt.Print(*item + "|")
			}
		}
		fmt.Println()
	}
}
