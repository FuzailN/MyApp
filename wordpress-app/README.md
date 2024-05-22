# WordPress Application

<img width="1435" alt="image" src="https://github.com/FuzailN/MyApp/assets/129302212/f18cf070-4150-4d01-8945-06332a40a524">

# WordPress Local Development Setup

This guide provides instructions on how to set up and run a WordPress site locally using Docker Compose.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Setup and Installation

### **Clone the Repository:**
```
git clone https://github.com/your-username/your-repository.git

cd wordpress-app
```

### Running the Application
Build the Docker Image:
From the go-app directory, build the Docker image:
```
docker build -t your-dockerhub-username/wordpress-app:latest .
```

### Start Docker Compose Services:
```
docker-compose up -d
```
This will start the application on port 8081.

### Deployment
Open your web browser and navigate to (http://localhost:8081) to see the application running.
