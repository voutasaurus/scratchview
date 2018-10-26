set -ex
GOOS=linux go build .
docker build -t scratchview .
rm scratchview
docker run scratchview
docker run scratchview dev
docker run scratchview etc/hosts
