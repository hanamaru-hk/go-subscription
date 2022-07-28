# go-subscription
A simple restful api with mongo db connection. It can store and list all the emails for subscription.

# Run in docker
```
docker build -t go-subscription .
docker run --rm -dp 2000:8090 go-subscription
```

# Interactive with docker
```
docker build -t go-subscription .
docker run --rm -it go-subscription
```

# Remove Containers
```
docker kill $(docker ps -q)
```

# Remove Images
```
docker rmi $(docker images -a -q)
```

# Reference
https://blog.codecentric.de/en/2020/04/golang-gin-mongodb-building-microservices-easily/
https://github.com/Andreas-Maier/task-management