EXECUTABLE=fetchall
SOURCE=main.go
OUTPUT=out.gif

# Build app
$(EXECUTABLE): main.go
	go build -o $(EXECUTABLE) $(main.go)
	./$(EXECUTABLE) http://gopl.io http://godoc.org http://golang.org

run:
	./$(EXECUTABLE) http://gopl.io http://godoc.org http://golang.org
