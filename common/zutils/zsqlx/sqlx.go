package zsqlx

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

const OpAnd = " and "
const OpOr = " or "

type builder struct {
	where []string
	args  []interface{}
	op    string
}

func (b builder) Args() []interface{} {
	return b.args
}
func (b builder) ArgsString() string {
	var s []string
	for _, arg := range b.args {
		s = append(s, cast.ToString(arg))
	}
	return strings.Join(s, ",")
}

type Builder struct {
	*builder
}

func NewBuilder(ops ...string) *Builder {
	op := OpAnd
	if len(ops) > 0 {
		op = ops[0]
	}
	return &Builder{
		builder: &builder{
			op: op,
		},
	}
}

func (b *Builder) Eq(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s = ?", cond))
	b.args = append(b.args, args)
	return b
}
func (b *Builder) Neq(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s != ?", cond))
	b.args = append(b.args, args)
	return b
}

func (b *Builder) ISNULL(cond string) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s IS NULL", cond))
	return b
}
func (b *Builder) ISNotNULL(cond string) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s IS NOT NULL", cond))
	return b
}
func (b *Builder) RLike(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s like ?", cond))
	str := cast.ToString(args)
	b.args = append(b.args, str+"%")
	return b
}
func (b *Builder) LLike(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s like ?", cond))
	str := cast.ToString(args)
	b.args = append(b.args, "%"+str)
	return b
}
func (b *Builder) Like(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s like ?", cond))
	str := cast.ToString(args)
	b.args = append(b.args, "%"+str+"%")
	return b
}

func (b *Builder) NotRLike(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s not like ?", cond))
	str := cast.ToString(args)
	b.args = append(b.args, str+"%")
	return b
}
func (b *Builder) NotLLike(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s not like ?", cond))
	str := cast.ToString(args)
	b.args = append(b.args, "%"+str)
	return b
}
func (b *Builder) NotLike(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s not like ?", cond))
	str := cast.ToString(args)
	b.args = append(b.args, "%"+str+"%")
	return b
}
func (b *Builder) Between(cond string, start interface{}, end interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s between ? and ?", cond))
	b.args = append(b.args, start, end)
	return b
}

func (b *Builder) In(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s in (?)", cond))
	b.args = append(b.args, args)
	return b
}
func (b *Builder) NotIn(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s not in (?)", cond))
	b.args = append(b.args, args)
	return b
}

func (b *Builder) Gt(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s > ?", cond))
	b.args = append(b.args, args)
	return b
}

func (b *Builder) Gte(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s >= ?", cond))
	b.args = append(b.args, args)
	return b
}
func (b *Builder) Lt(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s < ?", cond))
	b.args = append(b.args, args)
	return b
}
func (b *Builder) Lte(cond string, args interface{}) *Builder {
	b.where = append(b.where, fmt.Sprintf("%s <= ?", cond))
	b.args = append(b.args, args)
	return b
}

func (b *Builder) End() (where string, args []interface{}) {
	if len(b.where) == 0 {
		return "", b.args
	}
	return fmt.Sprintf("(%s)", strings.Join(b.where, b.op)), b.args
}

func format(where string, args []interface{}) (string, interface{}) {
	if len(args) > 0 {
		return where, args[0]
	}
	return where, nil
}
func Format(where string, args []interface{}) (string, interface{}) {
	return format(where, args)
}
func Eq(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Eq(cond, args).End())
}
func Neq(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Neq(cond, args).End())
}
func ISNULL(cond string) string {
	w, _ := NewBuilder().ISNULL(cond).End()
	return w
}
func ISNotNULL(cond string) string {
	w, _ := NewBuilder().ISNotNULL(cond).End()
	return w
}

func RLike(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().RLike(cond, args).End())
}

func LLike(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().LLike(cond, args).End())
}

func Like(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Like(cond, args).End())
}

func NotRLike(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().NotRLike(cond, args).End())
}

func NotLLike(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().NotLLike(cond, args).End())
}

func NotLike(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().NotLike(cond, args).End())
}

func Between(cond string, start interface{}, end interface{}) (string, interface{}) {
	return format(NewBuilder().Between(cond, start, end).End())
}

func In(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().In(cond, args).End())
}

func NotIn(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().NotIn(cond, args).End())
}
func Gt(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Gt(cond, args).End())
}
func Gte(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Gte(cond, args).End())
}
func Lt(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Lt(cond, args).End())
}
func Lte(cond string, args interface{}) (string, interface{}) {
	return format(NewBuilder().Lte(cond, args).End())
}

func ASC(str string) string {
	return str + " ASC"
}
func DESC(str string) string {
	return str + " DESC"
}
