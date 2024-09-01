# Car-Rental

This project is a Car Rental API built with Golang and PostgreSQL. The API supports CRUD operations for customers, cars, and bookings, as well as additional features such as membership programs, car rentals with drivers, and driver incentives.

## Project Overview

The Car Rental API is developed in two versions:

- Car Rental - v1: Basic CRUD operations for customers, cars, and bookings.
- Car Rental - v2: Enhancements including membership discounts, rentals with drivers, and incentive calculation for drivers

## Installation

- Clone the repository:

```bash
git clone https://github.com/mnaufalhilmym/Car-Rental.git
cd Car-Rental
```

- Install the dependencies:

```bash
go mod tidy
```

## Running The Application

- Ensure that you have created and configured `config.yml`. An example configuration file can be found in `config_*.yml`.

- Run the application

```bash
go run ./cmd
```

## Build and Running with Docker
- Build the Docker image:
```bash
docker build . -f Dockerfile -t docker.io/mnaufalhilmym/car-rental
```
- Run the Docker container:
```bash
docker run -p 8080:8080 --name car-rental -v $(pwd)/config.yml:/config.yml docker.io/mnaufalhilmym/car-rental
```
- Run with Docker Compose:
```bash
docker-compose -f compose.yml up
```

## API Documentation

[Postman API Documentation](https://www.postman.com/cloudy-satellite-99691/workspace/car-rental/collection/38046449-d24b9de5-8ffa-44a7-b964-651584c98585?action=share&creator=38046449)

## Database Design

ERD (Entity Relationship Diagram)

- Car Rental - v1:\
  ![ERDv1](./docs/erd-v1.png)

- Car Rental - v2:\
  ![ERDv2](./docs/erd-v2.png)
