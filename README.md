# movies
RESTful API for movies(something similar to IMDB)

1- Used MySQL  
2- Standard Golang packages for implementing the REST APIs.

Create Table Commands-  
Movies   
```
CREATE TABLE `movies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(90) NOT NULL,
  `popularity_score` decimal(5,2) NOT NULL,
  `director` varchar(45) NOT NULL,
  `imdb_score` decimal(5,2) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)
```   
Genre   
```
CREATE TABLE `genre` (
  `movie_id` int(11) NOT NULL,
  `name` varchar(90) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`movie_id`)
)
```   
Users   
```
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `is_admin` tinyint(4) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
)
```
User Auth   
```
CREATE TABLE `user_auth` (
  `user_id` int(11) NOT NULL,
  `auth_token` varchar(90) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
)
```

**Insert Queries**   
This will create an admin user, and the auth token. This auth token we need to pass as headers in every api call(Admin Group).
```
INSERT INTO `movies`.`users` (`username`, `email`, `is_admin`) VALUES ('abhinav', 'abhinav.maurya001@gmail.com', '1');
INSERT INTO `movies`.`user_auth` (`user_id`, `auth_token`) VALUES ('1', '4e17f8af577919d5c50841543b559a28');
```

This App consists 5 REST APIs   
1- PopulateDB with IMDB JSON file   
**Request-**
```
curl --location --request POST 'localhost:8080/api/movies/admin/populateDB' \
--header 'auth_token: 4e17f8af577919d5c50841543b559a28' \
--header 'Content-Type: application/json' \
--data-raw '{
    "list": [
        {
            "99popularity": 83,
            "director": "Victor Fleming",
            "genre": [
                "Adventure",
                " Family",
                " Fantasy",
                " Musical"
            ],
            "imdb_score": 8.3,
            "name": "The Wizard of Oz"
        },
        {
            "99popularity": 88,
            "director": "George Lucas",
            "genre": [
                "Action",
                " Adventure",
                " Fantasy",
                " Sci-Fi"
            ],
            "imdb_score": 8.8,
            "name": "Star Wars"
        }
        ...
    ]
}
```
2- Add Movie  
**Request**
```
curl --location --request POST 'localhost:8080/api/movies/admin/add' \
--header 'auth_token: 4e17f8af577919d5c50841543b559a28' \
--header 'Content-Type: application/json' \
--data-raw '{
    "movie": {
        "99popularity": 83,
        "director": "Abhinav Victor Fleming",
        "genre": [
            "Adventure",
            " Family",
            " Fantasy",
            " Musical"
        ],
        "imdb_score": 9.3,
        "name": "The Abhinav Wizard of Ozq"
    }
}'
```
3- Edit Movie  
**Request**
```
curl --location --request POST 'localhost:8080/api/movies/admin/edit' \
--header 'auth_token: 4e17f8af577919d5c50841543b559a28' \
--header 'Content-Type: application/json' \
--data-raw '{
    "movie": {
        "99popularity": 83,
        "director": "Abhinav Victor Fleming",
        "genre": [
            "Adventure",
            " Family",
            " Fantasy",
            " Musical"
        ],
        "imdb_score": 9.3,
        "name": "The Abhinav Wizard of Ozq"
    }
}'
```
4- Remove Movie  
**Request**
```
curl --location --request GET 'localhost:8080/api/movies/admin/remove/1' \
--header 'auth_token: 4e17f8af577919d5c50841543b559a28'
```
5- List of Movies  
**Request**
```
curl --location --request GET 'localhost:8080/api/movies/user/view'
```

Run Through Docker-   
```
docker-compose up --build
```