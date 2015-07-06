package main

import "fmt"
import "sort"

type bucket struct {
	time  float64
	count int
	total float64
}

type histogram []bucket

// Implement Sort
func (h histogram) Len() int      { return len(h) }
func (h histogram) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h histogram) Less(i, j int) bool {
	return h[i].time < h[j].time
}

// Sum the total entries in the histogram
func qrtentries(h []bucket) int {
	total := 0

	for _, v := range h {
		total += v.count
	}

	return total
}

func percentile(h []bucket, p int) float64 {
	var p_ix float64
	var cur_pctl float64
	var total int
	var pctl float64

	// Find the total number of entries in histogram
	total = qrtentries(h)

	fmt.Printf("total: %v\n", total)

	// Multiply the total number of values in the data set by the percentile, which will give you the index.
	p_ix = (float64(total) * float64(p)) * .01

	fmt.Printf("index: %f\n", p_ix)

	// Order all of the values in the data set in ascending order (least to greatest).
	var times []float64

	for _, v := range h {
		times = append(times, v.time)
	}

	var count []int

	for _, v := range h {
		count = append(count, v.count)
	}

	sort.Ints(count)
	sort.Float64s(times)
	fmt.Println(times)

	// Find the tgt percentile
	// If p_ix equal to tgt pctl then average current and next value together
	for i, _ := range h {
		if float64(i) > p_ix {
			pctl = times[i-1]
			fmt.Printf("cur_pctl: %v\n pctl: %v\n", cur_pctl, pctl)
			break
		}
	}

	return pctl
}

func main() {
	var hist []bucket

	/*
		    Make connection to mysql
			  Return histogram

		    75 90 95 99 999 9999
		    percentile(hist, )
	*/
	// p := [5]float32{75, 90, 95, 99, 99.9}

	hist = append(hist, bucket{time: 85, count: 1, total: 85})
	hist = append(hist, bucket{time: 34, count: 1, total: 34})
	hist = append(hist, bucket{time: 42, count: 1, total: 42})
	hist = append(hist, bucket{time: 51, count: 1, total: 51})
	hist = append(hist, bucket{time: 84, count: 1, total: 84})
	hist = append(hist, bucket{time: 86, count: 1, total: 86})
	hist = append(hist, bucket{time: 78, count: 1, total: 78})
	hist = append(hist, bucket{time: 85, count: 1, total: 85})
	hist = append(hist, bucket{time: 87, count: 1, total: 87})
	hist = append(hist, bucket{time: 69, count: 1, total: 69})
	hist = append(hist, bucket{time: 94, count: 1, total: 94})
	hist = append(hist, bucket{time: 74, count: 1, total: 74})
	hist = append(hist, bucket{time: 65, count: 1, total: 65})
	hist = append(hist, bucket{time: 56, count: 1, total: 56})
	hist = append(hist, bucket{time: 97, count: 1, total: 97})
	/*
		hist = append(hist, row{time: 0.000001, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.000003, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.000007, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.000015, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.000030, count: 1656, total: 0.044709})
		hist = append(hist, row{time: 0.000061, count: 24841, total: 1.067347})
		hist = append(hist, row{time: 0.000122, count: 11003, total: 0.862894})
		hist = append(hist, row{time: 0.000244, count: 1435, total: 0.226209})
		hist = append(hist, row{time: 0.000488, count: 287, total: 0.093543})
		hist = append(hist, row{time: 0.000976, count: 244, total: 0.181830})
		hist = append(hist, row{time: 0.001953, count: 299, total: 0.411185})
		hist = append(hist, row{time: 0.003906, count: 176, total: 0.566051})
		hist = append(hist, row{time: 0.007812, count: 157, total: 0.759050})
		hist = append(hist, row{time: 0.015625, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.031250, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.062500, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.125000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.250000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 0.500000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 1.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 2.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 4.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 8.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 16.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 32.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 64.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 128.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 256.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 512.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 1024.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 2048.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 4096.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 8192.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 16384.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 32768.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 65536.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 131072.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 262144.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 524288.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 1048576.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 2097152.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 4194304.000000, count: 0, total: 0.000000})
		hist = append(hist, row{time: 8388608.000000, count: 0, total: 0.000000})
	*/

	// sysbench/sysbench --test=sysbench/tests/db/oltp.lua --mysql-user=root prepare
	// sysbench/sysbench --test=sysbench/tests/db/oltp.lua --mysql-user=root run
	/*
		for _, pct := range p {
			// percentile(hist, p)
			fmt.Println(pct)
		}

		for k, v := range hist {
			fmt.Printf("%v: %+v\n", k, v.count)
		}
	*/

	fmt.Printf("orig: %v\n", hist)
	sort.Sort(histogram(hist))
	fmt.Printf("sorted: %v\n", hist)
	// percentile(hist, 90)
}
