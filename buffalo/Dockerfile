FROM golang:1.19.5-alpine  
ENV TZ=Asia/Kolkata
RUN apk add ca-certificates
RUN wget https://github.com/gobuffalo/cli/releases/download/v0.18.13/buffalo_0.18.13_Linux_x86_64.tar.gz --no-check-certificate
RUN tar -xvzf buffalo_0.18.13_Linux_x86_64.tar.gz
RUN mv buffalo /usr/local/bin/buffalo 

RUN apk update
RUN apk add pkgconfig 
RUN apk add openssl-dev 
RUN apk add net-tools 

WORKDIR /go/src/app

RUN apk update && apk add --virtual build-dependencies build-base gcc wget git


RUN git config --global http.sslVerify false
EXPOSE 3000



