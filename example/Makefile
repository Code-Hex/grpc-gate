.PHONY: example
example:
	@echo "+ $@"
	@echo "+ build server"
	go build -o bin/server -trimpath -mod=readonly \
        github.com/Code-Hex/grpc-gate/example/server
	@echo "+ build client"
	go build -o bin/client -trimpath -mod=readonly \
        github.com/Code-Hex/grpc-gate/example/client

.PHONY: docker/start-mysql
docker/start-mysql:
	docker run --rm -d -p 3306:3306 --name mysql-server-5.7 \
		-e MYSQL_ALLOW_EMPTY_PASSWORD=1 \
		-e MYSQL_DATABASE=mysql \
		mysql:5.7

.PHONY: docker/stop-mysql
docker/stop-mysql:
	docker stop mysql-server-5.7
