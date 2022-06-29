package utils

import (
	"database/sql"
	"nethttpcrud/db"
	"nethttpcrud/lib"
	"nethttpcrud/model"
)

func GetMoviesDetail(db *sql.DB) (movies []model.Movie) {

	rows, err := db.Query("SELECT * FROM movies")
	// check errors
	lib.CheckErr(err)
	for rows.Next() {
		var (
			id        int
			movieID   string
			movieName string
		)

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		lib.CheckErr(err)

		movies = append(movies, model.Movie{MovieID: movieID, MovieName: movieName})
	}
	return

}

func InsertMovieData(movieRequest model.Movie) {

	db := db.SetupDB()
	var lastInsertID int
	lib.PrintMessage("inserting the row")
	err := db.QueryRow("Insert into movies(movieid,moviename) values($1,$2) returning id", movieRequest.MovieID, movieRequest.MovieName).Scan(&lastInsertID)
	lib.CheckErr(err)

}

func UpdateMovie(movieRequest model.Movie, Id string) {

	db := db.SetupDB()
	_, err := db.Exec("update movies set movieid =$1,moviename=$2 where id =$3 ", movieRequest.MovieID, movieRequest.MovieName, Id)
	lib.CheckErr(err)
}

func DeleteMovie(Id string) {

	db := db.SetupDB()
	_, err := db.Exec("Delete from movies where id =$1 ", Id)
	lib.CheckErr(err)

}
