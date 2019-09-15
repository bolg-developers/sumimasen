api/build: api/clean
	mkdir api/build
	mkdir api/build/go
	mkdir api/build/java
	./scripts/protogen.py
	cp -r api/build/go/github.com/bolg-developers/bolg/proto server/proto

api/clean:
	rm -rf api/build

server/run:
	go run server/app/app.go

server/dotenv:
	cp -r server/.env.temp server/.env
