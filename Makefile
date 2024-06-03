BUILD_DIR = build
APP = deployer
CMD= cmd/deployer/main.go
APP_TOOL = secretSealer
CMD_TOOL= cmd/secretSealer/main.go
MC_DIR = build/bin
LOG_DIR= build/log
GIT_VER=$(shell git rev-parse HEAD)

LDFLAGS=-ldflags "-X main.version=${GIT_VER}"

.PHONY: dut app.darwin64 app.darwinArm lint clean distclean mrproper


# Build the project
help:
	@echo "cmd:"
	@echo ""
	@echo "  app            build all app for all os"
	@echo "  app.darwin64   build app for osx64"
	@echo "  app.darwinArm  build app for osx arm64"
	@echo "  app.linux64    build app for linux arm64"
	@echo ""
	@echo "  tool            build all sealer tools for all os"
	@echo "  tool.darwin64   build sealer tool for osx64"
	@echo "  tool.darwinArm  build sealer tool for osx arm64"
	@echo "  tool.linux64    build sealer tool for linux arm64"
	@echo ""
	@echo "  clean          remove dut binarys"
	@echo "  distclean       remove build folder"

all: app tool

app: app.windows app.windows64  app.darwin64 app.darwinArm app.linux64

app.windows:
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}.exe -v ${CMD}

app.windows64:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}64.exe -v ${CMD}

app.darwin64:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}-darwin -v ${CMD}

app.darwinArm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}-darwin-arm -v ${CMD}

app.linux64:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}-linux -v ${CMD}




tool: tool.windows tool.windows64  tool.darwin64 tool.darwinArm tool.linux64

tool.windows:
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP_TOOL}.exe -v ${CMD_TOOL}

tool.windows64:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP_TOOL}64.exe -v ${CMD_TOOL}

tool.darwin64:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP_TOOL}-darwin -v ${CMD_TOOL}

tool.darwinArm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP_TOOL}-darwin-arm -v ${CMD_TOOL}

tool.linux64:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP_TOOL}-linux -v ${CMD_TOOL}

clean:
	rm -f ${BUILD_DIR}/${APP}*
	rm -f ${BUILD_DIR}/${APP_TOOL}*

distclean:
	rm -rf ./build
