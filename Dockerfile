FROM golang:1.22-alpine AS builder

#Create and change to the app dire
WORKDIR /app

COPY go.* ./
RUN go mod download

# Copy local code to container image
COPY . .


#Build the binary file
WORKDIR cmd/api
RUN CGO_ENABLED=0 GOOSE=linux go build -mod=readonly -v -o main

FROM alpine:latest
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
ARG SERVICE
COPY --from=builder /app/cmd/api/main /main

# Run the web service on container startup.
CMD ["/main"]