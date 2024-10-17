# my-ls

`my-ls` is a Go-based implementation of the Unix `ls` command, designed to list files and directories with similar functionality to the original `ls` command. This project includes common flags like `-l`, `-R`, `-a`, `-r`, and `-t`, allowing users to control how they display file information.

## Features

- **List Directory Contents**: Displays the files and directories in a specified or current directory.
- **Supports Common Flags**:
  - `-l`: List detailed file information.
  - `-R`: Recursively list all files in subdirectories.
  - `-a`: Include hidden files in the output.
  - `-r`: Reverse the order of the file listing.
  - `-t`: Sort files by modification time.
  
## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Flags](#flags)
- [Examples](#examples)
- [File Structure](#file-structure)
- [Contributing](#contributing)
- [License](#license)
- [Authors](#authors)

## Installation

To install and run `my-ls`, you will need to have [Go](https://golang.org/dl/) installed.

1. Clone the repository:
   ```bash
   from gitea:
   git clone 

   from github:
   git clone https://github.com/DavJesse/ls-clone.git
   ```


2. Navigate to the project directory:
    ```bash
    from gitea:
    cd my-ls

    from github:
    cd ls-clone
    ```
3. Run program:
    ```bash
    ./run_my_ls.sh
    ```

## Usage
You can run my-ls with or without specifying a directory. By default, it will display the contents of the current directory.

    
    ./run_my_ls.sh [options] [directory]
    
## Flags
The following flags are supported:

- __-l:__ Displays detailed information about each file, such as permissions, ownership, size, and modification date (similar to ls -l).
- __-R:__ Recursively lists all files in subdirectories (similar to ls -R).
- __-a:__ Includes hidden files (files starting with a dot) in the listing (similar to ls -a).
- __-r:__ Reverses the order of the listing (similar to ls -r).
- __-t:__ Sorts the listing by modification time, newest first (similar to ls -t).

## Examples
1. List files in the current directory:
```bash
    ./run_my_ls.sh
```
2. List all files, including hidden ones:
```bash
    ./run_my_ls.sh -a
```
3. Recursively list files and directories:
```bash
    ./run_my_ls.sh -R
```
4. List files with detailed information:
```bash
    ./run_my_ls.sh -l
```
5. List files sorted by modification time:
```bash
    ./run_my_ls.sh -t
```
6. Reverse the order of listing:

```bash
    ./run_my_ls.sh -r
```
## Combining Flags
You can combine multiple flags as needed:
```bash
    ./run_my_ls.sh -laR
```
This would list all files (including hidden ones) with detailed information, recursively through subdirectories.

## File Structure
The project is organized as follows:
```perl
my-ls/
├── cmd/
│   └── my-ls/
│       └── main.go            # Main entry point for the application
├── internal/
│   ├── ls/
│   │   ├── display.go         # Handles display logic (e.g., -l formatting)
│   │   ├── flags.go           # Parses and manages command-line flags
│   │   ├── file_info.go       # Manages file metadata
│   │   ├── sorter.go          # Sorts files (e.g., by time, name, etc.)
│   │   └── recursive.go       # Handles recursive directory traversal
├── tests/
│   ├── ls_test.go             # Unit tests for ls functionality
├── go.mod                     # Module file for managing dependencies
└── README.md                  # Project documentation
```
- ```cmd/my-ls/main.go:``` The main entry point of the application.
- ```internal/ls:``` Contains core logic for file listing, flag parsing, sorting, and recursive functionality.
- ```tests/ls_test.go:``` Unit tests to ensure the correctness of the my-ls implementation.
## Contributing
We welcome contributions to improve my-ls!

1. Fork the repository.
2. Create a new branch:
    ```bash
    git checkout -b feature/my-new-feature
    ```
3. Make your changes.
4. Write tests to cover your changes.
5. Commit your changes
    ```bash
    git commit -m "Add some feature"
    ```
    or
    ```bash
    ./commit.sh "Add some feature"
    ```
6. Push to the branch:
    ```bash
    git push origin feature/my-new-feature
    ```
7. Open a pull request.

## Branch Naming Convention
- `feature/` for new features (e.g., feature/add-sorting).
- `bugfix/` for bug fixes (e.g., bugfix/fix-hidden-files).
- `hotfix/` for urgent fixes (e.g., hotfix/critical-error).
- `refactor/` for code refactoring (e.g., refactor/cleanup-code).

Please ensure that your code follows best practices and is properly documented.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.

# Authors
David Jesse Odhiambo
-
Apprentice Software Deveoper,
Zone01 Kisumu

Mokwa Moffat
-
Apprentice Software Deveoper,
Zone01 Kisumu