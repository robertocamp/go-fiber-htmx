# go-fiber-htmx
example site using fiber framework and "GOTH" stack


## project setup

the project module is called "gfgoth" for "Go Fiber GOTH stack"


1. git clone git@github.com:robertocamp/go-fiber-htmx.git
2. cd go-fiber-htmx
3. `go mod init gfgoth`
4. install air (if not already) `go install github.com/air-verse/air@v1.52.3`
5. create default air file at the root of the project: `air init`

## Air


"Air" is a tool used in Go development that provides live reloading of your Go applications. This means that every time you make a change to your code, the tool will automatically rebuild and restart your application, which can significantly speed up the development process by reducing the need to manually rebuild and restart the application yourself.

### Features of Air:
    - **Live Reloading**: Automatically detects changes in your Go source code and recompiles the code.
    - **Configuration**: Allows configuration through an `.air.toml` file where you can specify various options like build commands, log paths, and directories to watch for changes.
    - **Ease of Use**: Simple to set up and integrate into your Go development workflow.

    ### How to Install Air:
    To install Air, you can use the `go install` command. Follow these steps:

    1. **Install Air**:

    ```bash
    go install github.com/air-verse/air@v1.52.3
    ```

    This command will download and install Air from the GitHub repository.

    2. **Add Go bin Directory to PATH**:
    Ensure that your Go binary directory is in your PATH. This is typically `$HOME/go/bin` for installations using the default settings. You can add it to your PATH by adding the following line to your shell configuration file (`.bashrc`, `.zshrc`, etc.):
    ```bash
    export PATH=$PATH:$HOME/go/bin
    ```

    3. **Verify Installation**:
    Check if Air is installed correctly by running:
    ```bash
    air -v
    ```

### How to Use Air:
    1. **Initialize Air**:
    In the root directory of your Go project, you can generate a default configuration file:
    ```bash 
    air init
    ```

    This will create an `.air.toml` file where you can customize the settings if needed.

    2. **Run Air**:
    To start using Air for live reloading, simply run:
    ```bash
    air
    ```

    This will start the Air tool, which will watch your source code for changes and automatically rebuild and restart your application.

### Example Configuration (`.air.toml`):
    Here’s a basic example of what the `.air.toml` file might look like:
    ```toml
    # Config file for air

    [build]
    cmd = "go build -o ./tmp/main"
    bin = "./tmp/main"
    full_bin = ""
    include_ext = ["go", "tpl", "tmpl", "html"]
    exclude_dir = ["assets", "tmp"]
    exclude_file = []
    follow_symlink = true
    args_bin = []

    [log]
    color = "auto"
    time = false

    [serve]
    root = "."
    cmd = "./tmp/main"
    delay = 1000
    grace = 5000
    ignore = ["assets", "tmp"]
    ignore_file = []
    watch_dir = ["."]
    watch_ext = []
    exclude_dir = []
    include_dir = []
    ```

    With this setup, you should be ready to start developing your Go project with the convenience of live reloading provided by Air.
## links

- youTube
    * https://www.youtube.com/watch?v=x7v6SNIgJpE&t=25s