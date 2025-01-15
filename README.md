# Coding Challenges

I'll be going through John Crickett's [Coding Challenges](https://codingchallenges.fyi/) and implementing them in Go.

## Prerequisites

- You will need [Go](https://go.dev/doc/install) installed - ideally a version >= `1.23.4`.
- If you want to use the `go install` method of running the program, then your PATH needs to have the path to the Go install directory defined.

## Usage

### Build from Source

There are two options for building from source - cloning the repository or downloading an release archive file.

#### Clone Repo

1. Clone this repository
2. Navigate to the cloned repository directory, named `coding-challenges`.
3. Navigate to the desired challenge directory
4. Run `go build`
5. Run the compiled executable with `./<INSERT_CHALLENGE_DIRECTORY_NAME>`

### Download from GitHub Releases

1. Navigate to https://github.com/neil-b-patel/coding-challenges/releases/
2. Navigate to the desired release version 
3. Click either the `zip` or `tar.gz` to download an archive file of the source code
4. Navigate to the downloaded archive file
5. Unarchive the file
6. Navigate to the desired challenge directory
7. Run `go build`
8. Execute the program one of three ways:
  - Run `go run .`
  - Run the compiled executable with `./<INSERT_CHALLENGE_DIRECTORY_NAME>`
  - Run `go install` and you can run `<INSERT_CHALLENGE_DIRECTORY_NAME>` from any directory

### Install Go Package

[TODO]: Add exported functions to each package

This option still needs to be implemented.
The goal is to allow other Go modules to also leverage this module's packages directly in-code.
Since this is a learning project, this is not a high priority issue.
