package main

import "fmt"

type Options struct {
	StringOpt1 string
	StringOpt2 string
	IntOpt1    int
	IntOpt2    int
}

type Option func(*Options)

func InitOptions(options ...Option) *Options {
	optionInstance := new(Options)
	for _, opt := range options {
		opt(optionInstance)
	}
	return optionInstance
}

func WithStringOpt1(param string) Option {
	return func(opt *Options) {
		opt.StringOpt1 = param
	}
}

func WithStringOpt2(param string) Option {
	return func(opt *Options) {
		opt.StringOpt2 = param
	}
}

func WithIntOpt1(param int) Option {
	return func(opt *Options) {
		opt.IntOpt1 = param
	}
}

func WithIntOpt2(param int) Option {
	return func(opt *Options) {
		opt.IntOpt2 = param
	}
}

func main() {
	opt := InitOptions(WithIntOpt1(100), WithStringOpt1("asd"), WithIntOpt2(200), WithStringOpt2("opt2 field"))
	fmt.Println(opt)
}
