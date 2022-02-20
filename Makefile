BINPATH=bin
BINARY=${BINPATH}/ChernOpenGL.exe

RESPATH=res

SRCPATH=cmd/chernopengl
SOURCE=${SRCPATH}/main.go

.PHONY: build run clean

build:
	rm -rf ${BINPATH}/*
	tools/audit-gl-import.pl
	GOARCH=amd64 GOOS=windows go build -o ${BINARY} ${SOURCE}
	cp -rf ${RESPATH} ${BINPATH}/${RESPATH}

run: ${BINARY}
	${BINARY}

clean:
	if [ -d ${BINPATH} ]; then rm -rf ${BINPATH}; fi

${BINARY}: build
