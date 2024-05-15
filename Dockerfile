# https://docs.docker.com/language/golang/build-images/
FROM golang:1.21.4

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

COPY . .
# COPY *.go ./
# COPY ./**/*.go ./
RUN go build -o /griz

EXPOSE 8080

CMD ["/griz"]