# Stress-Tool

`stress-tool` is a simple and effective tool for performing stress tests on web applications. It sends a specified number of requests to a given URL with configurable concurrency.

## Features

- Send a large number of HTTP requests to a target URL.
- Configure the number of concurrent requests.
- Measure the performance and stability of your web application under load.

## Requirements

- Docker

## Installation

To use `stress-tool`, you need to have Docker installed on your machine. If you don't have Docker installed, you can download and install it from the [official Docker website](https://www.docker.com/get-started).

## Usage

You can run the `stress-tool` using the following Docker command:

```sh
docker run --rm brunobevilaquaa/stress-tool --url=<URL> --requests=<REQUESTS> --concurrency=<CONCURRENCY>
```

### Parameters

- `--url`: The target URL to which the requests will be sent. (e.g., `https://www.google.com/`)
- `--requests`: The total number of requests to send. (e.g., `1000`)
- `--concurrency`: The number of concurrent requests to send. (e.g., `10`)

### Example

To perform a stress test on `https://www.google.com/` with 1000 total requests and 10 concurrent requests, run the following command:

```sh
docker run --rm brunobevilaquaa/stress-tool --url=https://www.google.com/ --requests=1000 --concurrency=10
```

This command will initiate the stress test, sending requests to the specified URL and providing output on the performance and stability of the web application under the specified load.

---

Thank you for using `stress-tool`! We hope it helps you ensure the reliability and performance of your web applications.