# # Base image node to install Vue dependencies
# FROM node:14.2 AS build-stage

# WORKDIR /app

# # Install dependencies in separate step to cache
# COPY package.json ./
# RUN yarn install

# # Build dependencies into dist folder, copy line may change as we restructure
# COPY . .
# RUN yarn serve

# Base image changed to go to run Go commands
FROM golang:1.13-buster as production-stage

# Copy the dist folder into the new container
# COPY --from=build-stage /app/dist  /app/dist
WORKDIR /app/server

# Copy server files 
COPY ./ /app/server

# Build go dependencies
RUN go mod download

# Daemon to run go build automatically when new files are changed on host system
RUN go get github.com/githubnemo/CompileDaemon

# Expose port for binding
EXPOSE 1323

# Run go build and execute the ./main file generated
ENTRYPOINT CompileDaemon --build="go build -o main ." --command=./main