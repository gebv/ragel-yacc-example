gen:
	ragel -Z -G2 -o strparam_lex.go strparam.rl
	goyacc -o strparam.go strparam.y
	gofmt -w strparam.go strparam_lex.go
