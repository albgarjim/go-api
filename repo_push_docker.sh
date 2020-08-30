if [[ $# -eq 0 ]] ; then
    echo 'a version number is required in the form of X.X.X'
    exit 1
fi

echo "Building go app"
CGO_ENABLED=0 go build

sudo docker build -t registry.gitlab.com/aikone_fashion/goggers:$1 .
sudo docker push registry.gitlab.com/aikone_fashion/goggers:$1