module github.com/shengyanli1982/regula/examples/lazy

go 1.19

replace (
	github.com/shengyanli1982/regula => ../../
	github.com/shengyanli1982/regula/contrib/lazy => ../../contrib/lazy
)

require github.com/shengyanli1982/regula/contrib/lazy v0.0.0-00010101000000-000000000000

require (
	github.com/shengyanli1982/karta v0.1.11 // indirect
	github.com/shengyanli1982/regula v0.0.0-00010101000000-000000000000 // indirect
	github.com/shengyanli1982/workqueue v0.1.12 // indirect
	golang.org/x/time v0.5.0 // indirect
)
