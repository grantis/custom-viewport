### **Go Viewport Pager - Usage Documentation**  
A simple terminal-based viewport pager built using [Bubble Tea](https://github.com/charmbracelet/bubbletea). This allows users to scroll through text in a styled viewport with exit instructions.

---

## **Installation**  
First, install the package in your Go project:

```sh
go get github.com/yourusername/yourrepo@latest
```

---

## **Usage**

### **Quick Start**
You can display text in a terminal viewport by calling the `StartPager` function:

```go
package main

import (
	"fmt"
	"os"

	view "github.com/grantis/custom-viewport"
)

func main() {
	content := "Hello, this is a quick terminal viewport!"
	err := yourrepo.StartPager(content)
	if err != nil {
		fmt.Println("Error starting pager:", err)
		os.Exit(1)
	}
}
```

This will display the text inside a scrollable terminal viewport.

---

## **Usage with Piped Input**
You can also use the program to accept input from stdin:

```sh
cat somefile.txt | go run main.go
```

Or directly from a command:

```sh
echo "This is a test message" | go run main.go
```

---

## **Keybindings**
| Key        | Action                              |
|------------|-------------------------------------|
| `↑` / `↓`  | Scroll up/down                     |
| `PgUp` / `PgDn` | Scroll up/down faster      |
| `q` / `esc` / `ctrl+c` | Quit the viewport |

---

## **API Reference**
### **StartPager(content string)**
Displays the given `content` in a terminal viewport.

#### **Parameters**
- `content` *(string)*: The text to display inside the viewport.

#### **Returns**
- `error`: Returns an error if the pager fails to initialize.

#### **Example**
```go
err := view.StartPager("Hello World!")
if err != nil {
    log.Fatal(err)
}
```

---

## **Advanced Usage**
If you want to integrate this with another CLI application, you can modify the text dynamically before passing it to `StartPager()`.

```go
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/yourusername/yourrepo"
)

func main() {
	var content string

	// Check if data is piped in
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("could not stat STDIN:", err)
		os.Exit(1)
	}

	if fi.Mode()&os.ModeCharDevice == 0 {
		// Read from stdin
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("could not read from STDIN:", err)
			os.Exit(1)
		}
		content = string(bytes)
	} else {
		// Default text if no input is piped
		content = "Welcome to the Go Viewport Pager!"
	}

	// Start the pager
	err = yourrepo.StartPager(content)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
```

---

## **Contributing**
1. Fork the repository.
2. Create a feature branch: `git checkout -b my-new-feature`
3. Commit changes: `git commit -am 'Add new feature'`
4. Push the branch: `git push origin my-new-feature`
5. Open a Pull Request.

---

## **License**
This project is licensed under the MIT License.

---

This documentation provides everything needed to install, use, and contribute to your terminal viewport package. Let me know if you want any refinements! 🚀