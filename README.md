# scratchview

A view inside the `scratch` docker container.

# Prerequisites

Go and Docker must be installed.

# Usage

```
$ go get -u github.com/voutasaurus/scratchview
$ cd ${GOPATH:-~/go}/src/github.com/voutasaurus/scratchview
$ ./dockerbuild.sh
++ GOOS=linux
++ go build .
++ docker build -t scratchview .
Sending build context to Docker daemon   2.13MB
Step 1/3 : FROM scratch
 ---> 
Step 2/3 : COPY scratchview /
 ---> Using cache
 ---> 508c6ea3c002
Step 3/3 : ENTRYPOINT ["/scratchview"]
 ---> Using cache
 ---> a99cbf0a47f7
Successfully built a99cbf0a47f7
Successfully tagged scratchview:latest
++ rm scratchview
++ docker run scratchview
.dockerenv
dev
etc
proc
sys
++ docker run scratchview dev
core
fd
full
mqueue
null
ptmx
pts
random
shm
stderr
stdin
stdout
tty
urandom
zero
++ docker run scratchview etc/hosts
127.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.2	6b1107661ca7
```

## Further usage

To view the contents of a directory in the scratch container:
```
docker run scratchview ls <directory path>
```

Or a file:
```
docker run scratchview cat <file path>
```

To view the environment variables set:
```
docker run scratchview env
```
