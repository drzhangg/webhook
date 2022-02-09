pwd
ls
cat pwd.txt  | docker login -u drzhangg --password-stdin
docker build -t drzhangg/webhook:latest .
docker push drzhangg/webhook:latest

docker