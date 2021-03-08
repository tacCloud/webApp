#!/bin/bash
#https://devopsheaven.com/docker/dockerhub/2018/04/09/delete-docker-image-tag-dockerhub.html
USERNAME="rmccabe3701"
#DOCKERHUB_PASSWORD should be in the env
ORGANIZATION="${USERNAME}"
IMAGE="inventory-web-app"
TAG=$1

login_data() {
cat <<EOF
{
  "username": "$USERNAME",
  "password": "$DOCKERHUB_PASSWORD"
}
EOF
}

TOKEN=`curl -s -H "Content-Type: application/json" -X POST -d "$(login_data)" "https://hub.docker.com/v2/users/login/" | jq -r .token`

curl "https://hub.docker.com/v2/repositories/${ORGANIZATION}/${IMAGE}/tags/${TAG}/" \
-X DELETE \
-H "Authorization: JWT ${TOKEN}"
