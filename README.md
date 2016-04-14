## Requirements
	Go 1.6 installed

## Installing Go
- Using `homebrew`:

	```sh
	brew update
	brew install go
	```

- Pedestrian way:
	Download and install: https://golang.org/dl/

## Environment setup
	Create a folder where your go code will leave and add it to your env variables under the name `GOPATH`.
	For example, you could run:

	```sh
	mkdir $HOME/go
	```
	and then add
	
	```sh
	GOPATH=$HOME/go
	```

	to your dotfile (or to your Environment Variables, if Windows). The folder you use is arbitrary, here's mine:
	
	```sh
	echo $GOPATH
	/Users/mvalle/code/golang
	```
## Am I ready?

To make sure, let's greet the world. Create a file named `hello.go` and paste this:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

and then run (in the folder you saved the file) the following:


```sh
go run hello.go
```

You know what to expect. 

## Final notes
The whole burrito can be found at https://golang.org/doc/install
If something goes wrong we can troubleshoot at the start of the workshop. 
Else, you are good to Go!
