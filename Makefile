server/run:
	bazel run //server

api/build:
	bazel build //api:go_grpc
	./hack/get_grpc_file.sh

gazelle:
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod

init: dotenv install

dotenv:
	@cp -r server/.env.temp server/.env

install:
	@which bazel || echo 'bazel command is not found. see: https://docs.bazel.build/versions/master/install.html'