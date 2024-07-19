#!/bin/bash

echo start...

echo input:$1,$2
if [ $1 = "api" ]; then
    goctl api go --api $2.api --dir ../api/
elif [ $1 = "demo" ]; then
    goctl api go --api $2.api --dir ../demo/
elif [ $1 = "rpc" ]; then
    goctl rpc protoc $2.proto --go-grpc_out=../rpc/$2/  --go_out=../rpc/$2/ --zrpc_out=../rpc/$2/
elif [ $1 = "rpc_py" ]; then
    protoc -I ./ gpt_py.proto --go-grpc_out=../rpc/$2/ --go_out=../rpc/$2/
elif [ $1 = "aa" ]; then
    goctl api plugin -plugin goctl-swagger="swagger -filename  backpack.json -host localhost:8888" -api backpack.api -dir .
else
    echo input error!
fi


##goctl api plugin -plugin goctl-swagger="swagger -filename  backpack.json -host test.wapenai.cn"  -api backpack.api -dir .

#docker run -d -p 3000:8080 --name="swaager" -v /root/wapen/code/wapen/proto/backpack.json:/app/swagger.json swaggerapi/swagger-ui

#docker run -d --name gitlab-runner --restart always   -v /var/run/docker.sock:/var/run/docker.sock   -v /srv/gitlab-runner/config:/etc/gitlab-runner -v /root/.kube/config:/root/.kube/config   gitlab/gitlab-runner:latest
