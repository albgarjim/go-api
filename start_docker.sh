echo "Building go app"
CGO_ENABLED=0 go build 
cd ..
echo "Stopping container"
sudo docker-compose down; 
echo "Building container"
sudo docker-compose up --build