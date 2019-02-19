@echo off
taskkill /f /t /im launcher.exe & flutter build bundle & start /b launcher.exe