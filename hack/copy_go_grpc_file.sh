#!/bin/bash

for file in `\find ./bazel-bin/api/go_grpc/api -maxdepth 1 -type f | grep .pb.go`; do
    cp -r $file api/
done