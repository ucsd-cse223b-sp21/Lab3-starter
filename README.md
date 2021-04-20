# Lab 3
Use this repo as the template for Lab 3. Choose one member of your team and push their lab 2 code to their forked version of this repo.

## Download

1. Use the _fork link_ from the class website to create your private clone of the starter code.
2. Do `git clone https://github.com/ucsd-cse223b-sp21/lab3-XXX` where `XXX` is your private repo.
3. Go has expectations about environment variables. For convenience, you might set these variables in your `.bashrc` and/or `.bash_profile` files so that you don't have to execute the commands every time:
    ```
    export PATH=$PATH:/usr/local/go/bin     # making sure go is on path
    export GOPATH=<path-to-repo>
    export PATH=$PATH:$GOPATH/bin           # before running this, see NOTE 2 below
    ```
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

You will be building on top of your work in Lab 2. Choose a group member and push their lab2 code as the base for lab3. To do this set the remote url of the prior assignment and push. Feel free to update your Lab 2 solution if necessary. Extensive modifications to you lab 2 code may be required to meet the fault tolerance requirements

 ```
 cd chosen-lab2-directory
 git set remote-url lab3-forked-repo.git
 git push
 ```
See this [link](https://docs.github.com/en/github/getting-started-with-github/managing-remote-repositories) for debugging information on this task.


1. Modify methods in `src/triblab/lab2.go` following the requirements posted on class website (links below). Feel free to rename the file lab3.go.
    - [Lab 3](https://cseweb.ucsd.edu/classes/sp21/cse223B-a/lab3.html)  
    

2. Build and run local tests 
    ```
    $ cd src/triblab
    $ make all
    $ make test-lab2 #won't test lab3 requirements
    ```
    Unlike Lab 1, you do not get a comprehensive test-suite. You are encouraged to write more tests by yourself.
    
    Unlike Lab 2, there will be joins an failures in this system which must be handled. Consider using scripts wich make use of `kill` and `wait` to kill and syncronize processes. Coordinating on single machines can be done with file descriptors, or timers for course grain testing.


## Submit 

Save (and submit) your work with:
```
$ git commit -a -m MESSAGE
$ git push
```
