FROM golang:latest
WORKDIR /app
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite
COPY . .
RUN go get -d
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]