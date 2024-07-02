<h1 align="center">Todo Space 📑</h1>

## 💡 Objective
    Develop a TODO API using Golang and ScyllaDB that supports basic CRUD operations and includes 
    pagination functionality for the list endpoint.

## 📒 Requirements
- Set up a Golang project and integrate ScyllaDB as the database for storing TODO items. Ensure that items in the database are stored user-wise.
- Implement endpoints for creating, reading, updating, and deleting TODO items for a single user at a time. Each TODO item should have at least the following properties: id, user_id, title, description, status, created, updated.
- Implement a paginated list endpoint to retrieve TODO items.
- Provide support for filtering based on TODO item status (e.g., pending, completed).


## ⚙️ API Documentation

### `POST` /v1/todo

Create a new todo item.

-   **URL:** `http://localhost:8080/v1/todo`
-   **Method:** POST
-   **Body:** json
    ```json
    {
        "id": "004599c1-69d7-47b8-8306-a144d7265538",
        "user_id": "004599c1-69d7-47b8-8306-a144d7265538",
        "title": "Complete project tasks",
        "description": "Finish coding the backend and write documentation",
        "status": "Complete",
        "created": "2024-07-02T12:00:00Z",
        "updated": "2024-07-02T12:00:00Z"
    }



### `PUT` /v1/todo/:id

Update a specific todo item.

-   **URL:** `http://localhost:8080/v1/todo/:id`
-   **Method:** PUT
-   **URL Parameters:**
    -   `id`: ID of the todo item to update



### `DELETE` /v1/todo/:id

Delete a specific todo item.

-   **URL:** `http://localhost:8080/v1/todo/:id`
-   **Method:** DELETE
-   **URL Parameters:**
    -   `id`: ID of the todo item to delete



### `GET` /v1/todo/:id

Retrieve details of a specific todo item.

-   **URL:** `http://localhost:8080/v1/todo/:id`
-   **Method:** GET
-   **URL Parameters:**
    -   `id`: ID of the todo item to retrieve



### `GET` /v1/todo

Retrieve all todo items based on filters.

-   **URL:** `http://localhost:8080/v1/todo`
-   **Method:** GET
-   **Query Parameters:**
    -   `status`: Filter by status (`pending` or `completed`)
    -   `size`: Number of items to retrieve (default 10)
    -   `lastPageToken`: Offset for pagination (default 0)



## 🏃‍♂️ Run Locally

- Clone this repository.
    ```js
    https://github.com/srivalli123125/Todo-Space.git
    ```
- Install Docker Desktop.
- Run `docker-compose -f todospace-api.yml up` to initialize a Scylla-DB instance running on port 9042.
- Go from the main directory to `\cmd\api` and then run `go run main.go` to start the service.

Note: If you are facing issue while connecting the golang backend services with the ScyllaDB then run the below steps.
  - open ``cmd`` and enter the command to execute the cqlsh in the scylla-db instance.
  ```bash
    docker exec -it scylla-db cqlsh
  ```
  

  Note down the ``IP(172.23.0.2)``(might be change according to your docker configuration) and change it to the ``scylladb.go`` file.


  
  - Now in the cqlsh terminal create the ``KEYSPACE`` manually by the following command:
  ```bash
    CREATE KEYSPACE todo WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
  ```
  - Now enter the command to use the ``todo KEYSPACE``
  ```bash
    use todo;
  ```
  - Now create the ``todos`` table from the below query
  ```sql
    CREATE TABLE IF NOT EXISTS todos (id UUID,user_id UUID,title TEXT,description TEXT,status TEXT,created TEXT,updated TEXT,PRIMARY KEY (id, user_id));
  ```

## Features Implemented:

- Implemented CRUD routes for interaction b/w server and ScyllaDB.
- The API's are paginated for easy data retrieval.
- The application's DB part is Dockerized and is stateful through volumes.
- Support for filtering based on TODO item status (e.g., pending, completed).

## Current Architecture:

- Containerized approach to solving the problem statement.
- Two Interfaces one for the server and one for db are interacting between each other for the backend application.

## Future Scope:

- The current architecture is a very basic implementation of the problem statement.
- Depending upon the scale, the entire architecture can be **scaled horizontally** using nginx load balancing.
- Web can use a queueing mechanism like Rabbit or BullMQ to introduce pub-sub architecture to improve performance.
- The Go-Server could be containerized to improve deployment.
- Introduction to goroutines would increase the overall throughput of the service.  

## Design  Designs

1.Database Choice:

ScyllaDB: Chosen for its scalability and performance characteristics suitable for high-volume applications like TODO management. It's a NoSQL database designed to handle large amounts of data with low latency.

2.Go Language:

Efficiency: Go was chosen due to its efficiency in building concurrent and scalable applications. It integrates well with ScyllaDB via the gocql library, providing robust support for Cassandra-compatible databases.

3.RESTful API:

Endpoint Design: Followed RESTful conventions for CRUD operations (POST /todos, GET /todos/{id}, etc.). Pagination (GET /todos) and filtering (status={status}) were implemented as query parameters to provide flexibility in querying TODO items.

4.Error Handling:

Consistent Error Responses: HTTP status codes and descriptive error messages are returned for better API usability and debugging.

5.Concurrency and Performance:

Goroutines: Utilized Go's concurrency model (goroutines) for handling concurrent requests efficiently, especially beneficial in scenarios with multiple users interacting simultaneously.

6.Security:

Authentication: Authentication mechanisms (not shown in this basic example) should be implemented to secure endpoints, ensuring only authorized users can perform operations on their TODO items.

## 👨🏻‍💻 Developer's Talk


Developed by <a href="https://github.com/srivalli123125">Srivalli Bollineni</a>


<a href="https://github.com/srivalli123125/Todo-Space">This</a> is a small effort from my side to build a small scale project using Golang and ScyllaDB technologies. The experience taught me so many things, as well as the challenges involved in overcoming problems encountered during the development phase. I consider this project very relevant to me as a full-stack developer.

<br/>

---