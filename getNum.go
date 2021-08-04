package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) != 3 {
		log.Fatal("Input Type(sfc or amd) filename")
	}
	obType := os.Args[1]
	input := os.Args[2]
	if obType == "sfc" {
		sfc(input, "sfc-list")
	} else if obType == "amd" {
		amd(input, "amd-list")
	} else {
		log.Fatal("Input Type sfc or amd!!!")
	}
}

func sfc(input, output string) {
	in, err := os.Open(input)
	if err != nil {
		log.Print(err)
		return
	}
	defer in.Close()

	out, err := os.Create(output)
	if err != nil {
		log.Print(err)
		return
	}
	defer out.Close()

	listNum := []string{}
	checkNum := make(map[string]bool)
	//listPoint := []string{}
	//checkPoint := make(map[string]bool)
	reader := bufio.NewReader(in)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		slice := strings.Split(string(line), " ")
		fmt.Println(len(slice), slice)
		for i, str := range slice {
			if i == 0 {
				if _, ok := checkNum[str]; ok != true {
					// shell script で使えるように "num" というスタイルにする
					listNum = append(listNum, "\""+str+"\"")
					checkNum[str] = true
				}
			}
		}
		fmt.Println(listNum)
	}
	out.WriteString("[" + strings.Join(listNum, " ") + "]")
}

func amd(input, output string) {
	in, err := os.Open(input)
	if err != nil {
		log.Print(err)
		return
	}
	defer in.Close()

	out, err := os.Create(output)
	if err != nil {
		log.Print(err)
		return
	}
	defer out.Close()

	reader := csv.NewReader(in)
	listNum := []string{}
	checkNum := make(map[string]bool)
	count := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if count == 0 {
			count++
			continue
		}
		num := line[1]
		if _, ok := checkNum[num]; ok != true {
			// shell script で使えるように "num" というスタイルにする
			listNum = append(listNum, "\""+num+"\"")
			checkNum[num] = true
		}
	}
	fmt.Println(listNum)
	out.WriteString("[" + strings.Join(listNum, " ") + "]")
}
