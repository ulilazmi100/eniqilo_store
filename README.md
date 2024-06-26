# Eniqilo Store

## 🌄 Background

Eniqilo Store is a highly performant e-commerce backend application designed to serve store customers efficiently. Developed as part of Project Sprint Batch 2 Week 2, this project aims to showcase the ability to deliver high-quality, scalable applications quickly with rigorous testing and load testing.

---

## 🚀 Getting Started

### Prerequisites

- Go (1.19 or later)
- PostgreSQL
- [K6](https://k6.io/docs/get-started/installation/) for load testing
- WSL (if on Windows)

### Installation and Setup

1. **Clone the repository:**

    ```bash
    git clone https://github.com/ulilazmi100/eniqilo_store.git
    cd eniqilo_store
    ```

2. **Set up environment variables:**

    Create a `.env` file or export these variables in your shell:

    ```bash
    export DB_NAME=your_db_name
    export DB_PORT=5432
    export DB_HOST=localhost
    export DB_USERNAME=your_db_user
    export DB_PASSWORD=your_db_password
    export DB_PARAMS="sslmode=disable"  # or "sslrootcert=rds-ca-rsa2048-g1.pem&sslmode=verify-full" for production
    export JWT_SECRET=your_jwt_secret
    export BCRYPT_SALT=8 or 10  # or same or higher than 10 for production
    ```

3. **Run database migrations:**

    ```bash
    migrate -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?$DB_PARAMS" -path db/migrations up
    ```

4. **Start the server:**

    ```bash
    go run main.go
    ```

5. **Optional: Rollback migrations if needed:**

    ```bash
    migrate -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?$DB_PARAMS" -path db/migrations down
    ```

### Running with Docker

1. **Build Docker image:**

    ```bash
    docker build -t eniqilo_store .
    ```

2. **Run Docker container:**

    ```bash
    docker run -p 8080:8080 --env-file .env eniqilo_store
    ```

---

## 🧪 Testing

### Prerequisites

- [K6](https://k6.io/docs/get-started/installation/)
- Linux environment (WSL/MacOS)

### Running Tests

1. **Set environment variable:**

    ```bash
    export BASE_URL=http://localhost:8080
    ```

2. **Run the server.**

3. **Navigate to the `tests` folder and run the tests:**

    #### Regular testing

    ```bash
    BASE_URL=http://localhost:8080 make run
    ```

    #### Load testing

    ```bash
    BASE_URL=http://localhost:8080 make runLoadTest
    ```

    #### Timed testing

    ```bash
    BASE_URL=http://localhost:8080 make run_timed
    ```

    #### Debug testing (output to `output.txt`)

    ```bash
    BASE_URL=http://localhost:8080 make run_debug
    ```

---

## 📝 Requirements

### Functional Requirements

**Authentication & Authorization:**
- Register staff: `POST /v1/staff/register`
- Login staff: `POST /v1/staff/login`

**Product Management:**
- Add product: `POST /v1/product`
- Get products: `GET /v1/product`
- Edit product: `PUT /v1/product/{id}`
- Delete product: `DELETE /v1/product/{id}`

**SKU Search:**
- Search products: `GET /v1/product/customer`

**Checkout:**
- Register customer: `POST /v1/customer/register`
- Search customer: `GET /v1/customer`
- Checkout products: `POST /v1/product/checkout`
- Checkout history: `GET /v1/product/checkout/history`

### Non-Functional Requirements

- Develop on Linux (`WSL` is acceptable for Windows users).
- Backend server: Golang with any web framework.
- Database: PostgreSQL.
- Port: 8080.
- No ORM/Query generator; use raw queries.
- No external caching.
- Use Docker images and Docker registry.

### Further Readings:
- [Eniqilo Store's Requirements' Notion Page](https://openidea-projectsprint.notion.site/EniQilo-Store-93d69f62951c4c8aaf91e6c090127886)

---

## Database Migration

Use [golang-migrate](https://github.com/golang-migrate/migrate) for managing database migrations:

- Create migration:

    ```bash
    migrate create -ext sql -dir db/migrations add_user_table
    ```

- Execute migration:

    ```bash
    migrate -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?$DB_PARAMS" -path db/migrations up
    ```

- Rollback migration:

    ```bash
    migrate -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?$DB_PARAMS" -path db/migrations down
    ```

---

## 👥 Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/AmazingFeature`).
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the branch (`git push origin feature/AmazingFeature`).
5. Open a pull request.

---

## 📝 License

-

---

## 📚 Resources

- **Notion:** [Eniqilo Store's Requirements' Notion Page](https://openidea-projectsprint.notion.site/EniQilo-Store-93d69f62951c4c8aaf91e6c090127886)
- **Tests:** [Project Sprint Batch 2 Week 2 Test Cases](https://github.com/nandanugg/EniQiloStoreTestCasesPSW2B2?tab=readme-ov-file#for-load-testing)
- **Migrations:** [Golang Migration](https://github.com/golang-migrate/migrate)

---

## 📞 Contact

[Muhammad Ulil 'Azmi](https://github.com/ulilazmi100) - [@M_Ulil_Azmi](https://twitter.com/M_Ulil_Azmi)

Project Link: [https://github.com/ulilazmi100/eniqilo_store](https://github.com/ulilazmi100/eniqilo_store)

---