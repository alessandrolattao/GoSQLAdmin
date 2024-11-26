# Use the official Golang image as the base
FROM golang:1.23

# Install gow (a tool for live reloading)
RUN go install github.com/mitranim/gow@latest

# Set the working directory inside the container
WORKDIR /app

# Expose the port your application will run on
EXPOSE 8080

# Set the entrypoint to gow, watching for changes in the application
CMD ["gow", "-e=go,mod,html,png,css", "run", "cmd/gomyadmin/main.go"]

