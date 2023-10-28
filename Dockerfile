FROM golang:1.20.5 AS build-stage

WORKDIR /app
# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /auservice


# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY --from=build-stage /auservice /auservice
EXPOSE 8000
ENTRYPOINT ["/auservice"]