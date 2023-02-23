#!/bin/bash

docker run \
  --name raspberrypiLocal-cronTrigger \
  --env-file=.env \
  --env-file=.env.dev \
  --env-file=.env.secrets \
  --net mqtt-network \
  -itd \
  c4stus/raspberrypi:crontrigger
