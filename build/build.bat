SET GOPATH=%cd%\..
:again
	go run -a ../src/main/run.go
	set /p response=Rerun?
	if "%response%" equ "" cls & goto again