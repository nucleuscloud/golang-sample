FROM golang:1.18

WORKDIR /workplace

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY ./entrypoint.sh ./

RUN chmod u+x ./entrypoint.sh

COPY ./ ./

RUN go build -v -o bin/main .

ENTRYPOINT [ "/workplace/entrypoint.sh" ]
