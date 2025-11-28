# Task CLI — Simple Task Manager in Go

A lightweight command-line task tracker built with pure Go.  
Tasks are stored locally in a JSON file (`tasks.json`).

---

## How to Run

### 1. Clone the repository

```
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>
```

### 2. Build the executable 

From the project root: 
```
go build -o task-cli ./cmd/task-cli
```

### 3. Run the app

Use:
```
./task-cli <command>
```

Examples:

```
./task-cli help
./task-cli add "Buy groceries"
./task-cli list
```


## Commands

### Add task

```
task-cli add <task description>
```

### List tasks

```
task-cli list
task-cli list <status>
```

### Update a task

```
task-cli update <id> <new description>
```

### Mark task as in progress

```
task-cli mark-in-progress <id>
```

### Mark task as done

```
task-cli mark-done <id>
```

### Delete task

```
task-cli delete <id>
```

### Help 

```
task-cli help
```
