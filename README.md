# AB-Hack
Back-End Service for Hackathon Events

### Compiling Requirements
* Docker with support for Linux containers
* Docker compose
* GNU Make
* Internet connection

### Development Requirements
* See _Compiling Requirements_
* Go IDE (we like GoLand)
* Go compiler (latest)
* Docker Compose

### Dependencies
* PostgreSQL Database

### Building
* make build - Create project binaries and production docker image
* make run - Run service locally for deployment. Please make sure that you perform _make build_ in order to install dependencies and create project binaries

#### Notes

To be able to run the service normally you need to setup the database migrations on the first run.
You can find the database migration SQL files in `migrations` folder.

### Setting up the Environment Variables
Refer to the table below for environment variables used in AB-Hack :

  | Configuration | Example Value | Description | Required |
      | ------------- | ------------- | ----------- | ----------- |
  | <SERVER_BASE_PATH> | /api | AB-Hack base URL | false |
  | <SERVER_PORT> | 8080 | AB-Hack exposed port | false |
  | <DB_USERNAME> | userA | Postgres Username | true |
  | <DB_PASSWORD> | mySuperSecurePassword | Postgres Password | true |
  | <DB_HOST> | my-postgres.database | Postgres Host | true |
  | <DB_PORT> | 5432 | Postgres Port | false |
  | <DB_NAME> | myDB | Postgres Database Name | true |
  | <DB_SSL_ENABLED> | false | Postgres SSL Connection | false |
