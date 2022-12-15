APP=boiling

# Our current location
SECTION=$(shell pwd)

# OUTPUT
# N/A

# Executable
BIN_DIR=$(SECTION)/../../../bin
BIN_EXE=$(BIN_DIR)/$(APP)

# Source
SRC_DIR=.
SRC_MAIN=./main.go
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

.PHONY: my_port
my_port:
	lsof -i :8080

.PHONY: client
client:
	echo "nothing"

.PHONY: debug_me
debug_me:
	echo "SRC_FILES: ${SRC_FILES}"
	echo "SRC_MAIN: ${SRC_MAIN}"
	echo "BIN_EXE: ${BIN_EXE}"
	echo "BIN_DIR: ${BIN_DIR}"
	echo "APP: ${APP}"
