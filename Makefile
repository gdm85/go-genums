.PHONY := all examples examples-generated
.DEFAULT := all

all: genums examples-generated examples

genums:
	go build -o genums

examples-generated: examples/dayEnum/generatedEnum.go examples/colorEnum/generatedEnum.go examples/carEnum/generatedEnum.go

examples/dayEnum/generatedEnum.go examples/colorEnum/generatedEnum.go examples/carEnum/generatedEnum.go: %/generatedEnum.go:
	cd $* && go generate

examples: examples/dayEnum/dayEnum examples/colorEnum/colorEnum examples/carEnum/carEnum

examples/dayEnum/dayEnum: examples/dayEnum/generatedEnum.go
	cd examples/dayEnum && go build

examples/colorEnum/colorEnum: examples/colorEnum/generatedEnum.go
	cd examples/colorEnum && go build

examples/carEnum/carEnum: examples/carEnum/generatedEnum.go
	cd examples/carEnum && go build

clean:
	rm -f genums \
			examples/dayEnum/generatedEnum.go examples/dayEnum/dayEnum \
			examples/colorEnum/generatedEnum.go examples/colorEnum/colorEnum \
			examples/carEnum/generatedEnum.go examples/carEnum/carEnum
