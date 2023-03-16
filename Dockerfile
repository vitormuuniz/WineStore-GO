FROM golang
WORKDIR /app/src/winestore-go
ENV GOPATH=/app
COPY . /app/src/winestore-go/
RUN go mod tidy
RUN go build -o main .
CMD ["./main"]