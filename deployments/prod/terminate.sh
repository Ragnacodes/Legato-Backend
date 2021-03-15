#!/bin/bash
source env.sh;
IMAGE=magstream/legato_server:$TAG;
if [ ! "$(sudo docker ps -q --filter ancestor=$IMAGE 2> /dev/null)" ]; then
  echo "'$IMAGE' does not exist.";
else
  sudo docker stop $( sudo docker ps -q --filter ancestor=$IMAGE );
fi
