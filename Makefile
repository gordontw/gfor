
BIN = gfor
SOURCES = src/main.go src/yaml.go src/parse.go src/health.go src/cache.go

all:
	go build  -o $(BIN) $(SOURCES)
