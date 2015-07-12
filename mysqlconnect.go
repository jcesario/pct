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

	//statement, err := db.Prepare("SELECT * FROM QUERY_RESPONSE_TIME;")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := statement.Query()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//f	columns, err := rows.Columns()

	//	fmt.Println(columns)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rowcount := 1

	mysqloutput := make(map[string]mysqldata)

	for rows.Next() {
		derp := rows.Scan(&TIME, &COUNT, &TOTAL)
		if derp != nil {
			//	fmt.Println(derp)
			os.Exit(1)
		}
		/*


			fmt.Println("WHOOOOO", TIME, COUNT, TOTAL)

				fmt.Println(columns)
				fmt.Println(reflect.TypeOf(columns))

				fmt.Println("type of ", reflect.TypeOf(TIME))
				fmt.Println("type of ", reflect.TypeOf(COUNT))
				fmt.Println("type of ", reflect.TypeOf(TOTAL))

		*/

		/*

		   func (f *mysqldata) SetTime(time float64) {
		   	f.time = time
		   }

		   func (f *mysqldata) SetCount(count int) {
		   	f.count = count
		   }

		   func (f *mysqldata) SetTotal(total float64) {
		   	f.total = total

		   }

		*/

		mysqlarray := mysqldata{time: TIME, count: COUNT, total: TOTAL}
		//		fmt.Println("ohhh mannn ", mysqlarray)

		// mysqloutput := map[string]mysqldata{"TIME": }

		/*
			type mysqldata struct {
				time  float32
				count int
				total float32
			}
		*/

		poop := "ROW " + strconv.Itoa(rowcount)

		mysqloutput[poop] = mysqlarray

		fmt.Println(mysqloutput)

		rowcount++

		/*

			values := make([]interface{}, len(columns))
			fmt.Println(values)
			fmt.Println(len(values))

			scanArgs := make([]interface{}, len(values))

			fmt.Println(scanArgs)
			for i := range values {
				scanArgs[i] = &values[i]
			}

			fmt.Println("1", values)
			fmt.Println("2", len(values))
			fmt.Println("3", scanArgs)

			fmt.Println("4", scanArgs)

			//returnslice := make([]mysqldata, 0)

			// Fetch rows
			for rows.Next() {
				/*
					var time1 float64
					var count1 int
					var total1 float64
					err = rows.Scan(&time1, &count1, &total1)
					checkErr(err)
					fmt.Println(time1)
					fmt.Println(count1)
					fmt.Println(total1)

				// FUCK YOU  RawBytes from data

				// var time float32
				// var count int
				// var total float32

				//err = rows.Scan(&time, &count, &total)

				//		fmt.Println(time, count, total)

				fmt.Println("rows things", rows)

				err = rows.Scan(scanArgs...)

				fmt.Println("err things", err)

				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}

				//fmt.Println(err)
				record := make(map[string]interface{})

				//		record := make(map[string]mysqldata)

				// rowreturn := mysqldata{}
				// Now do something with the data.
				// Here we just print each column as a string.
				//		var value float32
				for i, col := range values {
					if col != nil {
						fmt.Printf("\n%s: type= %s\n", columns[i], reflect.TypeOf(col))
					}
					switch t := col.(type) {

					default:
						fmt.Printf("Unexpected type %T\n", t)
					case bool:
						fmt.Printf("bool\n")
						record[columns[i]] = col.(bool)
					case int:
						fmt.Printf("int\n")
						record[columns[i]] = col.(int)
					case int64:
						fmt.Printf("int64\n")
						record[columns[i]] = col.(int64)
						herp := col.(int64)
						println("herp", herp)
					case float64:
						fmt.Printf("float64\n")
						record[columns[i]] = col.(float64)
					case string:
						fmt.Printf("string\n")
						record[columns[i]] = col.(string)
					case []byte:
						fmt.Printf("uint8\n")
						record[columns[i]] = col.([]byte)
						derp := col.([]byte)
						println("float", derp)

					case time.Time:
						//				echo "derp"

						// record[columns[i]] = col.(string)

						//			fmt.Println("columns", i)
						//			fmt.Println("col", col)
						//			//case []byte: // -- all cases go HERE!
						//	fmt.Printf("[]byte\n")
						//	record[columns[i]] = string(col.([]byte))

						//				if record[columns[i]].(type) == int64 {
						//					continue
						//				} else {
						//					derp := col.([]byte)
						//				}

					}

					//		fmt.Println(columns)
					//			fmt.Println(f)
					//			fmt.Println(derp)

					//float64me := Float64frombytes(derp)
					//fmt.Printf(float64me)
					//			fmt.Println("derp")

				}

				/*			// Here we can check if the value is nil (NULL value)
					//			if col == nil {
					//				continue
					//value = "NULL"
					//			} else {
					//				fmt.Println(col)
					//				value = col
					//value = string(col)
					//			println(&scanArgs)
					//if i == 1 {
						//				a, _ := strconv.ParseFloat(value, 64)

					//	rowreturn.SetTime(time)
					}
					//if i == 2 {
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
				//}
				// fmt.Println(returnslice)


		*/

		if err = rows.Err(); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
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
