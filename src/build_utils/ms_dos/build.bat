SET "ScriptDir=%~dp0"

cd /d "%ScriptDir%\..\..\.."

go build -o gsm.exe .\src
