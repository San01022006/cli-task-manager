# CLI Task Manager

A fast and lightweight terminal-based task manager written in Go using Cobra.

Features:

- Add tasks with priorities
- Mark tasks as completed
- Delete tasks
- View pending or all tasks
- Clear completed tasks
- Task statistics and progress tracking
- JSON persistence
- Colored terminal output
- Docker support
- GitHub Actions CI

---

## Installation

### Clone Repository

```bash
git clone https://github.com/San01022006/cli-task-manager.git

cd cli-task-manager
```

### Install Dependencies

```bash
go mod tidy
```

### Build

```bash
go build -o tasks
```

---

## Quick Start

Add a task:

```bash
./tasks add "Learn Go" --priority high
```

List pending tasks:

```bash
./tasks list
```

Mark task complete:

```bash
./tasks done ab12cd34
```

Delete a task:

```bash
./tasks delete ab12cd34
```

Show statistics:

```bash
./tasks stats
```

---

# Commands

## Add Task

Syntax:

```bash
tasks add "Task title"
```

Default priority:

```bash
tasks add "Read book"
```

Output:

```text
✓ Task added [8f21ac4d]
```

---

## Add Task With Priority

Low:

```bash
tasks add "Buy groceries" --priority low
```

Medium:

```bash
tasks add "Finish assignment" --priority medium
```

High:

```bash
tasks add "Production deployment" --priority high
```

Short flag:

```bash
tasks add "Fix critical bug" -p high
```

---

## List Pending Tasks

```bash
tasks list
```

Example:

```text
+----------+---------------------+----------+---------+------------+
| ID       | TITLE               | PRIORITY | STATUS  | CREATED    |
+----------+---------------------+----------+---------+------------+
| ab12cd34 | Learn Go            | HIGH     | Pending | 2026-06-14 |
| ef56gh78 | Build CLI Project   | MEDIUM   | Pending | 2026-06-14 |
+----------+---------------------+----------+---------+------------+
```

---

## List All Tasks

Includes completed tasks.

```bash
tasks list --all
```

---

## Mark Task Completed

```bash
tasks done ab12cd34
```

Output:

```text
✓ Task completed
```

---

## Delete Task

Long form:

```bash
tasks delete ab12cd34
```

Alias:

```bash
tasks rm ab12cd34
```

Alias:

```bash
tasks del ab12cd34
```

Output:

```text
✓ Task deleted
```

---

## Clear Completed Tasks

Removes all completed tasks.

```bash
tasks clear
```

Output:

```text
✓ Completed tasks removed
```

---

## Task Statistics

```bash
tasks stats
```

Example:

```text
Total Tasks      : 20
Completed Tasks  : 12
Pending Tasks    : 8
High Priority    : 5
Progress         : 60.00%
```

---

# Data Storage

Tasks are stored locally in:

```text
~/.cli-tasks.json
```

Example JSON:

```json
[
  {
    "id": "ab12cd34",
    "title": "Learn Go",
    "done": false,
    "priority": "high",
    "created_at": "2026-06-14T12:30:00Z"
  }
]
```

---

# Help

Display available commands:

```bash
tasks --help
```

Display command-specific help:

```bash
tasks add --help
```

```bash
tasks list --help
```

```bash
tasks stats --help
```

---

# Makefile Usage

Install dependencies:

```bash
make deps
```

Build application:

```bash
make build
```

Run application:

```bash
make run
```

Install binary:

```bash
make install
```

Clean build files:

```bash
make clean
```

---

# Docker Usage

## Build Image

```bash
docker build -t cli-task-manager .
```

---

## Run Help

```bash
docker run --rm cli-task-manager --help
```

---

## Add Task

```bash
docker run --rm \
-v $HOME:/root \
cli-task-manager \
add "Learn Docker" --priority high
```

---

## List Tasks

```bash
docker run --rm \
-v $HOME:/root \
cli-task-manager \
list
```

---

# Project Structure

```text
cli-task-manager/
│
├── cmd/
│   └── root.go
│
├── .github/
│   └── workflows/
│       └── build.yml
│
├── main.go
├── go.mod
├── Dockerfile
├── Makefile
├── README.md
└── .gitignore
```

---

# CI/CD

GitHub Actions automatically:

- Installs dependencies
- Builds the project
- Runs tests
- Verifies successful compilation

on every push to:

```text
main
```

---

# Tech Stack

- Go 1.21
- Cobra CLI
- UUID
- TableWriter
- Fatih Color
- Docker
- GitHub Actions

---

# License

MIT License

Feel free to fork, modify, and use in your own projects.
