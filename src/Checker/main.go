package main

import (
	"fmt"
	generator "github.com/Chloe-Woahie/Roblox-Namesnipe-Finder/generator"
	validator "github.com/Chloe-Woahie/Roblox-Namesnipe-Finder/validator"
	inputLogic "github.com/Chloe-Woahie/Roblox-Namesnipe-Finder/inputLogic"
)

func fullScanCycle() {
	fmt.Println("Enter Word To Check: ")
    var rawName string
    fmt.Scanln(&rawName)
	combos := generator.ReturnCombos(rawName)
	validCombos := validator.ReturnValidCombos(combos)
	fmt.Println("Scan Completed!")
	fmt.Println("Valid Combos:")
	fmt.Println(validCombos)
}

func main() {
	for {
		fullScanCycle()
		fmt.Println("Search again? (y/n):")
		if inputLogic.AskForConfirmation() {
			continue
		} else {
			break
		}
    }
}