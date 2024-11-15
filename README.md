# Fromhel Session Feedbacks

This project is a web service for collecting session feedback ratings. It is built using Go, Gin, and MongoDB.

## Getting Started

### Prerequisites

- Go 1.16+
- MongoDB
- [Gin](https://github.com/gin-gonic/gin)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/FromhelStudio/fromhel-session-feedbacks.git
   cd fromhel-session-feedbacks
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Copy the example environment file and set your MongoDB URI:

   ```sh
   cp .env.example .env
   ```

   Edit the [.env](http://_vscodecontentref_/16) file to include your MongoDB URI:

   ```
   MONGODB_URI=mongodb://your_mongodb_uri
   ```

### Running the Server

1. Navigate to the `cmd/server` directory:

   ```sh
   cd cmd/server
   ```

2. Run the server:

   ```sh
   go run [main.go](http://_vscodecontentref_/17)
   ```

   The server will start on `http://localhost:8080`.

### API Endpoints

#### Create Rating

- **URL:** `/rating`
- **Method:** `POST`
- **Request Body:**

  ```json
  {
    "rating": 1,
    "feedback": "Your feedback here",
    "gameName": "bulletspeel" // bulletspeel or cordel only
  }
  ```

- **Response:**

  ```json
  {
    "message": "Rating created",
    "statusCode": 201
  }
  ```

#### Get Ratings

- **URL:** `/rating/:page`
- **Method:** `GET`
- **Query Parameters:**

  - `game` (required): The name of the game (either "bulletspeel" or "cordel")

- **Response:**

  ```json
  {
    "ratings": [
      {
        "id": "string",
        "game": "string",
        "rating": 1,
        "feedback": "string",
        "createdAt": "string"
      }
    ],
    "statusCode": 200
  }
  ```

#### Create Session

- **URL:** `/session`
- **Method:** `POST`
- **Request Body:**

  ```json
  {
    "gameName": "bulletspeel", // bulletspeel or cordel only
    "timespent": 0,
    "deaths": 0,
    "colorPicked": 0,
    "enemysKilled": 0,
    "gameFinished": true,
    "money": 0,
    "ammunation": 0,
    "items": true
  }
  ```

- **Response:**

  ```json
  {
    "message": "Created.",
    "statusCode": 201
  }
  ```

### Get Sessions

- **URL** `/session/:page`
- **Method** `GET`
- **Query Parameters:**

  - `game` (required): The name of the game (either "bulletspeel" or "cordel")

- **Response**
  ```json
  {
    "sessions": [
      {
        "Id": "UUID/string",
        "Game": "string",
        "Timespent": 1,
        "Deaths": 1,
        "ColorPicked": 1,
        "EnemysKilled": 1,
        "GameFinished": true,
        "Money": 1,
        "Ammunation": 1,
        "Items": true,
        "CreatedAt": "Date/string"
      }
    ],
    "statusCode": 200
  }
  ```

### Project Details

- **Rating Model:** Defined in [rating.go](http://_vscodecontentref_/18)
- **Rating Service:** Handles the creation of ratings in MongoDB, defined in [rating.go](http://_vscodecontentref_/19)
- **Create Rating Handler:** Handles the HTTP request for creating a rating, defined in [rating.go](http://_vscodecontentref_/20)
- **Get Ratings Handler:** Handles the HTTP request for retrieving ratings, defined in [rating.go](http://_vscodecontentref_/21)
- **Session Model:** Defined in [sessions.go](http://_vscodecontentref_/22)
- **Session Service:** Handles the creation of sessions in MongoDB, defined in [sessions.go](http://_vscodecontentref_/23)
- **Create Session Handler:** Handles the HTTP request for creating a session, defined in [sessions.go](http://_vscodecontentref_/24)
- **Router Initialization:** Sets up the routes and loads environment variables, defined in [router.go](http://_vscodecontentref_/25)
