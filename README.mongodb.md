# mongoDB

## container setup


1. **Start MongoDB with Docker**  
   Use this command to set up the MongoDB container, setting the same username, password, and database (`fiber`), as expected by your application:

   ```bash
   docker run --name mongodb -p 27017:27017 -d \
     -e MONGO_INITDB_ROOT_USERNAME=username \
     -e MONGO_INITDB_ROOT_PASSWORD=password \
     -e MONGO_INITDB_DATABASE=fiber \
     mongo
   ```

   This command:
   - Maps MongoDB’s default port `27017` to your localhost.
   - Initializes MongoDB with the username and password set in your `main.go` file’s MongoDB URI.
   - Sets up a database called `fiber`, as expected in your URI.

2. **Adjust MongoDB Connection Settings (if needed)**  
   In your code, ensure the `databaseConnection()` function's URI matches the above settings:
   
   ```go
   client, err := mongo.Connect(ctx, options.Client().ApplyURI(
       "mongodb://username:password@localhost:27017/fiber").SetServerSelectionTimeout(5*time.Second))
   ```

3. **Running the Application**  
   With the MongoDB container running, start your Go application using:
   
   ```bash
   go run main.go
   ```

4. **Verify the Connection**  
   If everything is set up correctly, you should see the "Database connection success!" message, indicating a successful connection.

## environment variables


1. **Install the `godotenv` Package**  
   First, add the `godotenv` package to your project:
   ```bash
   go get github.com/joho/godotenv
   ```

2. **Create a `.env` File**  
   In the root of your project, create a `.env` file with the following content:
   ```plaintext
   DB_URI=mongodb://localhost:27017
   DB_NAME=example_db
   ```

3. **Update `main.go` to Load Environment Variables**  
   Import `godotenv` and load the variables at the start of your `main()` function. Update the `databaseConnection()` function to use these environment variables.

   Here’s how to update your `main.go` file:

   ```go
   package main

   import (
       "context"
       "fmt"
       "log"
       "os"
       "time"

       "github.com/gofiber/fiber/v2"
       "github.com/gofiber/fiber/v2/middleware/cors"
       "github.com/joho/godotenv"
       "go.mongodb.org/mongo-driver/mongo"
       "go.mongodb.org/mongo-driver/mongo/options"
       "gfgoth/api/routes"
       "gfgoth/pkg/book"
   )

   func main() {
       // Load environment variables from .env file
       err := godotenv.Load()
       if err != nil {
           log.Fatal("Error loading .env file")
       }

       db, cancel, err := databaseConnection()
       if err != nil {
           log.Fatalf("Database Connection Error: %s", err)
       }
       fmt.Println("Database connection success!")

       bookCollection := db.Collection("books")
       bookRepo := book.NewRepo(bookCollection)
       bookService := book.NewService(bookRepo)

       app := fiber.New()
       app.Use(cors.New())
       app.Get("/", func(ctx *fiber.Ctx) error {
           return ctx.SendString("Welcome to the gfgoth mongo book shop!")
       })
       api := app.Group("/api")
       routes.BookRouter(api, bookService)

       defer cancel()
       log.Fatal(app.Listen(":8080"))
   }

   func databaseConnection() (*mongo.Database, context.CancelFunc, error) {
       ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
       
       dbURI := os.Getenv("DB_URI")
       dbName := os.Getenv("DB_NAME")
       
       client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI).SetServerSelectionTimeout(5*time.Second))
       if err != nil {
           cancel()
           return nil, nil, err
       }
       db := client.Database(dbName)
       return db, cancel, nil
   }
   ```

4. **Running the Application**  
   With the `.env` file created and `godotenv` installed, simply run your application as usual:

   ```bash
   go run main.go
   ```

This setup will load the `DB_URI` and `DB_NAME` values from the `.env` file, allowing you to configure the MongoDB connection without hardcoding values in your code. This makes it easy to switch environments or modify settings.


I think we are getting closer.

Can you explain a few things in order to complete the final piece of documentation.

When I start the local container:

   ```bash
   docker run --name mongodb -p 27017:27017 -d \
     -e MONGO_INITDB_ROOT_USERNAME=username \
     -e MONGO_INITDB_ROOT_PASSWORD=password \
     -e MONGO_INITDB_DATABASE=fiber \
     mongo
   ```


and then set the .env file:

```
DB_URI=mongodb://localhost:27017
DB_NAME=books
```

And the main.go has:

`"mongodb://username:password@localhost:27017/fiber")`


Can you explain the relationship between the `MONGO_INITDB_DATABASE=fiber`, the `DB_NAME=books` and the `localhost:27017/fiber` code?