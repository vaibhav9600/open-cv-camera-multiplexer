Here’s an enhanced and grammatically improved version of your `README.md` file for the Git repository:

---

# Video Stream Multiplexing Using Go Channels and gRPC

This repository contains code to multiplex a single video stream using the Go channels pattern and stream it via gRPC to different services.

## Prerequisites

Before running the application, make sure you have the following dependencies installed:

- **OpenCV**: Required for video processing.
- **pkg-config**: A tool to configure and compile code with libraries like OpenCV.
- **Go**: The Go programming language for building and running the server and clients.

## Setup Instructions

### Step 1: Install OpenCV

The error indicates that the `pkg-config` tool is missing from your system. This tool is essential for configuring and compiling code with libraries like OpenCV. Follow these steps to resolve the issue:

1. **Install `pkg-config`**:

   On macOS, use Homebrew to install `pkg-config` by running the following command in your terminal:

   ```bash
   brew install pkg-config
   ```

2. **Ensure OpenCV is Installed**:

   If OpenCV is not installed, you can install it using Homebrew:

   ```bash
   brew install opencv
   ```

3. **Set Environment Variables**:

   Configure your `PKG_CONFIG_PATH` environment variable so that `pkg-config` can locate OpenCV. Add the following line to your shell configuration file (`~/.zshrc` or `~/.bashrc`):

   ```bash
   export PKG_CONFIG_PATH=$(brew --prefix opencv)/lib/pkgconfig
   ```

   Then, reload your shell configuration:

   ```bash
   source ~/.zshrc   # or source ~/.bashrc if you're using bash
   ```

4. **Test the Configuration**:

   Run the following command to verify that `pkg-config` can locate OpenCV:

   ```bash
   pkg-config --modversion opencv4
   ```

   If the version number is returned, OpenCV is correctly installed.

---

### Step 2: Run the Server and Clients

1. **Start the Server**:

   Open a terminal and navigate to the `server` folder. Start the server by running:

   ```bash
   go run main.go
   ```

   If prompted, grant your code editor permission to access the camera.

2. **Start Client 1 and Client 2**:

   Open two new terminals:

   - In the first terminal, navigate to the `clients/client1` directory and run:

     ```bash
     go run main.go
     ```

   - In the second terminal, navigate to the `clients/client2` directory and run:

     ```bash
     go run main.go
     ```

   Ensure the server is running before starting the clients. After executing the commands, the respective video streams will be visible in each client’s terminal window.

### Step 3: Connection Duration Configuration

The client connection will automatically close after 10 seconds. This behavior is defined in the `server/grpcserver/grpcserver.go` file. You can modify this duration by changing the following value in the server code:

```go
// Adjust the connection timeout duration as needed
```

---

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit pull requests. Please make sure to follow the project's coding standards and write tests for any new features you add.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Let me know if you need further modifications or additions!