# Task Tracker CLI

A small command-line task manager written in Go as a practice project. Tasks are saved locally in `tasks.json`.

## Prerequisites

- [Go 1.24 or newer](https://go.dev/dl/)
- Git (only needed to clone the repository)

## Run the app

```bash
git clone https://github.com/dazakharova/task-tracker-cli.git
cd TaskTrackerCLI
go build -o task-cli ./cmd/task-cli
./task-cli help
```

On Windows, build with `go build -o task-cli.exe ./cmd/task-cli` and run commands with `./task-cli.exe`.

## Try the commands

Run these from the project root:

```bash
./task-cli add "Prepare for interview"
./task-cli list
./task-cli update 1 "Prepare Go project for interview"
./task-cli mark-in-progress 1
./task-cli list "in progress"
./task-cli mark-done 1
./task-cli list done
./task-cli delete 1
```

Task IDs are shown by `list`. Available statuses are `todo`, `in progress`, and `done`.

## Run the tests

```bash
go test ./...
```

The tests use temporary files, so they do not change your local `tasks.json`.
