package main

import (
	"fmt"
	"testing"
)

//Options 参数
type Options struct {
	StringOpt1 string
	StringOpt2 string
	IntOpt1    int
	IntOpt2    int
}

// Option a interface have Apply function to config
type Option interface {
	Apply(*Options)
}

//OptionFunc 配置 Options 参数的方法类型
type OptionFunc func(*Options)

// Apply 实现 Option接口中的 Appl y方法
func (of OptionFunc) Apply(o *Options) {
	of(o)
}

// InitOptions 创建 Options 实例的函数
func InitOptions(options ...Option) *Options {
	optionInstance := new(Options)
	for _, opt := range options {
		opt.Apply(optionInstance)
	}
	return optionInstance
}

// WithStringOpt1 修改参数
func WithStringOpt1(param string) Option {
	return OptionFunc(func(opt *Options) {
		opt.StringOpt1 = param
	})
}

// WithStringOpt2 修改参数
func WithStringOpt2(param string) Option {
	return OptionFunc(func(opt *Options) {
		opt.StringOpt2 = param
	})
}

// WithIntOpt1 修改参数
func WithIntOpt1(param int) Option {
	return OptionFunc(func(opt *Options) {
		opt.IntOpt1 = param
	})
}

// WithIntOpt2 修改参数
func WithIntOpt2(param int) Option {
	return OptionFunc(func(opt *Options) {
		opt.IntOpt2 = param
	})
}

func TestOptions(t *testing.T) {
	var optionList []Option
	withInt1, withInt2, withString1, withString2 := true, true, true, true
	if withInt1 {
		optionList = append(optionList, WithIntOpt1(123))
	}
	if withInt2 {
		optionList = append(optionList, WithIntOpt2(211))
	}
	if withString1 {
		optionList = append(optionList, WithStringOpt1("Opt1"))
	}
	if withString2 {
		optionList = append(optionList, WithStringOpt2("Opt2"))
	}
	opt := InitOptions(optionList...)
	fmt.Println(opt)
}
