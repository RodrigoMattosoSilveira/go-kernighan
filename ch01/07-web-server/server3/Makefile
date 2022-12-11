APP=server3

SECTION=$(shell pwd)
BIN=$(SECTION)/../../../bin

SOURCE=main.go
EXECUTABLE=$(BIN)/$(APP)
CLIENT=exercise_1_8
#OUTPUT=NA

# Build app
$(EXECUTABLE): $(SOURCE)
	go build -o $(EXECUTABLE) $(SOURCE)

.PHONY: run
run:
	$(EXECUTABLE) &

.PHONY: my_port
my_port:
	lsof -i :8080

.PHONY: slash
slash:
	$(BIN)/$(CLIENT) http://localhost:8080/

.PHONY: count
count:
	$(BIN)/$(CLIENT) http://localhost:8080/count
