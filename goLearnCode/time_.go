package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t, "\n", time.Now())
	t = time.Now()
	fmt.Printf("%s 今年还已经过了%d天,还有%d天\n", t.Month().String(), t.YearDay(), time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()-t.YearDay())
	fmt.Println(t.Format("2006年01月2日 15:04:05 ")) //格式化方法挺新奇的.
}
