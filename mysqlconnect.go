package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	//	"strconv"
)

type mysqldata struct {
	time  float32
	count int
	total float32
}

func (f *mysqldata) SetTime(time float32) {
	f.time = time
}

func (f *mysqldata) SetCount(count int) {
	f.count = count
}

func (f *mysqldata) SetTotal(total float32) {
	f.total = total
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/information_schema")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement, err := db.Prepare("SELECT * FROM QUERY_RESPONSE_TIME_READ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	columns, err := rows.Columns()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	returnslice := make([]mysqldata, 0)

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data

		var time float32
		var count int
		var total float32

		//err = rows.Scan(&time, &count, &total)

		//		fmt.Println(time, count, total)

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		//fmt.Println(err)

		rowreturn := mysqldata{}
		// Now do something with the data.
		// Here we just print each column as a string.
		//		var value float32
		for i, col := range scanArgs {
			// Here we can check if the value is nil (NULL value)
			//			if col == nil {
			//				continue
			//value = "NULL"
			//			} else {
			//				fmt.Println(col)
			//				value = col
			//value = string(col)
			println(&scanArgs)
			if i == 1 {
				//				a, _ := strconv.ParseFloat(value, 64)

				rowreturn.SetTime(time)
			}
			if i == 2 {
				//				b, _ := strconv.ParseInt(value, 0, 64)
				rowreturn.SetCount(count)
			}
			if i == 3 {
				//				c, _ := strconv.ParseFloat(value, 64)
				rowreturn.SetTotal(total)
			}
		}
		returnslice = append(returnslice, rowreturn)

		rowreturn.SetTime(0)
		rowreturn.SetTotal(0)
		rowreturn.SetCount(0)
		//			fmt.Println(columns[i], ": ", value)

		//		fmt.Println("-----------------------------------")
	}
	fmt.Println(returnslice)

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

/*

	for rows.Next() {
		mysqldata dump
		rows.Scan(&dump)
		fmt.Println("Row dump:", dump)
	}

	db.Close()

*/
