
BIN = gfor
SOURCES = src/main.go src/yaml.go 

all:
	go build  -o $(BIN) $(SOURCES)
