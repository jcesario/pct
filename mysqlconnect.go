package mysqlconnect 

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sort"
	"os"
)

// mysql struct holds the histogram data itself per histo interval time unit / row / bucket
type mysqldata struct {
	time  float64
	count int64
	total float64
}

// histogram format of mysqldata struct, row based on bucket
type MysqlHisto []mysqldata

// NewMysqlQrtBucket Public way to return a QRT bucket to be appended to a Histogram
func NewMysqlHisto(time float64, count int64, total float64) MysqlHisto {
	return MysqlHisto{time, count, total}

// Sort for mysql histogram format, in three methods, length, swap, less for Bubble Sort 
func (m MysqlHisto) Len() int {
	return len(h)
}
func (m MysqlHisto) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (m MysqlHisto) Less(i, j int) bool {
	return h[i].time < h[j].time
}
// Count for MysqlHisto
func (m MysqlHisto) Count() int64 {
	var total int64
	total = 0
	for _, v := range h {
		total += v.count
	}
	return total
}

func mysqlconnect(uname string, host string, port string) (m MysqlHisto) {
	var (
		TIME  float64
		COUNT int64
		TOTAL float64
	)

	db, err := sql.Open("mysql", "uname:@tcp(host:port)/information_schema")
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

	for rows.Next() {
		derp := rows.Scan(&TIME, &COUNT, &TOTAL)
		if derp != nil {
			os.Exit(1)
		}

		mysqlarray := mysqldata{time: TIME, count: COUNT, total: TOTAL}

		mysqloutput 

		fmt.Println(mysqloutput)

		if err = rows.Err(); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}
