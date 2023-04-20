# Flight Path

The Flight Path is a simple RESTful API developed in Go, designed to help users understand and track a particular person's flight path by querying a list of unordered flights. The API accepts a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
  - [Example Requests and Responses](#example-requests-and-responses)
- [Tests](#tests)
- [Production Considerations](#production-considerations)
- [Code Organization and Modularization](#code-organization-and-modularization)

## Features

- Efficiently calculates the flight path from a list of unordered flights
- Validates input data to avoid incorrect results
- Logs errors for easier debugging
- Returns a JSON response with the flight path or an error message
- Unit tests covering various use cases

## Installation

To install the Flight Path, you need to have [Go](https://golang.org/doc/install) installed on your machine. Then, follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/plar/flight-path.git
   ```

2. Change to the project directory:
   ```
   cd flight-path
   ```

3. Build the microservice:
   ```
   go build
   ```

4. Run the microservice:
   ```
   ./flight-path
   ```

The microservice will now be running on port 8080.

## Usage

### API Endpoints

#### POST /calculate

- `POST /calculate`: Calculates the flight path from a list of unordered flights.

The endpoint expects a request with a JSON payload containing a list of flights in the following format:

##### Request 

```json
[
  ["<source_airport_code>", "<destination_airport_code>"],
  ["<source_airport_code>", "<destination_airport_code>"],
  ...
]
```

Each flight is represented by an array of two strings: the source airport code and the destination airport code. 
The airport codes should be three-letter alphanumeric codes in uppercase format, according to the IATA airport code standard.

For example, a valid request body might look like:

```json
[
  ["LHR", "JFK"],
  ["JFK", "SYD"],
  ["SYD", "HKG"],
  ["HKG", "NRT"]
]
```

If the request body is not valid JSON or the format of the flight segments is incorrect, 
the API will return a `400 Bad/Request` error with an error message in the response body.

##### Response

If the API is able to find a valid flight path based on the input, 
it will return a `200/OK` response with a JSON payload containing the flight path in the following format:

```json
{
  "status": "success",
  "path": ["<start_airport_code>", "<end_airport_code>"]
}
```

The `path` field is an array of two strings representing the starting and ending airports of the flight path. 

If the API is unable to find a valid flight path, it will return a `400/Bad Request` error with an error message in the response body.

```json
{
   "status": "error",
   "message": "Failed to find flight path"
}
```

### Example Requests and Responses

1. **Request:**

   ```
   POST /calculate
   Content-Type: application/json

   [["SFO", "EWR"]]
   ```

   **Response:**

   ```json
   {
     "status": "success",
     "path": ["SFO", "EWR"]
   }
   ```

2. **Request:**

   ```
   POST /calculate
   Content-Type: application/json

   [["ATL", "EWR"], ["SFO", "ATL"]]
   ```

   **Response:**

   ```json
   {
     "status": "success",
     "path": ["SFO", "EWR"]
   }
   ```

3. **Request:**

   ```
   POST /calculate
   Content-Type: application/json

   [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
   ```

   **Response:**

   ```json
   {
     "status": "success",
     "path": ["SFO", "EWR"]
   }
   ```

4. **Request (Invalid Data):**

   ```
   POST /calculate
   Content-Type: application/json

   [["SFO"]]
   ```

   **Response:**

   ```json
   {
     "status": "error",
     "message": "Invalid request format"
   }
   ```

## Tests

To run the tests for the Flight Path, run the following command:

```
go test
```

This command will run all the unit tests, including various corner and invalid cases.

## Production Considerations

While the Flight Path Microservice provides a basic implementation for calculating flight paths from unordered flights, there are several aspects that should be addressed before deploying the service in a production environment:

1. **Security:** The service does not implement any authentication or authorization mechanisms. Depending on the use case, consider adding authentication (e.g., API tokens, OAuth) and authorization to control access to the API.

1. **Rate Limiting:** To protect the service from excessive usage or abuse, implement rate limiting to control the number of requests a client can make within a given timeframe.

1. **Logging and Monitoring:** Enhance logging and monitoring capabilities to capture more information about the service's performance, errors, and usage. This will aid in diagnosing issues and optimizing the service.

1. **Error Handling:** Although basic error handling and input validation are implemented, a more robust error handling approach should be considered to handle various edge cases and unexpected input formats.

1. **Documentation:** Consider using a tool like [Swagger](https://swagger.io/) to create interactive API documentation.

1. **Deployment and Infrastructure:** Set up a proper deployment process, including continuous integration and continuous deployment (CI/CD), and choose a suitable infrastructure solution, such as containerization (e.g., Docker) and orchestration (e.g., Kubernetes).

By addressing these considerations, you can ensure that the Flight Path is ready for production use and can handle real-world scenarios, providing a reliable and efficient service for users.

## Code Organization and Modularization

The current implementation of the Flight Path contains all the code in a single package, which may not be ideal for a production environment. To improve maintainability, readability, and scalability, consider organizing the code into packages and modules.

1. **Package Structure:** Organize the code into packages based on functionality. For example, create separate packages for handling API requests, flight path calculation logic, and utility functions. This separation of concerns will make it easier to understand, maintain, and test the code.

   ```
   flight-path/
   ├── api/
   │   ├── handlers.go
   │   └── responses.go
   ├── flightpath/
   │   ├── algorithm.go
   │   └── models.go
   ├── utils/
   │   └── validation.go
   └── main.go
   ```

2. **Modularize Functions:** Break down the code into smaller, more focused functions that can be tested and maintained independently. This will help to reduce complexity and make the code more modular, which is especially important as the service grows and evolves.

3. **Interfaces and Dependency Injection:** Use interfaces to define the behavior of different components and use dependency injection to improve testability and flexibility. This will allow you to swap out implementations easily, for example, when switching between different flight path calculation algorithms (or data storage solutions. or ...).

4. **Unit Tests:** Organize unit tests in a similar fashion, creating separate test files for each package and focusing on testing individual functions and components.

By organizing the code effectively into packages and modules, you can create a more maintainable and scalable Flight Path that is better suited for production use.
