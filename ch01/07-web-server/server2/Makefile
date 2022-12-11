SECTION=$(shell pwd)
BIN=$(SECTION)/../../../bin
APP=server2
SOURCE=main.go
EXECUTABLE=$(BIN)/$(APP)
#OUTPUT=out.gif

# Build app
$(EXECUTABLE): $(SOURCE)
	go build -o $(EXECUTABLE) $(SOURCE)

run:
	$(EXECUTABLE) &

.PHONY: listen
listen:
	lsof -i :8080

.PHONY: slash
slash:
	$(BIN)/fetch http://localhost:8080/

.PHONY: count
count:
	$(BIN)/fetch http://localhost:8080/count