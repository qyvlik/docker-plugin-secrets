VERSION ?= 0.0.1

create:
	docker build -t rootfsimage .
	docker create --name tmp rootfsimage
	mkdir -p plugin/rootfs/
	docker export tmp | tar -x -C plugin/rootfs/
	docker rm -vf tmp
	docker plugin create docker-plugin-secrets:$(VERSION) ./plugin
	docker rmi --force rootfsimage

enable:
	docker plugin enable docker-plugin-secrets:$(VERSION)

push:
	docker plugin push docker-plugin-secrets:$(VERSION)

.PHONY: clean
clean: clean-cache

.PHONY: clean-cache
clean-cache:
	rm -fr plugin/rootfs/.dockerenv plugin/rootfs/*
	docker plugin rm -f docker-plugin-secrets:$(VERSION)


