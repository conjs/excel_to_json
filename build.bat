SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags="-s -w" -v -o build/mac main.go

upx build/mac

::CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
::CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

::SET CGO_ENABLED=0
::SET GOOS=windows
::SET GOARCH=386
::go build -ldflags="-s -w" -v -o build/run.exe main.go
