package main

import (
	"fmt"

	"cleanmyworkspace/Clean"

	cmw "github.com/Mentor-Paris/CleanMyWorkspace"
)

// Affiche le workspace : nil => ".", sinon affiche la valeur
func printWorkspace(ws *[][]*string) {
	if ws == nil {
		fmt.Println("(workspace nil)")
		return
	}
	for _, row := range *ws {
		for j, cell := range row {
			if cell == nil {
				fmt.Print(".")
			} else {
				fmt.Print(*cell)
			}
			if j < len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	workspace := cmw.GenererateWorkSpace()

	fmt.Println("Souk initial:")
	printWorkspace(workspace)

	cleaned := Clean.CleanWorkSpace(workspace)

	fmt.Println("\nSouk après nettoyage:")
	printWorkspace(cleaned)
}
