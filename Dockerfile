FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y wget git gcc

RUN wget -P /tmp https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.11.5.linux-amd64.tar.gz
RUN rm /tmp/go1.11.5.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ADD . /code
WORKDIR /code
RUN ["apt-get", "install", "-y", "libgl1-mesa-dev", "libasound2-dev", "xorg-dev"]
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 1989
RUN ["ls"]
CMD ["gbc-in-cloud", "-s", "-c", "gamelist.json"]
