@echo off

set "workdir=%cd%"
set projectdir=D:/Projects/backend-winner-proj

@echo on

xcopy "%projectdir%/docproj/msg" "%workdir%/src/robot/dyj/pb" /E /Y /I

./makeproto.bat