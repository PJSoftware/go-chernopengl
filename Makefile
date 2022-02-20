BINPATH=bin
BINARY=${BINPATH}/ChernOpenGL.exe

RESPATH=res

SRCPATH=cmd/ChernOpenGL
SOURCE=${SRCPATH}/main.go

.PHONY: build run clean

build:
	rm -rf ${BINPATH}
	cp -rf ${RESPATH} ${BINPATH}/${RESPATH}
	GOARCH=amd64 GOOS=windows go build -o ${BINARY} ${SOURCE}

run: ${BINARY}
	${BINARY}

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi

${BINARY}: build
