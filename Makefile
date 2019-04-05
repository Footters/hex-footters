docker: dockerb docker-compose
dockerb: build-docker-auth build-docker-media
docker-compose: docker-compose down docker-compose up
build-docker-auth:
	docker build -f dockerfiles/auth.Dockerfile -t auth .
build-docker-media:
	docker build -f dockerfiles/media.Dockerfile -t media .
createContent:
	curl http://localhost:3000/contents
	[{"ID":1,"CreatedAt":"2019-03-19T12:37:24Z","UpdatedAt":"2019-03-19T12:37:24Z","DeletedAt":null,"urlName":"sevilla-fc-vs-real-betis-balompie","title":"Sevilla - Betis","description":"A live match event","-X POST   http://localhost:3000/contents  -H 'Cache-Control: no-cache' -H 'Content-Type: application/json' -d '{
	"urlName" : "sevilla-fc-vs-real-betis-balompie",
	"title" : "Badajoz - Murcia",
	"description" : "A live match event",
	"status": "live",
	"free": 1,
	"visible":1}'
