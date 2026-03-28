TEST_IMAGE := dotmgr-test

.PHONY: build test test-image clean

build:
	docker buildx bake

test:
	@docker run --rm -v $(CURDIR):/src -w /src $(TEST_IMAGE) go test ./...

test-image:
	docker build -t $(TEST_IMAGE) -f Dockerfile.test .

clean:
	rm -rf bin/
