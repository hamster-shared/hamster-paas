go build -o build/hamster-paas .
docker buildx build -t hamstershare/hamster-paas:latest --platform linux/amd64 --push .
ssh -i ~/Downloads/2.pem ubuntu@54.69.42.237 cd /home/ubuntu/github/hamster-paas && docker pull hamstershare/hamster-paas:latest && docker-compose up -d