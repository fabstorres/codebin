# CodeBin

CodeBin is a lightweight web app for sharing, storing, and previewing code snippets.
It’s inspired by Pastebin and designed for simplicity, speed, and modern web development.
Built with HTMX, Flask (or your backend), and friendly developer UX in mind.

---

## Features

- Create and share code snippets instantly
- Get a unique shareable link for each snippet
- Syntax highlighting for popular languages

---

## Setup

### 1. Clone the repo

```bash
git clone https://github.com/fabstorres/codebin.git
cd codebin
```

### 2. Run the application

```bash
go run main.go
```

The server will start and you can open:
http://localhost:8080

---

## Project Structure

- main.go - Entry point for the Go application
- templates/ - HTML templates rendered by Go
- static/ - Stylesheets, assets, and etc.

## Usage

### Add a New Snippet

1. Go to /new
2. Type or paste your code
3. Save - you’ll get a unique link like /bin/abc123

### View a Snippet

Visit any /bin/{id} URL to load the code view.

---

## Example Environment Variables (optional)

CODEBIN_PORT=8080  
CODEBIN_DATA_PATH=./data

---

## License

MIT License
Copyright (c) 2025 fabstorres

---

Made by [fabs](https://github.com/YourUsername)
