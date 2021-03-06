ifeq ($(PHPCFG),)
	PHPCFG=/usr/bin/php-config
endif

PHPEXE := $(shell $(PHPCFG) --php-binary)
PHPDIR := $(shell $(PHPCFG) --prefix)
BIN = gfor
SOURCES = src/main.go src/yaml.go src/parse.go src/health.go src/cache.go

export PATH := $(PHPDIR)/bin:$(PATH)
export CFLAGS := $(shell $(PHPCFG) --includes)
export LDFLAGS := -L$(shell $(PHPCFG) --prefix)/lib/

export GOPATH := $(PWD):$(GOPATH)
export CGO_CFLAGS := $(CFLAGS) $(CGO_CFLAGS)
export CGO_LDFLAGS := $(LDFLAGS) $(CGO_LDFLAGS)

all:
	make exec
	make phpext

phpext:
	go install ./zend
	go install ./phpgo
	go build -v -buildmode=c-shared -o php_gfor.so src/php_gfor.go src/main.go src/parse.go src/yaml.go src/health.go src/cache.go

exec:
	go build  -o $(BIN) $(SOURCES)

clean:
	rm -f $(GOPATH)/pkg/linux_amd64/zend.a
	rm -f $(GOPATH)/pkg/linux_amd64/phpgo.a
	rm -f hello.so
