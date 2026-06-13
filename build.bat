@echo off
go build -ldflags="-s -w -H=windowsgui" -o pcmediarc.exe main.go
echo done.
pause