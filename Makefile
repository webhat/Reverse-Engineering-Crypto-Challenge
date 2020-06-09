
all:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o Test10.x86
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  Test10.amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o Test10_x86.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o Test10_amd64.exe
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o Test10_osx.exe
