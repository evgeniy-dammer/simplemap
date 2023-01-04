build: clean
	go get -u github.com/valyala/fasthttp
	go get -u github.com/fasthttp/router
	go build -o tileserver

clean:
	rm -f tileserver

run:
	go run main.go

image: rmimage
	docker build -t tileserver .

rmimage:
	docker image rm -f tileserver

cont:
	docker run -dit --rm -p 8082:8082 --name tileserver tileserver

stopcont:
	docker stop tileserver

prune:
	docker image prune