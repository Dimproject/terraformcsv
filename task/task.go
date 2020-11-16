package task

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type VM struct {
	vmName   string
	template string
	ip       string
	cpu      string
	memory   string
}

func Process(srcFile, dstFile string) (string, string) {

	records, err := readData(srcFile)

	if err != nil {
		log.Fatal(err)
	}

	resultfile, err := os.Create(dstFile + ".csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer resultfile.Close()

	for _, record := range records {
		vm := VM{
			vmName:   record[0],
			template: record[1],
			ip:       record[2],
			cpu:      record[3],
			memory:   record[5],
		}
		if vm.template == dstFile {
			fmt.Printf("%s %s %s %s %s\n", vm.vmName, vm.template, vm.ip, vm.cpu, vm.memory)

			cellSlice := []string{vm.vmName, vm.template, vm.ip, vm.cpu, vm.memory}
			csvWriter := csv.NewWriter(resultfile)
			strWrite := [][]string{cellSlice}
			csvWriter.WriteAll(strWrite)
			csvWriter.Flush()
		}

	}
	return srcFile, dstFile
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
