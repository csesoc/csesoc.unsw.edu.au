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

# # Run python and run server
# EXPOSE 8080
# CMD [ "python3", "-m", "http.server", "8080"]

FROM golang:1.13-buster

COPY --from=builder /~ /~

WORKDIR /~/app/server/

RUN go mod download

# # Build and run the server
RUN go build -o  main .
EXPOSE 1323
CMD [ "./main" ]
