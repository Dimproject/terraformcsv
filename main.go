package main

import (
	"flag"
	"fmt"
	"training/terraformcsv/task"
)

func main() {
	srcFile := flag.String("src", "", "CSV file input.")
	dstFile := flag.String("dst", "", "Destination file output.")
	flag.Parse()

	if len(*srcFile) <= 0 {
		fmt.Printf("Invalid file. got=%v\n", *srcFile)
	}

	if len(*dstFile) <= 0 {
		fmt.Printf("Invalid file. got=%v\n", *dstFile)
	}

	//listTemplate := []string{"TPL_WIN16_Genuine", "TPL_WIN16_Genuine"}

	//for _, template := range listTemplate {
	task.Process(*srcFile, *dstFile)
	//}

}
