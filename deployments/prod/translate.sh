#!/bin/bash
source ./env.sh;
rm -rf Dockerfile;
envsubst < "template" > "Dockerfile";
