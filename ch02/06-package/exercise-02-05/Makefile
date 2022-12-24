APP=none

# Our current location
SECTION=$(shell pwd)

# OUTPUT
# N/A

# Executable
BIN_DIR=$(SECTION)/../../../bin
BIN_EXE=$(BIN_DIR)/$(APP)

# Source
SRC_DIR=.
SRC_MAIN=./cf.go
SRC_FILES=$(shell find $(SRC_DIR) -type f \( -name "*.go" -o -name "*.go" \))
OTHER_DEPS=

#Build output file
# N/A

# Build app
$(BIN_EXE): $(SRC_FILES) $(OTHER_DEPS)
	go build -o $(BIN_EXE) $(SRC_MAIN)

.PHONY: run
run:
	$(BIN_EXE)

.PHONY: test
test:
	richgo test -v ./...

.PHONY: bench
bench:
	go test -bench .