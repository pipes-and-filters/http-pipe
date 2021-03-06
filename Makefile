NAME=pipesandfilters/http
default: build

deps:
	go get -v github.com/Masterminds/glide
	glide install

build: deps 
	glide install
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o http-pipe ./http-pipe.go

test: deps
	@glide novendor|xargs go test -v

build-docker: build
	docker build -t $(NAME):$(TRAVIS_COMMIT) .

dockertravisbuild: build-docker
	@docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
	docker push $(NAME):$(TRAVIS_COMMIT)

