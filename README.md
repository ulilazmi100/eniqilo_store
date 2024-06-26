# Eniqilo Store

## 🌄 Background

Eniqilo Store is an backend application for a store to serve their customers, basically a highly performant e-commerce backend.
Eniqilo Store is part of Project Sprint Batch 2 Week 2 Project. Projectsprint is a sprint to create many resourceful projects in a short time with rigorous testing, load testing, and showcase. This initiative aims to demonstrate the ability to deliver high-quality, scalable applications quickly and effectively.

---

## 🚀 Getting Started

### Prerequisites

- Go (1.19 or later)
- Postgres
- [K6](https://k6.io/docs/get-started/installation/) for testing
- WSL (if on Windows)

### Running the Project

#### Environment Variables

Set the following environment variables:
```bash
export DB_NAME=your_db_name
export DB_PORT=5432
export DB_HOST=localhost
export DB_USERNAME=your_db_user
export DB_PASSWORD=your_db_password
export DB_PARAMS="sslmode=disabled" # this is needed because in production, we use `sslrootcert=rds-ca-rsa2048-g1.pem` and `sslmode=verify-full` flag to connect
# read more: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.SSL.html
export JWT_SECRET=your_jwt_secret
export BCRYPT_SALT=8 or 10 depending your requirements # don't use 8 in prod! use > 10
```

#### Running Migrations

```bash
migrate -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?$DB_PARAMS" -path db/migrations up
```

#### Running the Server

```bash
go run main.go
```

#### Optional cleanup if already done or something went wrong (Reverse the migrations if you wanted to)

```bash
migrate -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?$DB_PARAMS" -path db/migrations down
```

---

## 🧪 Testing

### Prerequisites
- [K6](https://k6.io/docs/get-started/installation/)
- A linux environment (WSL / MacOS should be fine)

### Environment Variables
- `BASE_URL` fill this with your backend url (eg: `http://localhost:8080`)

### Steps
1. Install [K6](https://k6.io/docs/get-started/installation/) and other prerequisites.
2. Run the server
3. Open the 'tests' folder 
4. Run the tests:
    #### For regular testing
    ```bash
    BASE_URL=http://localhost:8080 make run
    ```
    #### For load testing
    ```bash
    BASE_URL=http://localhost:8080 make runLoadTest
    ```

    #### For timed testing
    ```bash
    BASE_URL=http://localhost:8080 make run_timed
    ```

    #### For run for debugging, with result in the output.txt
    ```bash
    BASE_URL=http://localhost:8080 make run_debug
    ```

---

## 📝 Requirements

[Requirements' Notion](https://openidea-projectsprint.notion.site/EniQilo-Store-93d69f62951c4c8aaf91e6c090127886)

### Non-Functional Requirements

- Backend:
  - Golang with any web framework
  - Postgres database
  - Port: 8080
  - No ORM/Query generator; use raw queries
  - No external caching
  - Environment Variables:
    ```bash
    export DB_NAME=
    export DB_PORT=
    export DB_HOST=
    export DB_USERNAME=
    export DB_PASSWORD=
    export DB_PARAMS="sslmode=disabled" # this is needed because in production, we use `sslrootcert=rds-ca-rsa2048-g1.pem` and `sslmode=verify-full` flag to connect 
    # read more: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.SSL.html
    export JWT_SECRET=
    export BCRYPT_SALT=8 # don't use 8 in prod! use > 10
    ```

  - Use docker images and docker registry
  Read more in the [requirements' notion](https://openidea-projectsprint.notion.site/EniQilo-Store-93d69f62951c4c8aaf91e6c090127886)

### Database Migration

- Use [golang-migrate](https://github.com/golang-migrate/migrate) for managing database migrations:
  - Create migration:
    ```bash
    migrate create -ext sql -dir db/migrations add_user_table
    ```
  - Execute migration:
    ```bash
    migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
    ```
  - Rollback migration:
    ```bash
    migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations down
    ```

    Read more in the [requirements' notion](https://openidea-projectsprint.notion.site/EniQilo-Store-93d69f62951c4c8aaf91e6c090127886)

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