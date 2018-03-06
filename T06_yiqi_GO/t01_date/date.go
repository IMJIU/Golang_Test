package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//	funcParse(layout,value string)(Time,error),用来把一个字符串转换成
	//	日期。layout 与 Format 里的格式一样，2006 代表年，01 代表月，02 代表日，15 代表时，04 代表分，05 代表秒。
	d, err := time.Parse("01-02-2006", "06-17-2013")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)

	t := time.Now()
	t2 := t.Add(24 * time.Hour) //当前时间加24小时，即明天的这个时间
	d2 := t2.Sub(t)              //t2-t1,相差24小时
	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(d2)
	//t 小 t2
	if t.Before(t2) {
		fmt.Println("t<t2")
	}
	//t2大于 t
	if t2.After(t) {
		fmt.Println("t2>t")
	}
	//判断两个时间是否相等等
	if t.Equal(t) {
		fmt.Println("t=t")
	}
}
