#!/opt/homebrew/bin/zsh

set -ex;

export hamster_paas_version=$(date "+%Y%m%d%H%M%S")

#build
docker buildx build -t hamstershare/hamster-pass:${hamster_paas_version} --platform=linux/amd64 --push .

envsubst < deploy.yml | kubectl -n hamster apply -f -
kubectl -n hamster apply -f nginx.yml
