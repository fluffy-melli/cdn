@echo off
setlocal
for %%F in ("./packages/api/ui"\*) do (
    "%USERPROFILE%\go\bin\swag" f %%F
    "%USERPROFILE%\go\bin\swag" i %%F
)
go build -o main.exe ./main.go
main.exe