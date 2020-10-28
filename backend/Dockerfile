# Base image changed to go to run Go commands
FROM golang:1.13-buster as production-stage

# Set working server directory
WORKDIR /app/server

# Copy server files 
COPY ./ /app/server

# Build go dependencies
RUN go mod download

# Generate swag documentation
RUN go get github.com/swaggo/swag/cmd/swag
RUN swag init

# Daemon to run go build automatically when new files are changed on host system
RUN go get github.com/githubnemo/CompileDaemon

# Expose port for binding
EXPOSE 1323

# Install dependencies
RUN apt-get update
RUN apt-get install -y libkrb5-dev

# Manually install MongoDB Tools
RUN mkdir -p $GOPATH/src/github.com/mongodb
RUN cd $GOPATH/src/github.com/mongodb
RUN git clone https://github.com/mongodb/mongo-tools
RUN cd mongo-tools
RUN export GOROOT=/usr/local/go
RUN ./make build

# Run go build and execute the ./main file generated
ENTRYPOINT CompileDaemon --build="go build -o main ." --command="./main"