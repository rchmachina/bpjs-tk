# Use an official Node runtime as a base image
FROM node:18-alpine as builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./


# Copy the local code to the container
COPY . .

EXPOSE 3000

# Install dependencies
RUN npm install
