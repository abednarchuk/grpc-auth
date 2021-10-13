FROM golang:1.17.2

WORKDIR /auth_backend

COPY . .

RUN go mod tidy

EXPOSE 50051

CMD [ "go","run","." ]