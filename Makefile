TAG=ab-hack
SVC=$(TAG)
BUILDER=ab-hack-builder:1.16
export REVISION_ID=$(shell grep '"version":' version.json | cut -d\" -f4)
export BUILD_DATE=$(shell date '+%H:%M:%S-%Y/%m/%d')

RUN=docker run --rm \
	-v $(CURDIR):/opt/go/src/$(SVC) \
	-v $(HOME)/go/pkg/mod:/opt/go/pkg/mod \
	-w /opt/go/src/$(SVC) \
	-e GO111MODULE=on

build:
	docker build -f Dockerfile.build -t $(BUILDER) .
	$(RUN) -e CGO_ENABLED=0 -e GOOS=linux $(BUILDER) go build -o ./service \
		-ldflags "-s -X hackaton/version.buildID=$(REVISION_ID) -X hackaton/version.buildDate=$(BUILD_DATE)" \
	 	./cmd/hackaton/...

run:
	docker-compose up