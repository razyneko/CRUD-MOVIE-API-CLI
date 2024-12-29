// Seed initial data for movies
package data

import "github.com/movie-api/models"

var Movies = []models.Movie{
    {
        ID:    "1",
        Isbn:  "1234567890",
        Title: "The Great Adventure",
        Director: &models.Director{
            FirstName: "John",
            LastName:  "Doe",
        },
    },
    {
        ID:    "2",
        Isbn:  "2345678901",
        Title: "Mystery at Dawn",
        Director: &models.Director{
            FirstName: "Jane",
            LastName:  "Smith",
        },
    },
	{
        ID:    "3",
        Isbn:  "234573458901",
        Title: "Mystery Mary",
        Director: &models.Director{
            FirstName: "Mary",
            LastName:  "Jane",
        },
    },
	{
        ID:    "4",
        Isbn:  "23456785848901",
        Title: "Dawn of Justice",
        Director: &models.Director{
            FirstName: "Bruce",
            LastName:  "Wayne",
        },
    },
}
