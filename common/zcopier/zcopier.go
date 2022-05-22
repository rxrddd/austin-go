package zcopier

import (
	"austin-go/common/zutils/timex"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

//Copy copier.CopyWithOption default option
func Copy(toValue, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, DefOpt())
}

//DefOpt 默认copy选项
func DefOpt() copier.Option {
	return copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
		Converters: []copier.TypeConverter{
			{
				SrcType: time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(time.Time)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return s.Format(timex.DateTimeLayout), nil
				},
			},
			{
				SrcType: &time.Time{},
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					if src == nil {
						return "", nil
					}
					s, ok := src.(*time.Time)
					if !ok {
						return nil, errors.New("src type not matching")
					}
					return s.Format(timex.DateTimeLayout), nil
				},
			},
			{
				SrcType: copier.String,
				DstType: time.Time{},
				Fn: func(src interface{}) (interface{}, error) {
					parse, err := time.Parse(timex.DateTimeLayout, cast.ToString(src))
					if err != nil {
						return time.Time{}, nil
					}
					return parse, nil
				},
			},
			{
				SrcType: copier.String,
				DstType: &time.Time{},
				Fn: func(src interface{}) (interface{}, error) {
					parse, err := time.Parse(timex.DateTimeLayout, cast.ToString(src))
					if err != nil {
						return nil, nil
					}
					return &parse, nil
				},
			},
			{
				SrcType: copier.String,
				DstType: copier.Int,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(string)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return strconv.Atoi(s)
				},
			},
			{
				SrcType: copier.Int,
				DstType: copier.String,
				Fn: func(src interface{}) (interface{}, error) {
					s, ok := src.(int)

					if !ok {
						return nil, errors.New("src type not matching")
					}

					return strconv.Itoa(s), nil
				},
			},
		},
	}
}
