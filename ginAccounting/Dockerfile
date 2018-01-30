# mgo container issue로 한 컨테이너안에 api server, mongodb 같이배치

FROM 979156/idk:latest
ENV MONGODB_ADDR 127.0.0.1
WORKDIR /etc/yum.repos.d/
COPY mongodb-org-3.6.repo .
RUN yum install mongodb-org -y
WORKDIR /root/go/ginAccounting/
COPY main.go .
RUN go build
EXPOSE 8080
CMD mongod --config /etc/mongod.conf && /root/go/ginAccounting/ginAccounting 
