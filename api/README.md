# Project Automation Guide

This document explains how to automate key tasks in the project using predefined
commands. These tasks include building the application, running it, and managing
the database. Below is a breakdown of each command and its purpose.

## Available Commands

### `build`

- **Purpose**: Compiles the Go application and prepares the necessary executables.
- **Actions**:

  - Builds the main server into the `dist` directory.
  - Builds the file server binary.

  ### `run`

- **Purpose**: Launches the main application after building it.
- **Dependencies**: `build`
- **Actions**:

  - Clears the terminal and runs the application from the `dist` directory.

### `run-file`

- **Purpose**: Starts the file server.
- **Dependencies**: `build`
- **Actions**:

  - Clears the terminal and runs the file server binary.

### `build-seeder`

- **Purpose**: Compiles the database seeder tool.
- **Actions**:

  - Compiles the `seeder` executable for database population.

### `init-db`

- **Purpose**: Initializes the database schema.
- **Dependencies**: `build-seeder`
- **Actions**:

  - Runs the seeder with the `--init-db` flag to initialize the database structure.

### `seed-<entity>`

- **Purpose**: Seeds the database with sample data.
- **Dependencies**: `build-seeder`
- **Actions**:

  - Seeds various data types into the database using the following
    flags: - `--seed-users` - `--seed-groups` - `--seed-roles` -
    `--seed-resources` - `--seed-products` -
    `--seed-product-categories` - `--seed-permission`

### `nuke-db`

- **Purpose**: Completely clears the database.
- **Dependencies**: `build-seeder`
- **Actions**:

  - Runs the seeder with the `--nuke-db` flag to wipe all data from the database.

### `refresh-db`

- **Purpose**: Refreshes the database, clearing and re-seeding it.
- **Dependencies**: `build-seeder`
- **Actions**:

  - Runs the seeder with the `--refresh-db` flag to reset the database.

  ## Usage Examples

- **Build and run the app**: `bash run `

#### Initialize or seed the database:

```bash
    init-db # Initialize the DB schema
    seed-users # Seed users into the DB
    seed-groups # Seed groups into the DB
```

#### Clear or refresh the database:

```bash
    nuke-db      # Wipe all data from the DB
    refresh-db   # Reset and reseed the DB
```
