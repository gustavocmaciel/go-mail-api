# Go Mail API

An extremely small REST API  written in Go that *simulates* a simple mail backend.

Built with [Gorilla Mux](https://github.com/gorilla/mux), a powerful HTTP request router and dispatcher for Go.


## Getting Started

To get a local copy up and running, you can clone this repository using the following command:

```bash
git clone https://github.com/gustavocmaciel/mail-api.git
```

 Once you have cloned the repository, you can run the project using the go run command:

 ```bash
 go run cmd/main.go
 ```

### API Endpoints

- `GET /users`: Get all users.
- `GET /mails`: Get all mails.
- `GET /mailbox/{user}/{mailbox}`: Get all mails in the specified mailbox for the given user.
- `POST /user`: Add a new user to the system.

## Usage

Fell free explore the code and use it for educational purposes. You can also experiment and build upon it as needed.

### Note

This project is still a work in progress (WIP).

## License

This project is licensed under the [MIT License](LICENSE).
