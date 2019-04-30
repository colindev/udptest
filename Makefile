
all: bin/udptest.linux bin/udptest.mac bin/udptest.windows

bin/udptest.linux:
	GOOS=linux go build -o bin/udptest.linux src/*.go

bin/udptest.mac:
	GOOS=darwin go build -o bin/udptest.mac src/*.go

bin/udptest.windows:
	GOOS=windows go build -o bin/udptest.windows src/*.go
