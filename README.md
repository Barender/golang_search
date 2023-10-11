## Installation

To install this project, follow these steps:

1. Clone the repository to your local machine.
2. Install go 1.21 in your system.
3. Install the necessary dependencies by running `go mod download`.
4. Create a `.env` file in the root directory of the project and add the following environment variables:
   - `MONGO_URI`: the URI for your MongoDB database
   - `MONGO_DBNAME`: the name of the database you want to use
5. Start the server by running `go run main.go`.

## Running the project

To run the project, follow these steps:

1. Start the server by running `go run main.go`.
2. Open your web browser and navigate to `http://localhost:4000`.
3. You should see the homepage of the project.

This project is a simple REST API that allows you to retrieve and add products to a MongoDB database. The API has two endpoints: `/getAllProducts` and `/getFilteredProducts`. The `/products` endpoint returns a list of all products in the database, while the `/getFilteredProducts` endpoint returns matching products with given filtered request body.
