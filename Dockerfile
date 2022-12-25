FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY *.go ./

COPY . .
RUN go build -o /datacollector

CMD [ "/datacollector" ]