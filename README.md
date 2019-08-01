# movie-listing-app

Movie listing app using the OMDB( http://www.omdbapi.com/) API.

## How to run

### Using source code

If **Go** isn't installed in you pc, then install is first. **Go** version +1.10 required. You can find the guide in here: https://golang.org/doc/install.

Clone the github repository:

```
# clone the github repository
$ git clone git@github.com:nightfury1204/movie-listing-app.git
```

Move into the repository folder:

```
$ cd movie-listing-app/ 
```

Use the following command to run the app. Change the `YOUR_API_TOKEN` to your OMDB API token. This will run the server in port `8443`.

```
$ go run *.go run \
     --omdb-api-token=YOUR_API_TOKEN
```

Browse `http://localhost:8443` to use the app.

###  Using Docker

Run the Docker container using following command. Change the `YOUR_API_TOKEN` to your OMDB API token. 

```
$  docker run -it --name movie-listing-app -p 8443:8443 -e OMDB_API_TOKEN=YOUR_API_TOKEN nightfury1204/movie-listing-app:latest
```

Browse `http://localhost:8443` to use the app.


## Username and Password

Currently following username and password are hard coded into the system:

| Username | Password |
|----------|----------|
| user1    | pass1    |
| user2    | pass2    |
| user3    | pass3    |

