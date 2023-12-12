# Go Mail API

A small REST API  written in Go that *simulates* a simple mail backend.

Built with [Gorilla Mux](https://github.com/gorilla/mux), a powerful HTTP request router and dispatcher for Go.


## Getting Started

To get a local copy up and running, you can clone this repository using the following command:

```bash
git clone https://github.com/gustavocmaciel/go-mail-api.git
```

To simplify the setup process, we recommend using Docker for containerization. Follow the steps below to set up your development environment using Docker.

## Prerequisites

Make sure you have the following prerequisites installed on your system:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Steps
1. Run the following command to start the Docker containers:

   ```bash
   docker compose up -d
   ```

2. The application should now be accessible at http://localhost:8080.



## API Endpoints

### User Endpoints

#### Create User

- **Description:** Handles the creation of a new user.
- **Method:** `POST`
- **Endpoint:** `/create_user`

#### Get Users

- **Description:** Handles the retrieval of all users.
- **Method:** `GET`
- **Endpoint:** `/get_users`

### Mail Endpoints

#### Create Mail

- **Description:** Handles the creation of a new mail.
- **Method:** `POST`
- **Endpoint:** `/create_mail`

#### Get Mail

- **Description:** Handles the retrieval of a specific mail by ID.
- **Method:** `GET`
- **Endpoint:** `/get_mail/{mail_id}`

#### Mailbox

- **Description:** Handles the retrieval of emails in a specific mailbox for a user.
- **Method:** `GET`
- **Endpoint:** `/mailbox/{user}/{mailbox}`

## Usage

Fell free explore the code and use it for educational purposes. You can also experiment and build upon it as needed.

## License

This project is licensed under the [MIT License](LICENSE).
