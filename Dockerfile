# Run base image from node
FROM node:latest AS builder

# Set working directory
WORKDIR /~/app

# Copy files into container
COPY . /~/app

# Install yarn and build dependencies
RUN curl -o- -L https://yarnpkg.com/install.sh | bash
RUN yarn
RUN yarn build

# Multi stage build to enable code for static files to be served
FROM golang:1.13-buster

# Copy the files from stage builder into this stage
COPY --from=builder /~ /~

# Set the working directory
WORKDIR /~/app/server/

# Build dependencies 
RUN go mod download

# Build and run the server
RUN go build -o  main .
EXPOSE 1323
CMD [ "./main" ]
