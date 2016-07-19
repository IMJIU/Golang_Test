package main

import (
	"database/sql"
	"errors"
	"reflect"
	"SimpleDb"
)

type MyRows struct {
	*sql.Rows
	values map[string]interface{}
	columnName []string
}

/*根据字段名来取字段的值
name:字段名，区分大小写
value:用来接收字段值的变量，需传变量的地址,如&a
*/
func (this *MyRows) GetValue(name string, value interface{}) error {
	if this.values == nil || len(this.values) == 0 {
		return errors.New("没有调用 Next,或没有可用的行")
	}
	i, ok := this.values[name]
	if ok {
		err := SimpleDb.ConvertAssign(value, i)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("字段不存在，请注意大小写")
}

func (this *MyRows) Scan(dest ...interface{}) error {
	if this.values == nil || len(this.values) == 0 {
		return errors.New("没有调用 Next,或没有可用的行")
	}
	for i := 0; i < len(dest); i++ {
		err := SimpleDb.ConvertAssign(dest[i], this.values[this.columnName[i]])
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *MyRows) Next() bool {
	bResult := this.Rows.Next()
	if bResult {
		//如果成功，取所有列的数据到 values 里
		if this.columnName == nil || len(this.columnName) == 0 {
			this.columnName, _ = this.Rows.Columns()
		}
		if this.values == nil {
			this.values = make(map[string]interface{})
		}
		var arr []interface{}
		for i := 0; i < len(this.columnName); i++ {
			var inf interface{}
			arr = append(arr, &inf)
		}
		//将数据接收到 interface{}变量里
		this.Rows.Scan(arr...)
		for i := 0; i < len(this.columnName); i++ {
			this.values[this.columnName[i]] = reflect.ValueOf(arr[i]).Elem().Interface()
		}
	}
	return bResult
}
