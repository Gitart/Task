echo off >>null
rem SET CGO_ENABLED=0 

rem rem UNIX
rem SET GOOS=linux
rem SET GOARCH=amd64
rem SET CGO_ENABLED=0 



REM путь к компилятору
SET GOPATH=%CD%
SET GOROOT=C:\GO
SET PATH=%GOROOT%\BIN;%PATH%;
CLS

rem go get github.com/dancannon/gorethink

go build -o serv.exe
serv.exe
pause
