package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type obj struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (o obj) ToJSON() (string, error) {
	data, err := json.Marshal(o)
	if err == nil {
		return string(data), nil
	}
	return "", err
}

func (o obj) toString() (string, error) {
	return fmt.Sprintf("%s - %d ", o.Name, o.Age), nil
}

func tryDefer() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func tryDefer1() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i)
	}
}

func tryDefer2() (i int) {
	defer func() { i++ }()
	return 1
}

func myPanic() {
	fmt.Println("call panic here ")
	panic("I call this panic he he he, the program will stop")
}

func callPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("I am recover, trust me, you will be saved :D :D")
		}
	}()
	myPanic()
}

func caller() {
	fmt.Println("I am going to call panic ")
	callPanic()
	fmt.Println("After I call panic normal ")
}

func main() {
	o := obj{"tri", 24}
	data, err := json.Marshal(o)
	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
	s := "10"
	if _, err := strconv.Atoi(s); err == nil {
		fmt.Println(s + " can parse ")
	}
	fmt.Println(o.toString())

	db, err := sql.Open("mysql", "root:pass@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, e := db.Query("show databases;")
	if e != nil {
		panic(e.Error())
	}
	columns, e := result.Columns()
	if e != nil {
		panic(e.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for result.Next() {
		err = result.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
	}
	if err = result.Err(); err != nil {
		panic(err.Error())
	}

	tryDefer()
	fmt.Println("=================")
	tryDefer1()
	fmt.Println("=================")
	fmt.Println(tryDefer2())
	fmt.Println("=================")
	caller()

	var dict map[string]string
	dict = make(map[string]string)

	dict["hello"] = "xin chao"
	fmt.Println(dict["hello"])
}
