# Movie API

This is a simple Movie API implemented in Go. It allows you to perform basic CRUD (Create, Read, Update, Delete) operations on a collection of movies. Instead of using a fully-fledged database, the API uses structs and slices to store and manipulate movie data.

## Getting Started

To run the Movie API locally, follow the steps below:

1. Clone the repository:
   ```
   git clone https://github.com/DragonCoderz/Go-Movies-CRUD-API
   ```

2. Install the dependencies:
   ```
   go mod download
   ```

3. Build and run the application:
   ```
   go run main.go
   ```

The API server will start running on `http://localhost:8000`.

## Endpoints

The following endpoints are available in the API:

### Get All Movies

- URL: `/movies`
- Method: `GET`
- Response: Returns a JSON array containing all the movies.

### Get a Movie

- URL: `/movies/{id}`
- Method: `GET`
- Response: Returns the details of a specific movie identified by `{id}`.

### Create a Movie

- URL: `/movies`
- Method: `POST`
- Request Body: Send the movie details as JSON in the request body.
- Response: Returns the details of the created movie.

### Update a Movie

- URL: `/movies/{id}`
- Method: `PUT`
- Request Body: Send the updated movie details as JSON in the request body.
- Response: Returns the updated movie details.

### Delete a Movie

- URL: `/movies/{id}`
- Method: `DELETE`
- Response: Deletes the movie identified by `{id}`.

## Data Structures

The movie data is stored in memory using the following structs:

```go
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
```

The `Movie` struct represents a movie object with an ID, ISBN, title, and director details. The `Director` struct contains the first name and last name of the movie director.

## Dependencies

The following third-party packages are used in this project:

- `github.com/gorilla/mux`: A powerful URL router and dispatcher for building HTTP services in Go.

## Contributing

Contributions are welcome! If you find any issues or want to enhance the Movie API, feel free to submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
