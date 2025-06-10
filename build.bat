go build -ldflags="-s -w" -o win64.exe main.go && upx -9 win64.exe
