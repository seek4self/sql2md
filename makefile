GC=go
MAIN_GO_Dir=.
APP_NAME=sql2md
RELEASE_DIR=release

build-unix :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GC} build -ldflags '-s -w' -o ${RELEASE_DIR}/${APP_NAME}_unix ${MAIN_GO_Dir}

build-mac :
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 ${GC} build -ldflags '-s -w' -o ${RELEASE_DIR}/${APP_NAME}_mac ${MAIN_GO_Dir}

build-win :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ${GC} build -ldflags '-s -w' -o ${RELEASE_DIR}/${APP_NAME}_win.exe ${MAIN_GO_Dir}

build :
	make build-unix
	make build-mac
	make build-win

release :
	make build
	upx ${RELEASE_DIR}/${APP_NAME}_*
#	tar -czvf ${RELEASE_DIR}/${APP_NAME}_mac.tar.gz -C ${RELEASE_DIR}/ ${APP_NAME}_mac
#	tar -czvf ${RELEASE_DIR}/${APP_NAME}_unix.tar.gz -C ${RELEASE_DIR}/ ${APP_NAME}_unix
#	tar -czvf ${RELEASE_DIR}/${APP_NAME}_win.tar.gz -C ${RELEASE_DIR}/ ${APP_NAME}_win

clean :
	@if [ -d ${RELEASE_DIR} ] ; then rm -rf ${RELEASE_DIR}; fi