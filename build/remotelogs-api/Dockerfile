FROM    golang:latest

ENV     GOOS=linux      

RUN     mkdir /app

ADD     . /app

WORKDIR /app

RUN     go build -o bin/remotelogs-api cmd/remotelogs-api/*.go

EXPOSE  3000

CMD     ["/app/bin/remotelogs-api"]