package main

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"os"
	"strconv"
)

type mysqldata struct {
	time  float64
	count int
	total float64
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (f *mysqldata) SetTime(time float64) {
	f.time = time
}

func (f *mysqldata) SetCount(count int) {
	f.count = count
}

func (f *mysqldata) SetTotal(total float64) {
	f.total = total

}

func main() {
	var (
		TIME  float64
		COUNT int
		TOTAL float64
	)

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/information_schema")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statement, err := db.Prepare("SELECT cast(TIME as decimal(13,6)) AS TIME, COUNT, cast(TOTAL as decimal(13,6)) AS TOTAL FROM QUERY_RESPONSE_TIME;")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rowcount := 1

	mysqloutput := make(map[string]mysqldata)

	for rows.Next() {
		derp := rows.Scan(&TIME, &COUNT, &TOTAL)
		if derp != nil {
			os.Exit(1)
		}

		mysqlarray := mysqldata{time: TIME, count: COUNT, total: TOTAL}

		poop := "ROW " + strconv.Itoa(rowcount)

		mysqloutput[poop] = mysqlarray

		fmt.Println(mysqloutput)

		rowcount++
		if err = rows.Err(); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}
