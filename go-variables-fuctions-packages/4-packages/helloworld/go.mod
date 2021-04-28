module helloworld

go 1.16

require (
	github.com/myuser/calculator v0.0.0
	rsc.io/quote v1.5.2
)

replace github.com/myuser/calculator => ../calculator
