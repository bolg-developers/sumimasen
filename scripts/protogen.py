#!/usr/bin/env python3
import glob
import subprocess

gocmd = 'protoc -I . {} --go_out=plugins=grpc:api/build/go'
javacmd = 'protoc -I . {} --grpc-java_out=lite:api/build/java --java_out=api/build/java --plugin=protoc-gen-grpc-java=/usr/local/bin/protoc-gen-grpc-java'


if __name__ == '__main__':
	files = glob.glob('./api/**/*.proto', recursive=True)
	for f in files:
		print('build', f)
		subprocess.call(gocmd.format(f), shell=True)
		subprocess.call(javacmd.format(f), shell=True)