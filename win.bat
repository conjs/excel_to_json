SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags="-s -w" -v -o build/run.exe main.go
upx build/run.exe