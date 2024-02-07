FROM golang:alpine

# Create app directory
WORKDIR /app

# Bundle app source
COPY . .

RUN go build -o main .

CMD ["./main"]