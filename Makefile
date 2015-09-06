.PHONY := all examples intEnums stringEnums

all: go-genums examples

go-genums:
	go build

examples: examples/intEnums/generatedIntEnums.go examples/stringEnums/generatedStringEnums.go examples/intEnums/intEnums

examples/intEnums/generatedIntEnums.go:
	./go-genums Wolf piggy int examples/intEnums/intEnums.go > $@

examples/intEnums/intEnums:
	cd examples/intEnums && go build

examples/stringEnums/generatedStringEnums.go:
	./go-genums Color colour string examples/stringEnums/stringEnums.go > $@

clean:
	rm -f go-genums examples/stringEnums/generatedStringEnums.go examples/intEnums/generatedIntEnums.go examples/intEnums/intEnums
