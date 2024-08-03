# URL Shortener

Welcome to the URL Shortener project! This application allows you to shorten long URLs into concise, easy-to-share links and seamlessly redirect users to the original web addresses. Built with Go and deployed using Docker and Kubernetes, this project is designed to be lightweight, efficient, and easy to use.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Running Locally](#running-locally)
- [Deploying to Kubernetes](#deploying-to-kubernetes)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Shorten URLs**: Convert long URLs into short, user-friendly links.
- **Redirect URLs**: Access original URLs by using the shortened links.
- **Easy Deployment**: Deploy the application using Docker and Kubernetes.
- **Scalable**: Built to handle a large number of URL mappings.
- **Persistent Storage**: Utilizes SQLite for storing URL mappings.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following software installed on your system:

- [Go](https://golang.org/dl/) (version 1.21 or newer)
- [Docker](https://www.docker.com/get-started)
- [Kubernetes](https://kubernetes.io/docs/tasks/tools/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/url-shortener.git
   cd url-shortener
   ```
2. **Set up go modules**

    ```bash
    go mod tidy
    ```
## Usage

### Running Locally

You can run the application locally without docker using the Go command line tools.

1. **Run the application**
    ```bash 
    go run cmd/main.go
    ```
2. **Test the application**
    * Shorten a URL:
   ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"original_url":"https://example.com"}' http://localhost:8080/shorten
   ```
   * Access the shortened URL:
   ```bash
   curl -I http://localhost:8080/<shortID>
   ```
### Running with Docker

1. **Build the Docker image**
    ```bash 
    docker build -t url-shortener:latest .
    ```
2. **Run the Docker container**
    ```bash 
    Docker run -p 8080 url-shortener:latest
    ```
3. **Access the application**
Use the same curl commands as above to test the application

### Deploying to Kubernetes
To Deploy the URL shortener application on a Kubernetes cluster, follow these steps:

1. **Apply Kubernetes configurations**
    ```bash 
    kubectl apply -f kubernetes/deployment.yaml
   kubectl apply -f kubernetes/service.yaml
    ```
2. **Verify deployment**
    ```bash 
    kubectl get pods
   kubectl get svc
    ``` 
3. **Access the application**
Use the external IP provided by the Kubernetes service to interact with the application.

## Contributing
I welcome contributions from the community! If you'd like to contribute, please follow these steps:
   
   1. Fork the repository
   2. Create a new branch for your feature or bugfix
   3.  Make your changes
   4. Submit a pull request

Please ensure your code adheres to the project's coding standards and includes relevant tests

### License
This Project is licensed under the MIT License. See the [LICENSE](https://github.com/git/git-scm.com/blob/main/MIT-LICENSE.txt) file for details.

Thank you for using the URL Shortener project! If you have questions or feedback, feel free to open an issue or submit a pull request.