#!/bin/bash

docker run \
  --rm \
  --name raspberrypiLocal-cronTrigger \
  -v "$(pwd)"/src:/data \
  --workdir /data \
  --env-file=.env.prod \
  --env-file=.env.secrets \
  --net mqtt-network \
  -itd \
  raspberrypiLocal-cronTrigger-img \
  /bin/bash -c "sh run.sh"
