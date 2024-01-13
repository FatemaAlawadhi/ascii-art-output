package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
	"bufio"
	"fmt"
)

func TestMain(t *testing.T) {
	exec.Command("go" , "run" , "main.go", "--output=banner.txt", "hello", "standard" )
	
	var Output []string
	var OutputStr string
	file, err := os.Open("banner.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Output = append(Output, scanner.Text())
	}
	OutputStr = strings.Join(Output, "\n")

	expected :=                                                                                         
` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`

	if OutputStr != expected {
		t.Errorf("Expected %q, but got %q" , expected , OutputStr)
	}

}