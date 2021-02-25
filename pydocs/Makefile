gen:
	ragel -Z -G2 -o pydocs_lex.go pydocs.rl
	goyacc -o pydocs.go pydocs.y
	gofmt -w pydocs.go pydocs_lex.go
