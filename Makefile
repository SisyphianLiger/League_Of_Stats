TARGET = bin/GoLOL

gorun: clean build
		@./Backend/$(TARGET)

default: ./$(TARGET)

clean:
		@rm -f Backend/bin/*

build:
	@mkdir -p bin
	@cd Backend && go build -o $(TARGET)
	@cd ..

live:
	@live-server --port=8080 --entry-file=index.html
