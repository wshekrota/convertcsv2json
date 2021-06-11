package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

/*
	Function: input the text from a csv file and return it as a json string
	This output will be valid json that can be decoded
	requirements:
		csv has a column header line
		it does not use special characters(they will be excluded)
		marshal take []byte as input so this outputs byte
*/

func cvtc2j(list []string) []byte {

	first := true
	words := []string{}
	result := "["
	for r := range list {
		if first {
			words = strings.Split(list[r], ",")
			first = false
			continue
		}

		result += "{"
		each := strings.Split(list[r], ",")
		elen := len(each)
		for i := 0; i < elen; i++ {
			result += fmt.Sprintf("%q:%s", words[i], each[i])
			if i<elen-1 {
			result += ","
			}
		}
		result += "}"
		if r != len(list)-1 {
			result += ","
		} else {
			result += "]"
		}
	}
	return []byte(result)
}


// Example code
func main() {

// prereq input for cvtc2j is []string from a file
// so this requires you to read it line by line not as byte
// READFILE

	var dat = []string{}
	file, err := os.Open("./test.csv")
	if err != nil {
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}
	file.Close()

    
    var buf interface{}

    // decode .. here is the beef call to cvtc2j gets decoded
    j := cvtc2j(dat)

    fmt.Println(json.Unmarshal(j,&buf))

    // interface contains []map[string]int but cannot index an interface
    z := buf.([]interface {})
    fmt.Println(z[3]) // show map at 4th index or range aas below
//    for i,j := range z {
//        fmt.Println(i,"=>",j)
//    } 
}
