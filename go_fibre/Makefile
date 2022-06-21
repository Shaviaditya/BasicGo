build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -r -s '\.go$$' make run