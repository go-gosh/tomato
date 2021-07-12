test:
	go test -v ./...

gen_model:
	go generate app/ent/generate.go

gen: gen_model