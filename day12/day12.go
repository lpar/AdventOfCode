package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	inf, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer inf.Close()
	fmt.Println(part1(inf))
	inf.Seek(0, 0)
	fmt.Println(part2(inf))
}

func part1(input *os.File) float64 {
	total := 0.0
	dec := json.NewDecoder(input)
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if reflect.TypeOf(t).Kind() == reflect.Float64 {
			total += reflect.ValueOf(t).Float()
		}
	}
	return total
}

func part2(input *os.File) float64 {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}
	var f interface{}
	err = json.Unmarshal(data, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	return totalForObject(m)
}

func totalForObject(m map[string]interface{}) (total float64) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			if vv == "red" {
				return 0.0
			}
		case float64:
			total += vv
		case []interface{}:
			total += totalForArray(vv)
		case map[string]interface{}:
			total += totalForObject(vv)
		}
	}
	return total
}

func totalForArray(m []interface{}) (total float64) {
	for k, v := range m {
		switch vv := v.(type) {
		case float64:
			total += vv
		case []interface{}:
			total += totalForArray(vv)
		case map[string]interface{}:
			total += totalForObject(vv)
		}
	}
	return total
}
