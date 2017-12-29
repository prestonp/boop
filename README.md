Run nginx reverse proxy

```
docker run -d -p 80:80 -v /var/run/docker.sock:/tmp/docker.sock:ro jwilder/nginx-proxy
```

Run ci server

```
docker build -t boop .
docker run -e VIRTUAL_HOST=ci.preston.io --rm -d -p 8080:8080 --name boop boop
```
