

This project is Bank Account Management System is a backend application developed using Go and PostgreSQL. The project allows users to perform basic banking operations such as:

- Account creation: Users can create a new account by providing their name and an initial deposit amount.
- Deposit: Users can deposit money into their bank accounts.
- Withdraw: Users can withdraw funds from their bank accounts, ensuring they have sufficient balance.
- Transaction management: This includes transferring funds between accounts.


This project simulates a real-world banking system that helps users manage their financial transactions in a secure environment. It provides essential operations such as creating and managing bank accounts, depositing and withdrawing funds, and transferring money between accounts.

The project is particularly useful for developers who want to understand backend development practices, including:

- Using Go for building efficient, scalable systems.
- Interacting with PostgreSQL for data persistence.
- Implementing essential banking logic such as balance management, withdrawals, and transfers.
- integration testing of backend services.
- Creating Swagger documentation for APIs.



To get started with the Bank Account Management System, follow these steps:

	- make build	Compiles the Go source code and creates an executable in the build/ folder.
    - make run	    Builds and runs the application locally.
    - make test	    Runs all unit and integration tests in the project.
    - make fmt	    Formats all Go source files using go fmt.
    - make clean	Cleans the build artifacts and resets the build directory.
    - make swagger	Generates Swagger documentation (swagger.json) from annotated Go code.
