server/run:
	bazel run //server/cmd

api/build:
	bazel build //api:go_grpc
	./hack/get_grpc_file.sh

gazelle:
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod