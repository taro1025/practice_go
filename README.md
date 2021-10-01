# practice_go

Golang

Enable dependency tracking for the code you're about to write.

go mod init example/hello

Add new module requirements and sums.

go mod tidy

func init () is executed automatically

go install can call anywhere

go build can call from the current directly


get dependencies for code in the current directory.

go get .

go.mod

there declares the module path

Maybe, "go init mod" makes vertial path. example, execute "go init mod example/hello", attached current directory.

Note: TO write upper-case latter, function can be exported


Package

is a collection of source files in the same directory 

Repository

 contains one or more modules

Module

 is a collection of related Go packages

