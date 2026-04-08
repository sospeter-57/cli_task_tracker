# CLI Task Tracker

A simple command-line task management application built with Go. This tool allows you to manage your tasks directly from the terminal, tracking their status, creation, and update times.

## Features

- **Add Tasks**: Create new tasks with descriptions.
- **List Tasks**: View all tasks or filter by status (todo, in-progress, done).
- **Update Tasks**: Modify the description of existing tasks.
- **Delete Tasks**: Remove tasks by ID.
- **Mark Status**: Change task status to done or in-progress.
- **Persistent Storage**: Tasks are stored in a JSON file (`logs.json`) for data persistence.

## Installation

1. Ensure you have Go installed (version 1.24.3 or later).
2. Clone the repository:
   ```bash
   git clone <repository-url>
   cd cli_task_tracker
   ```
3. Build the application:
   ```bash
   go build -o task-tracker main.go
   ```

## Usage

Run the application with the following commands:

### Add a Task
```bash
./task-tracker add "Buy groceries"
```

### List All Tasks
```bash
./task-tracker list
```

### List Tasks by Status
```bash
./task-tracker list todo
./task-tracker list in-progress
./task-tracker list done
```

### Update a Task
```bash
./task-tracker update <id> "New description"
```

### Delete a Task
```bash
./task-tracker delete <id>
```

### Mark a Task as Done
```bash
./task-tracker mark-done <id>
```

### Mark a Task as In-Progress
```bash
./task-tracker mark-in-progress <id>
```

Replace `<id>` with the actual task ID (a number).

## Task Structure

Each task includes:
- **ID**: Unique identifier (auto-generated).
- **Description**: Task details.
- **Status**: Current state (todo, in-progress, done).
- **Created At**: Timestamp when the task was created.
- **Updated At**: Timestamp when the task was last modified.

## Data Storage

Tasks are stored in `logs.json` in the same directory as the executable. The file is automatically created if it doesn't exist.

## Contributing

Feel free to submit issues or pull requests to improve the application.

## License

This project is open-source. See the license file for details.

