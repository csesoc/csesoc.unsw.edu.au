# Base image changed to go to run Go commands
FROM golang:1.13-buster as production-stage

# Copy server files 
COPY ./ /app/server

# Set working server directory
WORKDIR /app/server

# Build go dependencies
RUN go mod download

# Daemon to run go build automatically when new files are changed on host system
RUN go get github.com/githubnemo/CompileDaemon

# Expose port for binding
EXPOSE 1323

# Run go build and execute the ./main file generated
ENTRYPOINT CompileDaemon --build="go build -o main ." --command="./main"