package zsqlx

import (
	"fmt"
	"strings"
)

type builderItem struct {
	builder *Builder
	op      string
}
type BuilderGroup struct {
	builders []*builderItem
}

func NewBuilderGroup() *BuilderGroup {
	return &BuilderGroup{
		builders: make([]*builderItem, 0),
	}
}

func (b *BuilderGroup) Group(builder *Builder, ops ...string) *BuilderGroup {
	op := OpAnd
	if len(ops) > 0 {
		op = ops[0]
	}
	item := &builderItem{
		builder: builder,
		op:      op,
	}
	b.builders = append(b.builders, item)
	return b
}

func (b *BuilderGroup) End() (string, []interface{}) {
	var where []string
	var args []interface{}
	for i, v := range b.builders {
		e, a := v.builder.End()
		if i == 0 {
			where = append(where, e)
		} else {
			where = append(where, v.op, e)
		}
		args = append(args, a...)
	}
	return fmt.Sprintf("(%s)", strings.Join(where, " ")), args
}
