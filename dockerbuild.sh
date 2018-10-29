set -ex
GOOS=linux go build .
docker build -t scratchview .
rm scratchview
docker run scratchview
docker run scratchview ls
docker run scratchview ls dev
docker run scratchview cat etc/hosts
docker run scratchview env
