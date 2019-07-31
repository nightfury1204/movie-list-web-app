package main

import "net/http"

var (
	searchRes = `{"Search":[{"Title":"Hello, My Name Is Doris","Year":"2015","imdbID":"tt3766394","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTg0NTM3MTI1MF5BMl5BanBnXkFtZTgwMTAzNTAzNzE@._V1_SX300.jpg"},{"Title":"Hello, Dolly!","Year":"1969","imdbID":"tt0064418","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BODJmZmFiNzQtMDJiYS00ZTgzLTljZGMtNjEzNzM4NmEyYjNiXkEyXkFqcGdeQXVyNjE5MjUyOTM@._V1_SX300.jpg"},{"Title":"Hello Ladies","Year":"2013–2014","imdbID":"tt2378794","Type":"series","Poster":"https://m.media-amazon.com/images/M/MV5BNjYxMjI3MzY3NF5BMl5BanBnXkFtZTgwMTgyNzg3MDE@._V1_SX300.jpg"},{"Title":"Hello I Must Be Going","Year":"2012","imdbID":"tt2063666","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMzkzMDc0Nzg5OF5BMl5BanBnXkFtZTcwMDU0MzAyOA@@._V1_SX300.jpg"},{"Title":"Hello Ladies: The Movie","Year":"2014","imdbID":"tt3762944","Type":"movie","Poster":"http://ia.media-imdb.com/images/M/MV5BMTQ5MjYxMjkwOV5BMl5BanBnXkFtZTgwODE3MjY0MzE@._V1_SX300.jpg"},{"Title":"Hello Brother","Year":"1999","imdbID":"tt0233856","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMjk1MDczMGQtY2RkNS00OGVhLWJhNzYtNWMwMzFhNTcyNjczXkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg"},{"Title":"Hello Ghost","Year":"2010","imdbID":"tt1848926","Type":"movie","Poster":"https://images-na.ssl-images-amazon.com/images/M/MV5BNDAyOTY2MzE4N15BMl5BanBnXkFtZTgwMjY0OTI5NDE@._V1_SX300.jpg"},{"Title":"Hello Again","Year":"1987","imdbID":"tt0093175","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTI1MDU1OTM2N15BMl5BanBnXkFtZTcwNjYzMjUyMQ@@._V1_SX300.jpg"},{"Title":"Hello Stranger","Year":"2010","imdbID":"tt1725995","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BYTE0Yzk1ZDEtN2E2Mi00Y2I0LTkzN2QtZjU4ODlhYTgxODgyXkEyXkFqcGdeQXVyNzI1NzMxNzM@._V1_SX300.jpg"},{"Title":"Oh, Hello on Broadway","Year":"2017","imdbID":"tt6987652","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BZmQ3YmM0NGMtYmRmNi00ZWY4LTk5MGYtYzUyODA4ODBlODE3XkEyXkFqcGdeQXVyMjQzNzk2ODk@._V1_SX300.jpg"}],"totalResults":"548","Response":"True"}`
	titleRes = `{"Title":"Hello","Year":"2008","Rated":"N/A","Released":"10 Oct 2008","Runtime":"129 min","Genre":"Drama, Romance","Director":"Atul Agnihotri","Writer":"Atul Agnihotri (screenplay), Chetan Bhagat (additional dialogue), Chetan Bhagat (book), Jalees Sherwani (lyrics), Alok Upadhyay (additional dialogue)","Actors":"Sharman Joshi, Amrita Arora, Sohail Khan, Isha Koppikar","Plot":"Hello... is a tale about the events that happen one night at a call center. Told through the views of the protagonist, Shyam, it is a story of almost lost love, thwarted ambitions, absence ...","Language":"Hindi","Country":"India","Awards":"N/A","Poster":"https://m.media-amazon.com/images/M/MV5BZGM5NjliODgtODVlOS00OWZmLWIzYzMtMTI2OWIzMTM1ZGRhXkEyXkFqcGdeQXVyNDUzOTQ5MjY@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"3.4/10"}],"Metascore":"N/A","imdbRating":"3.4","imdbVotes":"1,749","imdbID":"tt1087856","Type":"movie","DVD":"08 Dec 2008","BoxOffice":"N/A","Production":"N/A","Website":"N/A","Response":"True"}`
)

func main() {
	if err := http.ListenAndServe(":18443", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("s")
		if s != "" {
			w.Write([]byte(searchRes))
		} else {
			w.Write([]byte(titleRes))
		}
	})); err != nil {
		panic(err)
	}
}
