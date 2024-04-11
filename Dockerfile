FROM golang:1.22
LABEL authors="Bastian Venz"

WORKDIR /app
COPY go.mod Lets.go ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /goTimeouting


EXPOSE 8080
CMD ["/docker-gs-ping"]