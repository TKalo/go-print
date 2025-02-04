# Go Print

Go Print is a simple command-line tool that scans a directory for files based on specified include and exclude patterns and generates a structured Markdown report containing the file structure and content. This can be used to pass a project into an LLM to provide context for a conversation.

## Features

✅ **Configurable File Selection** - Specify which files to include or exclude using `config.yaml`.
✅ **Automatic Markdown Generation** - Outputs a structured `.md` file with file listings and content.
✅ **Language Detection** - Applies syntax highlighting for recognized file types.
✅ **Easy to Use** - Just run a single command, and the output is generated automatically.

## Installation

### **Prerequisites**
- Go **1.18+** installed on your system.

### **Clone and Build**
```sh
git clone https://github.com/yourusername/go-print.git
cd go-print
go build -o go-print
```

## Usage

### **Basic Usage**
Run the tool from the project directory:
```sh
./go-print
```
By default, it reads configuration from `config.yaml` and generates `output.md`.

### **Custom Configuration**
You can specify a different config file:
```sh
./go-print -config=myconfig.yaml
```

## Configuration

Modify `config.yaml` to customize the behavior:

```yaml
output_path: output.md  # Where the markdown file will be saved

included_paths:
  - "*"  # Include all files

excluded_paths:
  - *.tmp  # Exclude file types
  - folder # Exclude folder
  - file.txt # Exclude specific file
```

## Project Structure

```
go-print/
│── main.go           # Entry point
│── src/
│   ├── config.go     # Config loader
│   ├── files.go      # File filtering logic
│   └── print.go      # Markdown generator
│── README.md         # This file
```

## How It Works

1. **Loads Configuration**: Reads `config.yaml` for included/excluded files.
2. **Scans Directory**: Finds files matching the criteria.
3. **Generates Markdown**:
   - File structure is written in a hierarchical format.
   - File contents are added with syntax highlighting.

## License

This project is licensed under the **MIT License**.

