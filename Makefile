BINARY=bin/ChernOpenGL.exe
SOURCE=cmd/chernopengl/main.go

.PHONY: build run clean

build:
	GOARCH=amd64 GOOS=windows go build -o ${BINARY} ${SOURCE}

run: ${BINARY}
	${BINARY}

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi

${BINARY}: build
