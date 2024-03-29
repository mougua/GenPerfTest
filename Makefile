#ARCH
ARCH="`uname -s`"
LINUX="Linux"

all: clean build-test-linux build-test-windows

build-test-linux:
	GOOS=linux GOARCH=amd64	go build -o dist/linux/ ./src/gensipptest.go
	GOOS=linux GOARCH=amd64 go build -o dist/linux/ ./src/genpjtest.go

build-test-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/windows/ ./src/gensipptest.go
	GOOS=windows GOARCH=amd64 go build -o dist/windows/ ./src/genpjtest.go

build-diskbench-windows:
	GOOS=windows GOARCH=amd64
	GOOS=windows GOARCH=amd64 go build -o dist/windows/ ./src/diskbench.go

build-diskbench-linux:
	GOOS=linux GOARCH=amd64	go build -o dist/linux/ ./src/diskbench.go


build-genusers-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/windows/ ./src/genusers.go

build-reportlog2db-windows:
	GOOS=windows GOARCH=amd64 go build -o dist/windows/ ./src/reportlog2db.go

build-reportlog2db-linux:
	GOOS=linux GOARCH=amd64	go build -o dist/linux/ ./src/reportlog2db.go

clean:
	go clean
	rm -rf dist

real-clean: clean
	rm -rf release
	rm -rf genfile

THEME_PATH=theme1
release-theme1	: clean build-test-windows
	rm -rf release/$(THEME_PATH)
	mkdir -p release/$(THEME_PATH)
	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		echo $(ARCH); \
	else \
		echo $(ARCH); \
		dist/windows/genpjtest.exe -group 10003_ -idip 172.16.23.52 -reg 172.16.23.52:6060 -password 888888 -set "6101,6001,99|6300,6200,100"; \
		dist/windows/gensipptest.exe -group 10003_ -set "6101,6001,99|6300,6200,100"; \
		cp -r genfile/* release/$(THEME_PATH); \
		cp resource/pjsua/* release/$(THEME_PATH)/pj_test/; \
		cp resource/sipp/* release/$(THEME_PATH)/sipp_test/; \
	fi

THEME_PATH=theme-chitu
release-theme-chitu	: clean build-test-windows
	rm -rf release/$(THEME_PATH)
	mkdir -p release/$(THEME_PATH)
	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		echo $(ARCH); \
	else \
		echo $(ARCH); \
		dist/windows/genpjtest.exe -group 10000_ -idip 172.16.23.123 -reg 172.16.23.123:6060 -password 888888 -set "1000,6000,400"; \
		dist/windows/gensipptest.exe -group 10003 -set "6101,6001,100"; \
		cp -r genfile/* release/$(THEME_PATH); \
		cp resource/pjsua/* release/$(THEME_PATH)/pj_test/; \
		cp resource/sipp/* release/$(THEME_PATH)/sipp_test/; \
	fi


THEME_PATH=theme-prod
release-theme-prod	: clean build-test-windows
	rm -rf release/$(THEME_PATH)
	mkdir -p release/$(THEME_PATH)
	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		echo $(ARCH); \
	else \
		echo $(ARCH); \
		dist/windows/genpjtest.exe -idip 10.150.94.4 -reg 10.150.94.4:5060 -password 888888 -set "10001,11001,400"; \
		dist/windows/gensipptest.exe -set "10001,9999001,300"; \
		cp -r genfile/* release/$(THEME_PATH); \
		cp resource/pjsua/* release/$(THEME_PATH)/pj_test/; \
		cp resource/sipp/* release/$(THEME_PATH)/sipp_test/; \
	fi

THEME_PATH=theme-test
release-theme-test	: clean build-test-windows
	rm -rf release/$(THEME_PATH)
	mkdir -p release/$(THEME_PATH)
	@if [ $(ARCH) = $(LINUX) ]; \
	then \
		echo $(ARCH); \
	else \
		echo $(ARCH); \
		dist/windows/genpjtest.exe -group 10003_ -idip 10.150.94.5 -reg 10.150.94.5:6060 -password 888888 -set "6000,6000,451"; \
		dist/windows/gensipptest.exe -group 10003_ -set "6500,87790001,400"; \
		cp -r genfile/* release/$(THEME_PATH); \
		cp resource/pjsua/* release/$(THEME_PATH)/pj_test/; \
		cp resource/sipp/* release/$(THEME_PATH)/sipp_test/; \
	fi
