
# KopiKami API Documentation

## Introduction
KopiKami is a comprehensive inventory and transaction management system designed for coffee shop operations. This API provides endpoints for managing users, products, raw materials, inventory, transactions, reports, and dashboards. Role-based access control (RBAC) ensures secure and streamlined operations.

---

## Table of Contents
1. [Getting Started](#getting-started)
2. [Features](#features)
3. [Roles and Permissions](#roles-and-permissions)
4. [API Endpoints](#api-endpoints)
    - [Authentication](#authentication)
    - [Users Management](#users-management)
    - [Products](#products)
    - [Raw Materials and Batches](#raw-materials-and-batches)
    - [Inventory](#inventory)
    - [Transactions](#transactions)
    - [Reports](#reports)
    - [Dashboard](#dashboard)
5. [Database Schema](#database-schema)
6. [Dummy Data](#dummy-data)
7. [Setup and Deployment](#setup-and-deployment)

---

## Getting Started
### Prerequisites
- **Go**: v1.19 or higher
- **MySQL**: v8.0 or higher
- **Docker**: v20.10 or higher (optional for containerization)

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/budytabootie/kopikami.git
   cd kopikami
   ```

2. Set up your `.env` file with the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASS=password
   DB_NAME=kopikami
   JWT_SECRET=your_jwt_secret
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

5. Access the API at `http://localhost:8080`.

---

## Features
### Authentication
- **JWT-based authentication**
- Secure login and logout

### User Management (Admin Only)
- CRUD operations for users

### Product Management
- Manage coffee products and their details

### Raw Materials and Batches
- Manage raw materials and their inventory batches

### Inventory Management
- FIFO-based inventory updates
- Check current stock levels

### Transactions
- Create and view sales transactions

### Reports
- Generate sales and stock reports

### Dashboard
- View sales stats, inventory stats, and sales trends

---

## Roles and Permissions
### Admin
- Full access to all features

### Cashier
- Restricted to transaction and inventory viewing

---

## API Endpoints
### Authentication
- **POST /api/auth/register**: Register a new user
- **POST /api/auth/login**: Login and get a JWT token
- **POST /api/auth/logout**: Logout

### Users Management
- **POST /api/admin/users**: Create a user
- **GET /api/admin/users**: Get all users
- **GET /api/admin/users/:id**: Get user by ID
- **PUT /api/admin/users/:id**: Update user
- **DELETE /api/admin/users/:id**: Delete user

### Products
- **GET /api/admin/products**: Get all products
- **POST /api/admin/products**: Create a product
- **PUT /api/admin/products/:id**: Update a product
- **DELETE /api/admin/products/:id**: Delete a product

### Raw Materials and Batches
- **POST /api/admin/raw-materials**: Add raw material
- **GET /api/admin/raw-materials**: View raw materials
- **POST /api/admin/raw-material-batches**: Add raw material batch
- **GET /api/admin/raw-material-batches**: View raw material batches

### Inventory
- **GET /api/inventory**: Check stock levels

### Transactions
- **POST /api/cashier/transactions**: Create a transaction
- **GET /api/admin/transactions**: Get all transactions (Admin)
- **GET /api/cashier/transactions**: Get transactions by cashier

### Reports
- **GET /api/reports/sales**: Generate sales report
- **GET /api/reports/stock**: Generate stock report

### Dashboard
- **GET /api/dashboard/sales-stats**: View sales stats
- **GET /api/dashboard/inventory-stats**: View inventory stats
- **GET /api/dashboard/trends**: View sales trends

---

## Database Schema
Refer to the `models` directory for entity structures. Key tables:
- **Users**: Stores user details
- **Products**: Stores product information
- **Raw Materials**: Stores raw material details
- **Transactions**: Logs all sales transactions
- **Inventory Logs**: Tracks inventory changes

The following is the database schema design for the KopiKami application:

![Database Schema](kopikami.png)

---

## Dummy Data
### JSON Format
Refer to the file: `dummy_data.json`
- [Dummy Data (JSON)](dummy-data.json)

### SQL Format
Refer to the file: `dummy_data.sql`
- [Dummy Data (SQL)](dummy-data.sql)

These files contain pre-generated data for testing purposes. Data includes:
- **Products**
- **Raw Materials**
- **Inventory Logs**
- **Transactions** (with associated items)

---

## Setup and Deployment
### Docker Setup
1. Build the Docker image:
   ```bash
   docker build -t kopikami:latest .
   ```

2. Run the container:
   ```bash
   docker-compose up
   ```

### Testing
- Use Postman or cURL to test the API.
- A `tests.http` file is included for quick testing.

### Future Improvements
- Add role-specific dashboards
- Implement email notifications for low stock
- Enhance reporting with graphs and CSV export

---

For more information, contact satyabudinugroho@gmail.com or open an issue in the repository.
