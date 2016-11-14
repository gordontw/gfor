
BIN = gfor
SOURCES = src/main.go src/yaml.go src/parse.go

all:
	go build  -o $(BIN) $(SOURCES)
