build: clean
	go get -u github.com/valyala/fasthttp
	go build -o front

clean:
	rm -f front

run:
	go run main.go

image: rmimage
	docker build -t front .

rmimage:
	docker image rm -f front

cont:
	docker run -dit -p 8080:8080 --rm --name front front

stopcont:
	docker stop front

prune:
	docker image prune