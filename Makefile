.PHONY := all examples intEnums stringEnums

all: go-genums examples

go-genums:
	go build

examples: examples/intEnums/generatedIntEnums.go examples/stringEnums/generatedStringEnums.go examples/intEnums/intEnums examples/stringEnums/stringEnums

examples/intEnums/generatedIntEnums.go:
	cd examples/intEnums && go generate

examples/intEnums/intEnums:
	cd examples/intEnums && go build

examples/stringEnums/generatedStringEnums.go:
	cd stringEnums && go generate

examples/stringEnums/stringEnums:
	cd examples/stringEnums && go build

clean:
	rm -f go-genums examples/stringEnums/generatedStringEnums.go examples/intEnums/generatedIntEnums.go examples/intEnums/intEnums
