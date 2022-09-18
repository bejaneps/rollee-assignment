run:
	go run cmd/main.go

build-docker:
	docker build -t "rollee" -f Dockerfile .

run-docker:
	docker run -p "7171:7171" rollee