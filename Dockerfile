# Use a minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary build file from the root directory into the container
COPY main .

# Set the entry point for the container
CMD ["./main"]