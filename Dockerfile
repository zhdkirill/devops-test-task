# Compile the Go app using Alpine Golang image
FROM golang:alpine as builder
RUN mkdir /build/
ADD /app/ /build/
WORKDIR /build/
RUN go build

# Create smaller image without build runtime
FROM alpine
# Create user to run the app under
RUN mkdir /app/ ; adduser -S -D -H -h /app/ ervcp
USER ervcp
COPY --from=builder /build/ /app/
WORKDIR /app/
EXPOSE 8080
ENTRYPOINT ["./ervcp"]
