# RSS Posts Agregator API
### Introduction
This API allows to create RSS feeds. Users can follow different feeds. All feeds posts are fetched and updated in the database with some time interval. Users can get the newest posts according to their feeds.
### Technologies Used
* [Go](https://go.dev/)
* [Chi](https://github.com/go-chi/chi/)
* [SQLS](https://github.com/jmoiron/sqlx/)
* [Goose](https://github.com/pressly/goose/)
* PostgreSQL
### Installation Guide
* Clone this repository [here](https://github.com/Taras-Rm/rss-agg)
* The develop branch is - `dev`
* Run npm install to install all dependencies
* Create an .env file localy. See .env.example for assistance.
### Usage
* Run `go build && ./rss-agg` to start the application.
### Migration
* For running migrations, go to sql/schema and execute `goose postgres postgres://{user}:{password}@{host}:{port}/{dbname} up`
### API Endpoints
| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| POST | /v1/users | To create a new user |
| GET | /v1/users | To retrieve all users |
| POST | /v1/feeds | To create a new feed |
| GET | /v1/feeds | To retrieve all feeds |
| POST | /v1/feed_follows | To follow a feed |
| GET | /v1/feed_follows | To retrieve all followed feeds |
| DELETE | /v1/feed_follows/:feed_follow_id | To delete feed follow |
| GET | /v1/posts | To retrieve the newest posts for users |
| GET | /v1/healthz | To check api health |
| GET | /v1/err | To check error respons |
