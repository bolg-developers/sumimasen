#!/bin/bash

cp -r ./bazel-bin/api/android_grpc/android_grpc.jar .
jar -xvf ./android_grpc.jar 
cp -r ./bolg_developers/sumimasen/api android/app/src/main/java/bolg_developers/sumimasen
rm -rf android_grpc.jar 
rm -rf bolg_developers