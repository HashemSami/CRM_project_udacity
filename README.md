## Overview:

The project represents the backend of a customer relationship management (CRM) web application. As users interact with the app via some user interface, the server will support all of the functionalities:

- Getting a list of all customers
- Getting data for a single customer
- Adding a customer
- Updating a customer's information
- Removing a customer

## Motivation:

Many applications you use daily include the functionality to read, write, update, and delete content provided by APIs.

This project allows users to make API requests to a backend server to interact with customer data. After completing this project, you'll solidify your skills in designing and creating a server (and effectively, a basic API) to support dynamic CRUD applications in the real world.

## Getting Started:

- #### Install GO:
  - Download the latest version of Go

> **_Note:_** you need to have newer versions of Go (1.18 and above) to be able to run the project without issues, as this project is leveraging Go workspaces feature which enables you to run the project with multiple modules.

- #### Clone Github repo

  - Clone the project repository

- #### Run the Project:
  - Open your preferred terminal and navigate to the root directory of the project.
  - Run the project server by executing the below command:
  ```bash
  go run github.com/hashemsami/crm_project_udacity
  ```

## API:

#### Get Customers:

`/customers`

Method: `GET`

Description: Get all Customers from database.

example API return:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```

#### Get Customer:

`/customers/{id}`

Method: `GET`

Description: Get a Customer that has the path id.

example API return:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```

#### Create Customer:

`/customers`

Method: `POST`

Description: Create new Customer in the database.

example POST Body:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```

example API return:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```

#### Update Customer:

`/customers/{id}`

Method: `PATCH`

Description: Update data for Customers in the database.

example PATCH Body:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```

example API return:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```

#### Delete Customer:

`/customers/{id}`

Method: `DELETE`

Description: Delete a Customer from database.

example API return:

```bash
{
    "id": "1",
    "name": "Hashem",
    "role": "employee",
    "email": "test@test.com",
    "phone": 12216695,
    "contacted": true
}
```
