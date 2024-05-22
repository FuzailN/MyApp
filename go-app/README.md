# Go Application

<img width="1440" alt="image" src="https://github.com/FuzailN/MyApp/assets/129302212/61cf4a73-0738-4fa7-ad0e-48887cde99e9">

# Instructions on how to set up and run the extended application locally. 
## Table of Contents
- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Setup and Installation](#setup-and-installation)
- [Running the Application](#running-the-application)
- [Deployment](#deployment)

## Introduction
This project is an extended Golang application that includes a setup for linting, building a Docker image, and deploying using Docker Compose. This README provides instructions on how to set up and run the application locally.

## Prerequisites
Before you begin, ensure you have the following installed on your machine:
```
Golang
Docker
Docker Compose
```
## Setup and Installation
Clone the Repository:
```
git clone https://github.com/FuzailN/MyApp
cd go-app
```

## Install Dependencies: (Optional)
Navigate to the go-app directory and download the Go module dependencies:
```
cd go-app
go mod download
```

## Running the Application
### Build the Docker Image:
From the go-app directory, build the Docker image:
```
docker build -t your-dockerhub-username/golang-app:latest .
Start Docker Compose Services:
```
## Docker Compose to start the services
```
docker-compose up -d
```
This will start the application on port 8082.
## Access the Application:

## Deployment
Open your web browser and navigate to (http://localhost:8082) to see the application running.


