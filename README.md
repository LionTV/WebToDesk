```markdown
# WebToDesk

WebToDesk is a Go-based tool to convert websites into standalone desktop applications for macOS and Windows.

## Features

- Convert any website (URL or local files) into a desktop app.
- Customizable window size, title, and debug mode.
- macOS `.dmg` and Windows `.exe` output options.

## Installation

### Prerequisites

- Go 1.16 or higher
- For macOS: `create-dmg` (install via Homebrew: `brew install create-dmg`)
- For Windows: [WiX Toolset](https://wixtoolset.org/) (for `.msi`) or [Inno Setup](https://jrsoftware.org/isdl.php)

### Build Instructions

#### macOS

1. Install `create-dmg` if not already installed:

   ```sh
   brew install create-dmg
   ```

2. Build the app and create a `.dmg`:

   ```sh
   go build -o WebToDesk main.go
   create-dmg WebToDesk
   ```

#### Windows

1. Build the `.exe` file:

   ```sh
   go build -o WebToDesk.exe main.go
   ```

2. (Optional) Create an installer with WiX Toolset or Inno Setup.

## Usage

Convert a website to a desktop app with the following command:

```sh
./WebToDesk <url or folderPath> <AppName>
```

### Example

Convert a URL to a desktop app:

```sh
./WebToDesk "https://example.com" "MyWebApp"
```

Convert a local folder to a desktop app:

```sh
./WebToDesk "./path/to/local/website/" "MyLocalApp"
```
