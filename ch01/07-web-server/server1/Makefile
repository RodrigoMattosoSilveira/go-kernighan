APP=server1

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

.PHONY: client
client:
	$(BIN)/$(CLIENT) :8080
