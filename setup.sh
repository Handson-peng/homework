sudo apt install golang -y
sudo add-apt-repository ppa:longsleep/golang-backports -y
sudo apt update
sudo apt install golang-go -y
echo "export GOROOT=/usr/lib/go" >> ~/.profile
echo "export PATH=\$GOROOT/bin:\$PATH" >> ~/.profile
source ~/.profile
sudo apt install docker.io -y
docker run -d -p 27017:27017 --name mongodb mongo:4.4
go install