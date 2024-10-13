windres.exe -i app.rc -o defaultRes_windows_amd64.syso

go build -tags=tempdll -buildmode=exe -ldflags="-s -w -H windowsgui" -o bin/HoneyOrange.exe .

upx bin/HoneyOrange.exe

@REM pause