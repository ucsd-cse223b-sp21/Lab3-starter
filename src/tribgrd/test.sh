#!/bin/bash -x

if ! [[ $GOPATH ]]; then 
    echo "GOPATH variable not set"; 
    exit 1
fi

# test descriptions
declare -a messages=(
    'your code passes the basic test'
    'your rpc client does not return a nil list as an empty list'
    'your rpc client clears the list when it wants to return an empty list'
    'ServeBack() blocks on serving'
    'ServeBack() sends a false to Ready channel on listen error'
    'ServeBack() returns error when the address is not resolvable'
    'ServeBack() uses the storage provided'
    'ServeBack() can spawn two back-ends in one single process'
    'your rpc client can connect to another rpc server implementation'
    'your rpc server can accept from another rpc client implementation'
    'your rpc client works when the server restarts after the client creates'
    'your rpc client works when the server restarts between two rpc calls'
    'ServeBack() returns error when the listening socket is shutdown'
    'your implementation handles concurrent list appends'
    'your rpc client returns error when the connection is lost'
)

echo "(compiling)"
go install -v trib/... triblab/... tribgrd/...

echo "(starts testing)"
i=0
for b in `ls ${GOPATH}/bin/tg-l1*`; do
    bname=$(basename $b)
    echo -n "... test $bname  ${messages[$i]} ==> "
    ${GOPATH}/bin/tg-run ${GOPATH}/log/$bname $b
    i=$((i+1))
done

