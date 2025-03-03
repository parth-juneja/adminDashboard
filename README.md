# Admin Dashboard

This project is an Admin Dashboard built with Go. It provides a backend CRUD API to manage customers (here specifically in a gym membership).

I have made use of Repository pattern to isolate the DB choice to the user, 
I am currently using Postgres and pgadmin via Docker containers, but ultimately it's upto you whichever DB you're comfortable with.

Also attached is a sample Curl file to test API via Postman, just paste the Curls and enjoy!

Before running, create a .env file with the appropriate DB connection keys and create a container running postgres and using pgadmin, create the necessary table, i have not included low level steps as these are trivial and the service should only have single responsibility.


## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/adminDashboard.git
    ```
2. Navigate to the project directory:
    ```sh
    cd adminDashboard
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```
