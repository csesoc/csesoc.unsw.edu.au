# Use this Dockerfile for production builds
# Use node image to install Vue dependencies
FROM node:14.2 as vue-dependencies

# Set the working directory for the project in container
WORKDIR /app 

# Install Dependencies
COPY ./frontend ./
RUN yarn build

# At this stage the dist/ folder should be in the container
# and hence static files will be served from dist
FROM golang:1.13-buster

# Set the working directory for the project in container
WORKDIR /app 

# Copy dist/ folder
COPY --from=vue-dependencies /app/dist/ ./dist

# Copy server files 
COPY ./backend ./

# Build go dependencies
RUN go mod download

# Expose port for binding
EXPOSE 1323

# Build the server binary
RUN go build

# Run the server
ENTRYPOINT ./m
