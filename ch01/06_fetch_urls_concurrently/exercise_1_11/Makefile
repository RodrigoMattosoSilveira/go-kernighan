APP=exercise_1_11

SECTION=$(shell pwd)
BIN=$(SECTION)/../../../bin

SOURCE=main.go
EXECUTABLE=$(BIN)/$(APP)
OUTPUT=$(SECTION)/../../../output/$(APP).txt

#Build data file
$(OUTPUT): $(EXECUTABLE)
	$(EXECUTABLE) \
		google.com \
		youtube.com \
		facebook.com \
		baidu.com \
		wikipedia.org \
		qq.com \
		taobao.com \
		yahoo.com \
		tmall.com \
		amazon.com \

# Build app
$(EXECUTABLE): $(SOURCE)
	go build -o $(EXECUTABLE) $(main.go)

run_one:
	$(EXECUTABLE) google.com

run_many:
	$(EXECUTABLE) \
		google.com \
		youtube.com \
		facebook.com \
		baidu.com \
		wikipedia.org \
		qq.com \
		taobao.com \
		yahoo.com \
		tmall.com \
		amazon.com \



