# Base image node to install Vue dependencies
FROM node:14.13.1 AS build-stage

WORKDIR /app

# Install dependencies in separate step to cache
COPY package.json ./
RUN yarn install

# Build dependencies into dist folder, copy line may change as we restructure
COPY . .
RUN yarn build
RUN npm ci -g serve
ENTRYPOINT serve -s dist

