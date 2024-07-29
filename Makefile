build:
	@cd cmd && go build -o ../bin/huffman
run: build
	@./bin/huffman
