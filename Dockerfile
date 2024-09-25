FROM golang:1.12-alpine

# Add git, bash and dependencies for building
RUN apk add --no-cache git gcc musl-dev


WORKDIR /app

COPY . .

RUN go mod download
RUN go build main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]