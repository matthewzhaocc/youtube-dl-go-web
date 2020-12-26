#!/bin/bash
mkdir downloads/
# build with static linking netgo package
GOOS=linux go build -tags netgo -a -v