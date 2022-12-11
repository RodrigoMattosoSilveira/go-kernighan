APP=exercise_1_8

SECTION=$(shell pwd)
BIN=$(SECTION)/../../../bin

SOURCE=main.go
EXECUTABLE=$(BIN)/$(APP)
#OUTPUT=out.gif

# Build app
$(EXECUTABLE): $(SOURCE)
	go build -o $(EXECUTABLE) $(SOURCE)

run:
	$(EXECUTABLE) google.com