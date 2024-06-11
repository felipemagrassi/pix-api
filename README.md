# PIX Api 

Simple API to recreate a Pix receiver backend

## Running the project

Run the following command to start the project:

```bash
make up
```

The project will be available at `http://localhost:8080`

Having 5 endpoints:
    - POST /receiver
    - GET /receiver/{id}
    - GET /receiver/
    - PUT /receiver/{id}
    - DELETe /receiver/{id}

## Seeding the database

Run the following command to seed the database with sample accounts

```bash
make seed
```

## Running the tests

Run the following command to run the tests:

```bash
make test
```

## Documentation

If you are running the project locally, you can access the API documentation at `http://localhost:8080/docs/index.html`

## Roadmap

- [] Create bank tables to validate correct banks that support pix operation
- [] Same receiver having multiple Pix Keys
