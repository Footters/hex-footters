docker: dockerb docker-compose
dockerb: build-docker-auth build-docker-media
build-docker-auth:
	docker build -f dockerfiles/auth.Dockerfile -t auth .
build-docker-media:
	docker build -f dockerfiles/media.Dockerfile -t media .
docker-compose: 
	docker-compose down
	docker-compose up
