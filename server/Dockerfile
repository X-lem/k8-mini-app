# Use an official Golang runtime as a parent image
FROM golang:latest AS build-stage

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application inside the container
RUN go build -o /k8-mini-app

FROM gcr.io/distroless/base-debian12 AS run-stage
WORKDIR /app

# Copy executable file from build-stage image into run-stage image
COPY --from=build-stage /k8-mini-app /k8-mini-app

# Expose port 80
EXPOSE 80

# Run the Go application when the container starts
CMD ["/k8-mini-app"]