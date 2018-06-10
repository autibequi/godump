# Golang Code Dumpster

This is a repository to keep tutorial/experiments go codes

## Install Go

Download the latest golang binary at `https://golang.org/dl/` and unzip it

```sh
tar -C /usr/local -xzf go.XXX.tar.gz
```

Then setup the go environment

```sh
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

And setup a workspace (in this case on `$HOME/projects/go`)

```sh
export GOPATH=$HOME/projects/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

## Install Project

Clone this repository on `$GOPATH/src/github.com/pedroorez/`.

This repo uses Golang's Dep to manage it's dependencies. Install Dep running:

```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

This will install a binary on `$GOPATH/bin` so make sure it's in your `$PATH`.

After install `dep` you must run `dep ensure` inside the repository folder to install all dependencies.


