boop
---

boop is a really crappy CI server to handle building and deploying my blog on git pushes.

## setup

Compile and build docker image

```
make && make image
```

Run nginx reverse proxy

```
docker run -d -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock:ro jwilder/nginx-proxy
```

Run ci server

```
docker build -t boop .
docker run -e AWS_ACCESS_KEY_ID=<ACCESS KEY> \
	-e AWS_SECRET_ACCESS_KEY=<SECRET KEY> \
	-e VIRTUAL_HOST=ci.preston.io \
	--rm -d -p 8080:8080 --name boop boop
```
