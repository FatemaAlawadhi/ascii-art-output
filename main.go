package main

import (
	"ascii-art-output/pkg"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the input
	input := os.Args
	if len(input) > 4 {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
	} else if len(input) < 2 {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
	} else {
		var Text string
		var TextArray []string
		var Banner []string
		var FinalText []string
		var FileName string
		var TextString string

		for i := 0; i < len(input); i++ {
			if strings.HasPrefix(input[i], "--output=") {
				if i != 1 || (i==1 && len(input)==2){
					fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
					return
				}
			}
			if input[i] == "shadow" || input[i] == "standard" || input[i] == "thinkertoy" {
				if (len(input) == 3 && i != 2) || (len(input) == 4 && i != 3) || len(input) == 2 {
					fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
					return
				}
			}
		}

		if len(input) == 4 {
			Text = string(input[2])
			TextArray = strings.Split(Text, "\\n")
			Banner = pkg.Banner(input[3])
			FinalText = pkg.PrintString(Banner, TextArray)
			if strings.HasPrefix(input[1], "--output=") {
				FileName = strings.TrimPrefix(input[1], "--output=")
				pkg.CreateFile(FileName, FinalText)
			}
		} else if len(input) == 3 {
			if strings.HasPrefix(input[1], "--output=") {
				Text = string(input[2])
				TextArray = strings.Split(Text, "\\n")
				Banner = pkg.Banner("No Banner")
				FinalText = pkg.PrintString(Banner, TextArray)
				FileName = strings.TrimPrefix(input[1], "--output=")
				pkg.CreateFile(FileName, FinalText)
			} else {
				Text = string(input[1])
				TextArray = strings.Split(Text, "\\n")
				Banner = pkg.Banner(input[2])
				FinalText = pkg.PrintString(Banner, TextArray)
				TextString = strings.Join(FinalText, "")
				fmt.Print(TextString)
			}
		} else if len(input) == 2 {
			Text = string(input[1])
			TextArray = strings.Split(Text, "\\n")
			Banner = pkg.Banner("No Banner")
			FinalText = pkg.PrintString(Banner, TextArray)
			TextString = strings.Join(FinalText, "")
			fmt.Print(TextString)
		}

	}
}
