
## Install go(lang)
 with [homebrew](http://mxcl.github.io/homebrew/):
```Shell
$ sudo brew install go
```
with [apt](http://packages.qa.debian.org/a/apt.html)-get:

```Shell
$ sudo apt-get install golang
```

## Setup the workspace:
### Add Environment variables:

Go has a different approach of managing code, you'll need to create a single Workspace for all your Go projects. 

We'll add some environment variables into shell config. One of does files located at your home directory  `bash_profile`,  `bashrc`  or  `.zshrc`  (for Oh My Zsh Army)

```
$ vi .bashrc
```
Then add those lines to export the required variables
```
# This is actually your .bashrc file

export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

### Create your workspace:

Create the workspace directories tree:

```
$ mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin
```

## Project Setup

### Setting Up Codebase

1. Clone this Repository inside $GOPATH/src directory.

2. Install [glide](https://github.com/Masterminds/glide/blob/master/README.md) for dependency management .
with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
$ brew install glide
```
3. Install all project dependencies.
```
$ cd $GOPATH/src/smallcase
$ glide install
```

### Setting Up Database
1. Install [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/).
2. Execute all the DDL queries in MySQL present in 
> **$GOPATH/src/smallcase/resources/queries.txt**

### Running the application
1. Run the following commands on the terminal
```
$ cd $GOPATH/src/smallcase
$ source ./.autoenv.zsh
$ go run main.go
```
>**Note:**  you might have to change application configurations according to your setup.  (ref $GOPATH/src/smallcase/config.json
