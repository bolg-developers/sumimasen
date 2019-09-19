#!/bin/bash

cp -r ./bazel-bin/api/android_grpc/android_grpc.jar .
jar -xvf ./android_grpc.jar 
cp -r ./bolg_developers android/bolg/app/src/main/java
rm -rf android_grpc.jar 
rm -rf bolg_developers