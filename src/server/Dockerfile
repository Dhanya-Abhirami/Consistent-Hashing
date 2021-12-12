FROM golang:alpine

WORKDIR /server

RUN  go mod init server

RUN go get github.com/gin-gonic/gin

RUN go get github.com/swaggo/swag/cmd/swag

RUN go get github.com/swaggo/gin-swagger

RUN go get github.com/swaggo/files

COPY ./ ./

RUN swag init

RUN go build

CMD ["go", "run", "."]

EXPOSE 8080