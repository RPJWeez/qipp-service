##Qipp Service

This is a rest service for "twitter clone" called Qipp.  I created this as a way to learn Go.

### Tech Stack

* Go language
* gorilla/mux for http routing
* gorm for database connection/ORM/migration
* postgresql (in cloud, provided by heroku)
* user authentication/OAuth provided by Auth0

docker for container runtime

### Run locally

* create .env file in /env containing the following:
```
export JWKS_ENDPOINT=""
export JWT_ISSUER=""
export JWT_AUD=""
export DB_HOST=""
export DB_PORT=""
export DB_DATABASE=""
export DB_USER=""
export DB_PASS=""
```
* build: `./build.sh` - I developed this on windows to I needed a script to remember how to build for linux so it would run in an alpine docker image
* run docker image: `docker-compose up --build`

### TODO
* in mem database for local dev
* prometheus metrics
* produce kafka/mq messages?


