server/run:
	bazel run //server

api/build: api/build/go api/build/android

api/build/go:
	bazel build //api:go_grpc 
	./hack/copy_go_grpc_file.sh

api/build/android:
	bazel build //api:android_grpc --incompatible_disable_deprecated_attr_params=false
	./hack/copy_android_grpc_file.sh

gazelle:
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod

init: dotenv install

dotenv:
	@cp -r server/.env.temp server/.env

install:
	@which bazel || echo 'bazel command is not found. see: https://docs.bazel.build/versions/master/install.html'