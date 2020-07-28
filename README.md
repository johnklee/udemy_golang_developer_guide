## Introduction
Keep sample code and note from practice of udemy course https://www.udemy.com/course/go-the-complete-developers-guide/

## Section1 - Getting Started

### Environment Setup
([course link](https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797236#overview)) Please follow the course to prepare the Go runtime and development environment. To make sure your go is ready to go, try command below:
```console
# go version
go version go1.13.4 linux/amd64
```

## Section2 - A Simple Start 

### Boring OI'Hello World
([course link](https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797244#overview))

### Five Important Questions
([course link](https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797246#overview))<br/>
Five questions to ask about:
* How do we run the code in our project?
* What does `package main` mean?
* What does `import "fmt"` mean?
* What's that `func` thing?
* How is the `main.go` file organized?

To run `main.go` by command below:
```console
// Make sure you have main.go under below path
# pwd
/github_root/udemy_golang_developer_guide/code/helloworld

# go run main.go
Hi there
```

Basically, the Go support commands as below ([more](https://golang.org/cmd/)):
</br>
![git command](images/S2_1.PNG)
<br/>
So we can build tne file `main.go` and produce executable file `main` as below:
```console
# go build main.go
# ls
main  main.go

# ./main
Hi there
```
