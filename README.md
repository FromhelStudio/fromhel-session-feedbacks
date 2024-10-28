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

   Edit the [.env](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2F.env%22%2C%22path%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2F.env%22%2C%22scheme%22%3A%22file%22%7D%7D) file to include your MongoDB URI:

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
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

### API Endpoints

#### Create Rating

- **URL:** `/rating`
- **Method:** [`POST`](command:_github.copilot.openSymbolFromReferences?%5B%22%22%2C%5B%7B%22uri%22%3A%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Frouter%2Frouter.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22pos%22%3A%7B%22line%22%3A24%2C%22character%22%3A8%7D%7D%5D%2C%2255c069d4-181a-4600-b0ca-450ebaa5cdf4%22%5D "Go to definition")
- **Request Body:**

  ```json
  {
    "rating": 1,
    "feedback": "Your feedback here"
  }
  ```

- **Response:**

  ```json
  {
    "message": "Rating created",
    "statusCode": 201
  }
  ```

### Project Details

- **Rating Model:** Defined in [rating.go](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Fservices%2Frating.go%22%2C%22path%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Fservices%2Frating.go%22%2C%22scheme%22%3A%22file%22%7D%7D)
- **Rating Service:** Handles the creation of ratings in MongoDB, defined in [rating.go](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Fservices%2Frating.go%22%2C%22path%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Fservices%2Frating.go%22%2C%22scheme%22%3A%22file%22%7D%7D)
- **Create Rating Handler:** Handles the HTTP request for creating a rating, defined in [handlers.go](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Fweb%2Fhandlers.go%22%2C%22path%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Fweb%2Fhandlers.go%22%2C%22scheme%22%3A%22file%22%7D%7D)
- **Router Initialization:** Sets up the routes and loads environment variables, defined in [router.go](http://_vscodecontentref_/#%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Frouter%2Frouter.go%22%2C%22path%22%3A%22%2FUsers%2Fgabrielaraujolima%2FFromhel%2Ffromhel-session-feedbacks%2Finternal%2Frouter%2Frouter.go%22%2C%22scheme%22%3A%22file%22%7D%7D)
