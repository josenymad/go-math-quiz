# Go Math Quiz

This is the first program I've written in Go. It extracts questions and answers from a CSV file and runs through the quiz with a timer.

## Usage 

1. [Download](https://go.dev/dl/) and install Go
2. Fork and clone this repo
3. Navigate to the new directory on you machine
4. Start the quiz using: `go run .`

As a default, the program will use the `problems.csv` file. You can create and add new quizzes using the same format and run those using the `-csv` flag:

`go run . -csv=<new_quiz_file.csv>` 

As a default, the timer is set for 30 seconds. You can change this using the `-limit` flag:

`go run . -limit=<new_time_limit>`

Use the `--help` flag for a breakdown of these flags:

`go run . --help`