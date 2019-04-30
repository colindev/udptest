
all: bin/udptest.linux bin/udptest.mac bin/udptest.exe

bin/udptest.linux:
	GOOS=linux go build -a -o bin/udptest.linux src/*.go

bin/udptest.mac:
	GOOS=darwin go build -a -o bin/udptest.mac src/*.go

bin/udptest.exe:
	GOOS=windows go build -a -o bin/udptest.exe src/*.go
