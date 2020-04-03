IMG := ryantate13/hello-golang
VERSION := 1.0.0
PORT := 3000

build: dockerignore
	docker build -t $(IMG) -t $(IMG):$(VERSION) --build-arg VERSION=$(VERSION) --build-arg PORT=$(PORT) .

run: build
	docker run -it --rm -e PORT=$(PORT) -p $(PORT):$(PORT)  $(IMG)

publish: build
	docker push $(IMG)
	docker push $(IMG):$(VERSION)

dockerignore:
	@echo '*' > .dockerignore
	@for f in $$(cat Dockerfile | grep COPY | grep -v -e --from | awk '{print $$2}'); do \
		echo "!$$f" >> .dockerignore; \
	done