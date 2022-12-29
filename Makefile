BUILD_DIR = build
APP = deployer
MC_DIR = build/bin
LOG_DIR= build/log
GIT_VER=$(shell git rev-parse HEAD)

LDFLAGS=-ldflags "-X main.version=${GIT_VER}"

.PHONY: dut app.darwin64 app.darwinArm lint clean distclean mrproper


# Build the project
all:
	@echo "cmd:"
	@echo ""
	@echo "  app            build all app for all os"
	@echo "  app.darwin64   build app for osx64"
	@echo "  app.darwinArm  build app for osx arm64"
	@echo "  app.linux64    build app for linux arm64"
	@echo ""
	@echo "  lint           go linter"
	@echo ""
	@echo "  clean          remove dut binarys"
	@echo "  distclean       remove build folder"

app: app.windows app.windows64  app.darwin64 app.darwinArm app.linux64


app.windows:
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}.exe -v cmd/deployer/deployer.go

app.windows64:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}64.exe -v cmd/deployer/deployer.go

app.darwin64:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}-darwin -v cmd/deployer/deployer.go

app.darwinArm:
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}-darwin-arm -v cmd/deployer/deployer.go

app.linux64:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/${APP}-linux -v cmd/deployer/deployer.go


lint:
	golint -set_exit_status $(shell go list ./...)

clean:
	-rm -f ${BUILD_DIR}/${DUT}-*
distclean:
	rm -rf ./build
