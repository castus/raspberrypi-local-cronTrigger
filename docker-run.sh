#!/bin/bash

docker run \
  --rm \
  --name raspberrypiLocal-cronTrigger \
  -v "$(pwd)"/src:/data \
  --workdir /data \
  --env-file=.env \
  --env-file=.env.dev \
  --env-file=.env.secrets \
  --net mqtt-network \
  -itd \
  c4stus/raspberrypi:crontrigger \
  /bin/bash -c "sh run.sh"
