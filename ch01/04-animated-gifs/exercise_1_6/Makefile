EXECUTABLE=exercise_1_6
SOURCE=main.go
OUTPUT=out.gif

# Build GIF
$(OUTPUT): $(EXECUTABLE)
	./$(EXECUTABLE) > $(OUTPUT)

# Build app
$(EXECUTABLE): main.go
	go build -o $(EXECUTABLE) $(main.go)

run:
	./$(EXECUTABLE) > $(OUTPUT)