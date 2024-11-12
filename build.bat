@echo off

echo "syncproto ..."
call ./syncproto.bat
echo "syncproto finish"

echo "makeproto ..."
call ./makeproto.bat
echo "makeproto finish"

echo "go build ..."
go build -ldflags "-H windowsgui -w -s" -o robot.exe ./src/main.go
echo "go build finish"

pause