# Golang Code Dumpster

This is a repository to keep tutorial/experiments go codes

## Running Go Code

To execute a gofile run it directly as below:

```
go run hello.go
```

## Install

Clone this repository on your `src` folder inside your `$GOPATH`.

This repo uses Golang's Dep to manage it's dependencies. Install Dep running:

```
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

This will install a binary on `$GOPATH/bin` so make sure it's in your `$PATH`.

After install `dep` you must run `dep ensure` inside the repository folder to install all dependencies.

After that just run any go code directly
