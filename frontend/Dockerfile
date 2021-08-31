# Base image node to install Vue dependencies
FROM node:14.2 AS build-stage

WORKDIR /app

# Install dependencies in separate step to cache
COPY package.json ./
RUN yarn install

# Build dependencies into dist folder, copy line may change as we restructure
COPY . .
gi
ENTRYPOINT yarn serve build