# 📝 todo-cli-go

A simple command-line To-Do list manager written in Go.  
Designed for Linux and Unix-based systems.

## ✅ Features

- Add, list, complete, and delete tasks
- Tasks saved locally in `~/.todo/tasks.json`
- Lightweight and cross-platform (Linux/macOS)     

## ⚙️ Installation

### Option 1: Clone and Build from Source

```bash
git clone https://github.com/mkumar2307/todo-cli-go.git
cd todo-cli-go
go build -o todo
```

Then move the binary into your $PATH:     
```bash
mv todo ~/.local/bin/    # or /usr/local/bin/ (may require sudo)
```

Make sure ~/.local/bin is in your PATH:
```bash
export PATH="$HOME/.local/bin:$PATH"
```

### Option 2: Install via go install

```bash
go install github.com/mkumar2307/todo-cli-go@latest
```

This installs the binary in $HOME/go/bin. Add it to your PATH if needed.     

```bash
export PATH="$HOME/go/bin:$PATH"
```

## 🧪 Usage

```bash
todo add "Finish Go project"
todo list
todo done 1
todo delete 2
```

## 📘 Example Output

```bash
$ todo add "Learn Go basics"
Task added.

$ todo list
1. [ ] Learn Go basics

$ todo done 1
Task marked as done.

$ todo list
1. [x] Learn Go basics
```

## 💡 Contributing

Contributions are welcome! Feel free to open issues or PRs for:      

      
	•	Feature requests        
	•	Bug fixes          
	•	Improvements


## 🛡 License

This project is licensed under the MIT License.
