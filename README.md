# Gobf

This project is a simple Brainf*ck interpreter written in Go, wrapped in an HTTP server that allows clients to send Brainfuck code for interpretation. The frontend, built with Svelte, provides an intuitive interface for users to input Brainfuck code and receive results in real-time.

## Features

- **Brainfuck Interpreter**: Written in Go for efficient processing.
- **HTTP Server**: Accepts POST requests with Brainfuck code, interprets it, and returns the output.
- **Frontend with Svelte**: Easy-to-use interface for sending Brainfuck code to the server and displaying results.

## Tech Stack

- **Backend**: Go, gorilla mux
- **Frontend**: Svelte

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/AndreiAlbert/gobf.git
   cd brainfuck-interpreter
   cd backend
   go run main.go
   cd frontend
   npm install
   npm run dev
   ```

