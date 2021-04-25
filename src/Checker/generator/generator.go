package generator

import (
	//"fmt"
	sPlus "github.com/Chloe-Woahie/Roblox-Namesnipe-Finder/generator/stringsPlus"
)

func ReturnCombos(rawName string) []string {
	combos := []string{}
	combos = append(combos, rawName)
	if len(combos) <= 19 {
		combos = append(combos, rawName + "s")
	}
	for i := 1; i < (len(rawName) - 1); i++ {
		combos = append(combos, sPlus.ReplaceChar(rawName, '_', i))
		if len(combos) <= 19 {
			combos = append(combos, sPlus.InsertChar(rawName, '_', i))
		}
	}
	for i := 0; i < len(rawName); i++ {
		combos = append(combos, sPlus.ReplaceChar(rawName, 'x', i))
		if len(combos) <= 19 {
			combos = append(combos, sPlus.InsertChar(rawName, 'x', i))
		}
	}
	combos = append(combos, sPlus.CreateFlippedLICombos(rawName)...)
	return combos
}