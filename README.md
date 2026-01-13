### RDBMS Implementation 

#### Overview

This is  simple relational database management system (RDBMS) implemented using Golang .The system supports declaring tables with a few column data types, CRUD operations, basic indexing, primary and unique keying, and table joins. The interface is SQL-like and includes an interactive REPL (Read–Eval–Print Loop) mode for direct database interaction.The syatem also has a trivial web application that demonstrates CRUD operations using the custom RDBMS.
This project is a from-scratch implementation of a simple Relational Database Management System (RDBMS) built as part of the Pesapal Junior Developer Challenge.
The goal of the challenge is not to rely on credentials or prior experience, but to demonstrate systems thinking, problem-solving ability, and depth of understanding by building a working database engine and proving its usability through real interaction.

### Key Features

#### 1. Database Management

* Create multiple databases usinf `CREATE` command.
* Switch between databases using a `USE` command.
* Databases are persisted as directories on disk

#### 2. Table Management

* Create tables with predefined schemas
* Schemas define:
  * Column names
  * Column data types
  * Primary keys
  * Unique constraints (extensible)


#### 3. CRUD Operations

The RDBMS supports full CRUD(CREATE, READ (SELECT), UPDATE , DELETE) functionality:

#### 4. Joins

The engine supports basic relational joins using a nested-loop join strategy.

#### 5. SQL-like Interface with Interactive REPL

An interactive REPL (Read–Eval–Print Loop) is provided to interact with the database using SQL-like commands.

Features of the REPL:

* Persistent session state
* Live execution of database commands
* Immediate feedback

Run the REPL:

```bash
go run cmd/repl/main.go
```

Example REPL session:

![REPL to demonstrate functionality](https://github.com/jimmymuthoni/pesapal_rdbms_challenge/blob/aaac16bc07f791d25e57d77c62c6bb3188f5e3cf/rdbms_interface.png)


#### 6. Trivial Web Application (CRUD Demonstration)

To demonstrate real-world usage beyond the CLI, I built a trivial web application on top of the RDBMS.

The web app:

* Exposes HTTP endpoints for CRUD operations
* Translates HTTP requests into database operations
* Uses the same database engine as the REPL

#### Example Endpoints

| Method | Endpoint    | Operation |
| ------ | ----------- | --------- |
| POST   | /users      | Create    |
| GET    | /users      | Read      |
| PUT    | /users/{id} | Update    |
| DELETE | /users/{id} | Delete    |

Start the web app:

```bash
go run web/main.go
```

Example usage:

```bash
curl -X POST localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"id":1,"name":"Jimmy"}'
```

Itearction on the database via web app: 

![Wen Ineraction](https://github.com/jimmymuthoni/pesapal_rdbms_challenge/blob/aaac16bc07f791d25e57d77c62c6bb3188f5e3cf/accessing_db_via_web.png)

---

## Project Structure

```
pesapal-mini-rdbms/
├── cmd/
│   └── repl/          # Interactive SQL REPL
│       └── main.go
├── database/          # Core RDBMS engine
│   ├── engine.go
│   ├── storage.go
│   ├── schema.go
│   └── index.go
├── sql/               # SQL-like parser
│   └── parser.go
├── web/               # Trivial web CRUD app
│   └── main.go
├── data/              # On-disk database storage
├── schemas.json       # Persisted table schemas
└── README.md
```

#### Conclusion

This project demonstrates:

* Understanding of relational database fundamentals
* Ability to design and implement persistent systems
* Practical use of Go for systems programming
* Translating low-level engines into usable interfaces (REPL + Web)

Even if not used in production, the project serves as a strong learning artifact and a public portfolio piece showcasing real engineering work.


