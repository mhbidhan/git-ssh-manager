@echo off

SET "TargetDir=C:\Program Files\GitSSHManager"
SET "ScriptDir=%~dp0"

IF NOT EXIST "%TargetDir%" (
    mkdir "%TargetDir%"

    if errorlevel 0 (
        echo Target directory "%TargetDir%" created successfully.
    ) else (
        echo ERROR: Failed to create target directory "%TargetDir%".
        goto :end_script
    )
) ELSE (
    echo Target directory "%TargetDir%" already exists.
)

SET "SourceFile=%ScriptDir%\gsm.exe"

IF NOT EXIST "%SourceFile%" (
    echo ERROR: Source file "%SourceFile%" not found in the current directory.
    goto :end_script
) ELSE (
    copy /y "%SourceFile%" "%TargetDir%" >nul

    if errorlevel 0 (
        echo File "%SourceFile%" copied successfully to "%TargetDir%".
    ) else (
        echo ERROR: Failed to copy "%SourceFile%" to "%TargetDir%".
    )
)

start /wait cmd /c "%ScriptDir%\update_path.bat 

:end_script
echo Installation complete.
pause
