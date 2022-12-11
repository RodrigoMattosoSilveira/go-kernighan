APP=lissajous

SECTION=$(shell pwd)
BIN=$(SECTION)/../../../bin


SOURCE=main.go
EXECUTABLE=$(BIN)/$(APP)
OUTPUT=$(SECTION)/../../../output/$(APP).gif

#Build data file
$(OUTPUT): $(EXECUTABLE)

# Build app
$(EXECUTABLE): $(SOURCE)
	go build -o $(EXECUTABLE) $(SOURCE)

.PHONY: run
run:
	$(EXECUTABLE) > $(OUTPUT)

