## **📄 README.md (Updated Usage Documentation)**
Here’s a **clean and well-structured README** you can use:

```markdown
# Terminal Viewport Pager

A terminal-based **scrollable text viewer** built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).  

Supports:
- **File Viewing** (`pager file.txt`)
- **Piped Input** (`cat file.txt | pager`)
- **Direct Usage as a Go Package** (`viewportpager.StartPager(content, title)`)

---

## **📦 Installation**
### **1. Install the CLI**
First, install the CLI globally:
```sh
go install github.com/grantis/pagercli/cmd/pager@latest
```
Now you can run:
```sh
pager file.txt
```

---

### **2. Using `viewportpager` as a Package**
If you want to use this as a **Go package** in your own projects:
```sh
go get github.com/grantis/viewportpager
```

Then, in your Go code:
```go
package main

import (
    "log"
    "github.com/grantis/viewportpager"
)

func main() {
    content := "Hello, this is a terminal pager example!"
    title := "Demo Pager"

    err := viewportpager.StartPager(content, title)
    if err != nil {
        log.Fatal(err)
    }
}
```

Run it:
```sh
go run main.go
```

---

## **🚀 Features**
✅ **Scrollable text viewport**  
✅ **Supports keyboard & mouse input**  
✅ **Works with both files and piped input**  
✅ **Reusable as a Go package**  

---

## **📌 CLI Usage**
### **View a File**
```sh
pager file.txt
```

### **View Piped Input**
```sh
cat file.txt | pager
```

### **Run Without Input**
```sh
pager
```
(_Displays a default welcome message._)

---

## **⌨️ Controls**
| Key        | Action               |
|------------|----------------------|
| `↑` / `↓`  | Scroll up/down       |
| `PgUp` / `PgDn` | Fast scroll |
| `q` / `esc` / `ctrl+c` | Quit |

---

## **🔧 Development**
### **1. Clone the Repository**
```sh
git clone https://github.com/grantis/terminal-viewport.git
cd terminal-viewport
```

### **2. Install Dependencies**
```sh
cd viewportpager
go mod tidy
cd ../pagercli
go mod tidy
```

### **3. Run the CLI Locally**
```sh
go run cmd/pager/main.go file.txt
```

### **4. Install the CLI**
```sh
go install ./cmd/pager
```
Now you can run:
```sh
pager file.txt
```


---

## **📜 License**
MIT License © 2024 [Grant Rigby](https://github.com/grantis)
```

---

### **✅ Next Steps**
1. **Replace `grantis` with your GitHub username**.
2. **Push this updated README to GitHub**:
   ```sh
   git add README.md
   git commit -m "Updated usage documentation"
   git push origin main
   ```
3. **Start Using It!**
   ```sh
   pager file.txt
   cat file.txt | pager
   ```

---

### **🎯 Now Your Project is Fully Ready!**
- ✅ **CLI is globally installable**
- ✅ **Library is importable in other projects**
- ✅ **Documentation is complete**
- ✅ **Code is modular & clean**

**What’s next?** Do you want to add **RSS support**, **themes**, or **new features**? 🚀🔥