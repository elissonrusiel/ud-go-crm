# Project Title

This project is a CRM API for Udacity Go Developer course.

## Installation 

Describe the installation process here. This might include:

1. Cloning the repo: `git clone https://github.com/elissonrusiel/ud-go-crm.git`
2. Navigating into the project directory
3. Run project with the command: `go run main.go`

## Build

To build the project, run the following command:

```bash
go build
```

## Run

To run the project, run the following command:

```bash
go run main.go
```

## Usage

1. Run the application: `go run main.go`
2. Open documentation in a browser: `localhost:3000`


## Project Rubric: CRM Backend 

## Best Practices

| Criteria                               | Submission Requirements                                                                                                                                                                                                                                                                         |
|----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Set up a proper Go environment**     | The project requires only a simple `go run` command to launch the application.                                                                                                                                                                                                                  |
| **Document the project in the README** | The project README contains a description of the project, and also includes instructions for installation, launch, and usage.                                                                                                                                                                   |
| **Organize and write clean code**      | - Syntax and semantics of language features (e.g., functions, variables, loops, etc.) are well-formed and free from errors or warnings in the console <br> - Data structures, handlers, routes, imports, and other assets are organized logically (e.g., grouped together) and are easy to find |
| **Build an intuitive user experience** | Users can interact with the application (i.e., make API requests) by simply using Postman or cURL.                                                                                                                                                                                              |

## Data

| Criteria                                            | Submission Requirements                                                                                                                                                                                                                                              |
|-----------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Create a Customer struct**                        | Each customer includes: <br> - ID <br> - Name <br> - Role <br> - Email <br> - Phone <br> - Contacted (i.e., indication of whether or not the customer has been contacted) <br><br> Data is mapped to logical, appropriate types (e.g., Name should not be a `bool`). |
| **Create a mock "database" to store customer data** | Customers are stored appropriately in a basic data structure (e.g., slice, map, etc.) that represents a "database."                                                                                                                                                  |
| **Seed the database with initial customer data**    | The "database" data structure is non-empty. That is, prior to any CRUD operations performed by the user (e.g., adding a customer), the database includes at least three existing (i.e., "hard-coded") customers.                                                     |
| **Assign unique IDs to customers in the database**  | Customers in the database have unique ID values (i.e., no two customers have the same ID value).                                                                                                                                                                     |

## Server

| Criteria                                                     | Submission Requirements                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
|--------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Serve the API locally**                                    | The API can be accessed via `localhost`.                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| **Create RESTful server endpoints for CRUD operations**      | The application handles the following 5 operations for customers in the "database": <br> - **Getting a single customer** through a `/customers/{id}` path <br> - **Getting all customers** through the `/customers` path <br> - **Creating a customer** through a `/customers` path <br> - **Updating a customer** through a `/customers/{id}` path <br> - **Deleting a customer** through a `/customers/{id}` path <br><br> Each RESTful route is associated with the correct HTTP verb. |
| **Return JSON in server responses**                          | The application leverages the `encoding/json` package to return JSON values (i.e., not text, etc.) to the user.                                                                                                                                                                                                                                                                                                                                                                           |
| **Serve static HTML at the home ("/") route**                | The home route is a client API endpoint, and includes a brief overview of the API (e.g., available endpoints). Note: This is the only route that does not return a JSON response.                                                                                                                                                                                                                                                                                                         |
| **Set up and configure a router**                            | The application uses a router (e.g., `gorilla/mux`, `http.ServeMux`, etc.) that supports HTTP method-based routing and variables in URL paths.                                                                                                                                                                                                                                                                                                                                            |
| **Create and assign handlers for requests**                  | The Handler interface is used to handle HTTP requests sent to defined paths. There are five routes that return a JSON response, and are each is registered to a dedicated handler: <br> - `getCustomers()` <br> - `getCustomer()` <br> - `addCustomer()` <br> - `updateCustomer()` <br> - `deleteCustomer()`                                                                                                                                                                              |
| **Includes basic error handling for non-existent customers** | If the user queries for a customer that doesn't exist (i.e., when getting a customer, updating a customer, or deleting a customer), the server response includes: <br> - A `404` status code in the header <br> - `null` or an empty JSON object literal or an error message                                                                                                                                                                                                              |
| **Set headers to indicate the proper media type**            | An appropriate `Content-Type` header is sent in server responses.                                                                                                                                                                                                                                                                                                                                                                                                                         |
| **Read request data**                                        | The application leverages the `io/ioutil` package to read I/O (e.g., request) data.                                                                                                                                                                                                                                                                                                                                                                                                        |
| **Parse JSON data**                                          | The applications leverages the `encoding/json` package to parse JSON data.                                                                                                                                                                                                                                                                                                                                                                                                                |

## Suggestions to Make Your Project Stand Out

1. Create an additional endpoint that updates customer values in a batch (i.e., rather than for a single customer).
2. Upgrade the mock database to a real database (e.g., PostgreSQL).
3. Deploy the API to the web.