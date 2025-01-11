# Inventory Management and Sales API

## **Overview**
This is an API for managing an inventory and sales system for a coffee shop. The application provides features such as managing products, transactions, and inventory. It implements secure and scalable backend best practices using the following technologies:

- **Language**: Go (Golang)
- **Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL
- **Authentication**: JWT
- **Containerization**: Docker

---

## **Features**
1. **User Authentication**:
   - User registration and login with JWT-based authentication.
   - Role-based access control.

2. **Product Management**:
   - Create, read, update, and delete products.
   - Unique validation for product names.

3. **Transaction Management**:
   - Record sales transactions.
   - Automatically deduct stock from the inventory.

4. **Inventory Management**:
   - Batch-based inventory tracking.
   - FIFO (First In, First Out) stock consumption.

5. **API Documentation**:
   - Interactive API documentation using Swagger.

---

## **Setup and Installation**

### Prerequisites
- Docker & Docker Compose
- Go 1.20+
- MySQL (or use Docker)

### Installation Steps

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-repo/inventory-app.git
   cd inventory-app
   ```

2. **Environment Configuration**:
   Copy the example environment file and update it as needed:
   ```bash
   cp .env.example .env
   ```

3. **Run the Application with Docker**:
   ```bash
   docker-compose up --build
   ```

4. **Run the Application Locally**:
   - Install dependencies:
     ```bash
     go mod tidy
     ```
   - Run the application:
     ```bash
     go run main.go
     ```

5. **Access API Documentation**:
   Open [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser.

---

## **Folder Structure**
```plaintext
.
├── controllers    # API controllers (handle HTTP requests)
├── docs           # Auto-generated Swagger documentation
├── middlewares    # Custom middlewares (e.g., authentication)
├── models         # Database models
├── repositories   # Data access layer
├── services       # Business logic layer
├── tests          # Unit and integration tests
├── main.go        # Application entry point
├── go.mod         # Dependencies
└── docker-compose.yml  # Docker configuration
```

---

## **API Endpoints**
### **Authentication**
- `POST /api/v1/auth/register`: Register a new user.
- `POST /api/v1/auth/login`: Authenticate and retrieve a JWT.

### **Products**
- `GET /api/v1/products`: Get all products.
- `POST /api/v1/products`: Add a new product.
- `PUT /api/v1/products/:id`: Update an existing product.
- `DELETE /api/v1/products/:id`: Delete a product.

### **Transactions**
- `POST /api/v1/transactions`: Record a new transaction.
- `GET /api/v1/transactions`: Get all transactions.

### **Inventory**
- `POST /api/v1/inventory`: Add new stock to inventory.
- `GET /api/v1/inventory/:productID`: Get inventory details for a specific product.

---

## **Testing**

### **Run Unit and Integration Tests**
```bash
go test ./...
```

### **Sample Test Commands**
- Ensure products can be added, retrieved, updated, and deleted.
- Verify transactions reduce inventory stock appropriately.

---

## **Contributing**
1. Fork the repository.
2. Create a new branch for your feature/fix.
3. Commit your changes and open a pull request.

---

## **License**
This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## **Contact**
For support or inquiries, please email [your-email@example.com](mailto:your-email@example.com).
