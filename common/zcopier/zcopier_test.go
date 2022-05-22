package zcopier

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Test struct {
	Field1 time.Time
	Field2 *time.Time
	Field3 *time.Time
}

type Res struct {
	Field1 string
	Field2 string
	Field3 string
}
type Test1 struct {
	Field1 time.Time
	Field2 *time.Time
	Field3 string
	Field4 *time.Time
	Field5 int
	Field6 int
	Field7 string
}

type Res1 struct {
	Field1 string
	Field2 string
	Field3 string
	Field4 string
	Field5 string
	Field6 string
	Field7 int
}

func TestCopy(t *testing.T) {

	now := time.Now()
	a := Test{
		Field1: now,
		Field2: &now,
		Field3: nil,
	}
	var b Res

	Copy(&b, a)
	t.Log(b)
	if b.Field1 != now.Format(timex.DateTimeLayout) {
		t.Error("field1 err")
	}
	if b.Field2 != now.Format(timex.DateTimeLayout) {
		t.Error("field2 err")
	}
	if b.Field3 != "" {
		t.Error("field3 err")
	}

	var res1 = Res1{
		Field1: now.Format(timex.DateTimeLayout),
		Field2: now.Format(timex.DateTimeLayout),
		Field3: "这是一个抄抄写写",
		Field4: "",
		Field5: "10",
		Field6: "我晓得",
		Field7: 222,
	}
	var test1 Test1
	Copy(&test1, res1)
	t.Log(test1)
	assert.Equal(t, test1.Field1.Format(timex.DateTimeLayout), now.Format(timex.DateTimeLayout))
	assert.Equal(t, test1.Field2.Format(timex.DateTimeLayout), now.Format(timex.DateTimeLayout))
	assert.Equal(t, test1.Field3, "这是一个抄抄写写")
	assert.Nil(t, test1.Field4)
	assert.Equal(t, test1.Field5, 10)
	assert.Equal(t, test1.Field7, "222")
}
