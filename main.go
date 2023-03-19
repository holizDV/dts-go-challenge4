package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/holizDV/dts-go-challenge4/controller"
	"github.com/holizDV/dts-go-challenge4/model"
)

func main() {
	var data model.Data

	jsonFile, err := os.Open("assets/student.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)

	number := os.Args[1]
	num, err := strconv.Atoi(number)

	if err != nil {
		fmt.Println("Args must be number")
		fmt.Printf("(Error : %v)\n", err)
		return
	}

	controller.FetchDataStudent(num, data.Students)

}
