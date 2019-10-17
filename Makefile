all:
	go build -o bin/oneplanet

clean:
	rm -f bin/oneplanet

run:
	./bin/oneplanet
