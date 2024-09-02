package utils

import (
	"context"
	"reflect"
)

type (
	WhereBuilder struct {
		whereCount int
		where      string
		whereArgs  []interface{}
	}

	JoinBuilder struct {
		join     string
		joinArgs []interface{}
	}

	JoinConditionBuilder struct {
		conditionCount int
		condition      string
		conditionArgs  []interface{}
	}
)

func In_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	default:
		panic("unexpected reflect.Kind")
	}

	return
}

func Contains(val interface{}, array interface{}) bool {
	exists, _ := In_array(val, array)
	return exists
}

func Request(ctx context.Context, key string) string {
	val := ctx.Value("request." + key)
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case string:
		return v
	default:
		return ""
	}
}

func NewWhereBuilder() *WhereBuilder {
	return &WhereBuilder{
		whereCount: 0,
		where:      "",
		whereArgs:  make([]interface{}, 0),
	}
}

func (wb *WhereBuilder) And(where string, args ...interface{}) {
	if wb.whereCount != 0 {
		wb.where += " AND"
	}

	wb.where += " " + where
	wb.whereArgs = append(wb.whereArgs, args...)
	wb.whereCount++
}

func (wb *WhereBuilder) Or(where string, args ...interface{}) {
	if wb.whereCount != 0 {
		wb.where += " OR"
	}

	wb.where += " " + where
	wb.whereArgs = append(wb.whereArgs, args...)
	wb.whereCount++
}

func (wb *WhereBuilder) Get() (string, []interface{}) {
	return wb.where, wb.whereArgs
}

func NewJoinConditionBuilder() *JoinConditionBuilder {
	return &JoinConditionBuilder{
		conditionCount: 0,
		condition:      "",
		conditionArgs:  make([]interface{}, 0),
	}
}

func (jcb *JoinConditionBuilder) And(condition string, args ...interface{}) {
	if jcb.conditionCount != 0 {
		jcb.condition += " AND"
	}

	jcb.condition += " " + condition
	jcb.conditionArgs = append(jcb.conditionArgs, args...)
	jcb.conditionCount++
}

func (jcb *JoinConditionBuilder) Or(condition string, args ...interface{}) {
	if jcb.conditionCount != 0 {
		jcb.condition += " OR"
	}

	jcb.condition += " " + condition
	jcb.conditionArgs = append(jcb.conditionArgs, args...)
	jcb.conditionCount++
}

func (jcb *JoinConditionBuilder) Get() (string, []interface{}) {
	return jcb.condition, jcb.conditionArgs
}

func NewJoinBuilder() *JoinBuilder {
	return &JoinBuilder{
		join: "",
	}
}

func (jb *JoinBuilder) InnerJoin(table string, condition string, args ...interface{}) {
	jb.join += " INNER JOIN"
	jb.join += " " + table + " ON " + condition
	jb.joinArgs = append(jb.joinArgs, args...)
}

func (jb *JoinBuilder) LeftJoin(table string, condition string, args ...interface{}) {
	jb.join += " LEFT JOIN"
	jb.join += " " + table + " ON " + condition
	jb.joinArgs = append(jb.joinArgs, args...)
}

func (jb *JoinBuilder) RightJoin(table string, condition string, args ...interface{}) {
	jb.join += " RIGHT JOIN"
	jb.join += " " + table + " ON " + condition
	jb.joinArgs = append(jb.joinArgs, args...)
}

func (jb *JoinBuilder) Get() (string, []interface{}) {
	return jb.join, jb.joinArgs
}
