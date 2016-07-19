package main

import (
	"encoding/json"
	"fmt"
)
/**
在上面的代码中 Age 明明是 int 解析后成了 float64。
这是因为 Go 中规定，
Json 中的布尔值会被解析为布尔值，
Json 中的所有数字(整型，浮点型)将被解析为 float64，
Json 中的string，被解析为 string 类型，
Json 中的数组被解析为 interface{}数组，
Json 中的空值解为nil。
*/
func main() {
	str := `{"userName":"zhangThree","Age":19}`
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m)
	for k, v := range m {
		switch v.(type) {
		case float64:
			fmt.Println(k, "  int ,value:", v)
		case string:
			fmt.Println(k, "   string ,value:", v)
		default:
			fmt.Println(k, "unkown type")
		}
	}
}
