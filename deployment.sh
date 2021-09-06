#!/bin/bash

imageId=$(docker images -q adamo/deployment 2> /dev/null)
if [ ! -z $imageId ]; then
  docker run --network adamo --interactive --tty --name deployment --rm --volume /var/run/docker.sock:/var/run/docker.sock adamo/deployment
else
  echo "Image adamo/deployment doesn't exists, did you run then install script?"
fi
