# Rate Limiter in Go

This project implements a basic **Rate Limiter** in Go, using the `golang.org/x/time/rate` package for controlling the rate of incoming HTTP requests. The rate limiter allows up to **2 requests per second** with a **burst size of 4** requests. If the rate limit is exceeded, the server responds with a `429 Too Many Requests` status and a relevant message.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Rate Limiter Behavior](#rate-limiter-behavior)
- [Testing](#testing)

## Installation

1. Make sure you have **Go** installed on your machine. If you don't, you can install Go from the [official Go website](https://golang.org/dl/).
   
2. Clone the repository to your local machine:

    ```bash
    git clone <repository_url>
    cd <project_directory>
    ```

3. Install dependencies (e.g., `golang.org/x/time/rate`):

    ```bash
    go get golang.org/x/time/rate
    ```

4. Start the Go application:

    ```bash
    go run main.go
    ```

## Usage

- Once the server is running on `localhost:8080`, you can access the `/ping` endpoint:

    ```bash
    curl http://localhost:8080/ping
    ```

    This will return a JSON response with a `200 OK` status if the request is within the allowed rate limit.

## Rate Limiter Behavior

- **Rate Limit Configuration**:
  - Up to **2 requests per second**.
  - **Burst size**: 4 requests in rapid succession.
  
- **Rate Limiting**:
  - After exceeding the limit, further requests will receive a **429 Too Many Requests** response.
  
- **Example Response on Success** (`200 OK`):
  
    ```json
    {
      "status": "Success",
      "body": "Hi, Your Server is Up and Running"
    }
    ```

- **Example Response on Rate Limiting** (`429 Too Many Requests`):

    ```json
    {
      "status": "Request Failed",
      "body": "Reached API Limits, try again later."
    }
    ```


## Testing

You can run tests for the application by using the following command:

```bash
go test
