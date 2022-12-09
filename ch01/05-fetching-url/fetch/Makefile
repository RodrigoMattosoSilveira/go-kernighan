EXECUTABLE=fetch
SOURCE=main.go
OUTPUT=out.gif

# Build app
$(EXECUTABLE): main.go
	go build -o $(EXECUTABLE) $(main.go)
	./$(EXECUTABLE) http://gopl.io

run:
	./$(EXECUTABLE) http://gopl.ioooo