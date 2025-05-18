@echo off

SET "ScriptDir=%~dp0"
SET "ZipFileName=%ScriptDir%gsm.zip"
SET "File1=%ScriptDir%\gsm.exe"
SET "File2=%ScriptDir%\assets\install.bat"
SET "File3=%ScriptDir%\update_path.bat"


start /wait cmd /c "%ScriptDir%\build.bat 

copy /y "%ScriptDir%\..\..\..\gsm.exe" %ScriptDir%

echo Checking for source files...

IF NOT EXIST "%File1%" (
    echo ERROR: Source file "%File1%" not found in the current directory.
    goto :end_script
)

IF NOT EXIST "%File2%" (
    echo ERROR: Source file "%File2%" not found.
    echo Ensure the 'assets' folder and 'install.bat' are present relative to this script.
    goto :end_script
)

IF NOT EXIST "%File3%" (
    echo ERROR: Source file "%File2%" not found.
    echo Ensure the 'update_path.bat' are present relative to this script.
    goto :end_script
)

echo Creating zip archive: "%ZipFileName%"

powershell.exe -NonInteractive -NoProfile -Command "Compress-Archive -Path """%File1%""", """%File2%""", """%File3%""" -DestinationPath """%ZipFileName%""" -Force"

if errorlevel 0 (
    echo Zip archive "%ZipFileName%" created successfully.
) else (
    echo ERROR: Failed to create zip archive "%ZipFileName%".
    echo Ensure you have write permissions in the current directory.
)

del %File1%

:end_script
