# Running the Application

Hello, thank you for taking the time to review my coding challenge homework.  to run the project you need to have docker, docker-compose and golang installed on your system.  First you will need to initialize the database, which is a dockerized postgres DB.  Once you clone this repository, navigate to the base directory and run the command:

```docker-compose up```

This will utilize the docker-compose yml and the sql-init.sql to set up a test database with some users and some rentals.

Next you will want to run the unit tests, which can be achieved by running `go test ./...` from the root directory of the project.

Finally, you can either: 

Run the application using the command `go run cmd/web/main.go` 
OR 
Build the application using go `build cmd/web/main.go` then run it using the executable called `main` that will be built.

Either way, the application will now be running and listening to port 8080.  You can navigate to the rentals endpoint described below at `http://localhost:8080/rentals` in any browser or using postman.

Happy Testing!

# Coding Challenge

Thanks for applying for a backend role at Outdoorsy. We've put together this code challenge, which should take around 3-4 hours to complete.

## Functionality
The task is to develop a rentals JSON API that returns a list of rentals that can be filtered, sorted, and paginated. We have included files to create a database of rentals.

Your application should support the following endpoints.

- `/rentals/<RENTAL_ID>` Read one rental endpoint
- `/rentals` Read many (list) rentals endpoint
    - Supported query parameters
        - price_min (number)
        - price_max (number)
        - limit (number)
        - offset (number)
        - ids (comma separated list of rental ids)
        - near (comma separated pair [lat,lng])
        - sort (string)
    - Examples:
        - `rentals?price_min=9000&price_max=75000`
        - `rentals?limit=3&offset=6`
        - `rentals?ids=3,4,5`
        - `rentals?near=33.64,-117.93` // within 100 miles
        - `rentals?sort=price`
        - `rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`

The rental object JSON in the response should have the following structure:
```json
{
  "id": "int",
  "name": "string",
  "description": "string",
  "type": "string",
  "make": "string",
  "model": "string",
  "year": "int",
  "length": "decimal",
  "sleeps": "int",
  "primary_image_url": "string",
  "price": {
    "day": "int"
  },
  "location": {
    "city": "string",
    "state": "string",
    "zip": "string",
    "country": "string",
    "lat": "decimal",
    "lng": "decimal"
  },
  "user": {
    "id": "int",
    "first_name": "string",
    "last_name": "string"
  }
}
```

## Notes
- Running `docker-compose up` will automatically generate a postgres database and some data to work with. Connect and use this database.
- Write production ready code.
- Please make frequent, and descriptive git commits.
- Use third-party libraries or not; your choice.
- Please use Golang to complete this task.
- Feel free to add functionality as you have time, but the feature described above is the priority.
- Please add tests

## What we're looking for
- The functionality of the project matches the description above
- An ability to think through all potential states
- In the README of the project, describe exactly how to run the application and execute the tests

When complete, please push your code to Github to your own account and send the link to the project or zip the project (including the `.git` directory) and send it back.

Thank you and please ask if you have any questions!
