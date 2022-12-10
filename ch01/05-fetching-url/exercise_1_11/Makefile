EXECUTABLE=exercise_1_11
SOURCE=main.go
OUTPUT=data.txt

#Build data file
$(OUTPUT): $(EXECUTABLE)
	./$(EXECUTABLE) google.com

# Build app
$(EXECUTABLE): main.go
	go build -o $(EXECUTABLE) $(main.go)

run:
	./$(EXECUTABLE) \
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

