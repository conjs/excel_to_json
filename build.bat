SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags  -v -o build/run main.go



SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags  -v -o build/run.exe main.go