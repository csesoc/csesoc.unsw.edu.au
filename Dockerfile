# Run base image from node
FROM node:latest AS builder

# Copy files into container
COPY . /~/app/

# Set working frontend directory
WORKDIR /~/app/frontend

# Install yarn build dependencies
RUN yarn
RUN yarn build

# Multi stage build to enable code for static files to be served
FROM golang:1.13-buster AS server

# Copy the files from stage builder into this stage
COPY --from=builder /~ /~

# Set the working backend directory
WORKDIR /~/app/backend/server/

# Build dependencies 
RUN go mod download

# Build and run the server
RUN go build -o  main .
EXPOSE 1323
CMD [ "./main" ]
