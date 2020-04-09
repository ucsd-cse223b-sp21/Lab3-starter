# Lab 2

## Environment setup
You need to [install Go](https://golang.org/doc/install) to work with these assignments. While 
you are free to use any platform that supports Go environment, we recommend using Linux with basic tools 
like `git`, `make` and `gcc` installed. For you convenience, we created a AWS AMI image pre-configured
with proper tools which you can use to launch your AWS EC2 instance and start working right away.
Follow the instructions [here](./ec2-setup.md) to set up the EC2 instance.


## Download

1. Use the _fork link_ from the class website to create your private clone of the starter code.
2. Do `git clone https://github.com/ucsd-cse223b-sp20/lab2-XXX` where `XXX` is your private repo.
3. Go has expectations about environment variables. For convenience, you might set these variables in your `.bashrc` and/or `.bash_profile` files so that you don't have to execute the commands every time:
    ```
    export PATH=$PATH:/usr/local/go/bin     # making sure go is on path
    export GOPATH=<path-to-repo>
    export PATH=$PATH:$GOPATH/bin           # before running this, see NOTE 2 below
    ```
**NOTE 1:** For AWS machines, only commands added to `.bash_profile` file seems to stick.  
**NOTE 2:** If you had worked with prevous labs on the same machine, make sure to remove `$GOPATH/bin` added for previous labs from the `PATH`. You can simply print the `PATH`, remove the `.../lab1-XXX/bin` folder from the result and set it back to `PATH`. 

## Link to Upstream

Link your clone to the "upstream" to get any updates
```
$ make upstream
```

after this you can get "updates" (in case we modify the starter code), with

```
$ make update
```

## Do the Assignment 

You will be building on top of your work in Lab 1. Copy your solution for Lab 1 into this repository before you start working on Lab 2, and feel free to update your Lab 1 solution if necessary.

1. Implement the `"todo"` methods in `src/triblab/lab2.go` following the requirements posted on class website (links below).
    - [Lab Home](https://cseweb.ucsd.edu/classes/sp20/cse223B-a/labs.html)
    - [Lab 1](https://cseweb.ucsd.edu/classes/sp20/cse223B-a/lab1.html)
    - [Lab 2](https://cseweb.ucsd.edu/classes/sp20/cse223B-a/lab2.html)  
    

2. Build and run local tests 
    ```
    $ cd src/triblab
    $ make all
    $ make test-lab2
    ```
    Unlike Lab 1, you do not get a comprehensive test-suite. You are encouraged to write more tests by yourself.


## Submit 

Save (and submit) your work with:
```
$ git commit -a -m MESSAGE
$ git push
```
