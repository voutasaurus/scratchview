# scratchview

A view inside the `scratch` docker container.

# Prerequisites

Go and Docker must be installed.

# Usage
```
$ ./dockerbuild.sh
++ GOOS=linux
++ go build .
++ docker build -t scratchview .
Sending build context to Docker daemon  2.187MB
Step 1/3 : FROM scratch
 ---> 
Step 2/3 : COPY scratchview /
 ---> Using cache
 ---> dc8ea0a309df
Step 3/3 : ENTRYPOINT ["/scratchview"]
 ---> Using cache
 ---> a5f959489f67
Successfully built a5f959489f67
Successfully tagged scratchview:latest
++ rm scratchview
++ docker run scratchview
.dockerenv
dev
etc
proc
sys
++ docker run scratchview ls
.dockerenv
dev
etc
proc
sys
++ docker run scratchview ls dev
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
++ docker run scratchview cat etc/hosts
127.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.2	a55f28dd7496
++ docker run scratchview env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=5c00aaf12b45
HOME=/
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
