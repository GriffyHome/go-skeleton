# Go Project Skeleton

This repository provides a skeleton structure for Go projects. It includes a basic project setup with necessary configurations, dependencies, and database connection setups for SQL and Cassandra. Users can clone this repository, change the module name, install dependencies, and remove unnecessary components to fit their needs.

## Getting Started

Follow the steps below to get your project up and running.

### Prerequisites

- Go (version 1.16 or higher)
- Git

### Clone the Repository

1. Clone the repository to your local machine:
    ```sh
    git clone https://github.com/your-username/go-project-skeleton.git
    cd go-project-skeleton
    ```

### Change the Module Name

2. Change the module name to your desired module name. Open the `go.mod` file and replace `your-module-name` with your desired module name:
    ```go
    module your-module-name
    ```

3. Run the following command to update the module name throughout the project:
    ```sh
    go mod tidy
    ```

### Install Dependencies

4. Install the required dependencies:
    ```sh
    go mod download
    ```

### Database Setup

5. The `pkg/db` folder contains configurations for different databases (SQL and Cassandra). Remove any unnecessary database configurations that you do not need for your project.

    - To remove the SQL configuration:
        ```sh
        rm -rf pkg/db/sql
        ```

    - To remove the Cassandra configuration:
        ```sh
        rm -rf pkg/db/cassandra
        ```
