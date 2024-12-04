# net-cat
**Welcome to net-cat!**

```bash
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'

```





### About the Project

- This project consists of recreating NetCat in a Server-Client architecture that can run in server mode on a specified port, listening for incoming connections. It can also operate in client mode, connecting to a specified port and transmitting information to the server. NetCat, or nc, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for various networking tasks, including opening TCP connections, sending UDP packets, and listening on arbitrary TCP and UDP ports. For more information about NetCat, you can inspect the manual by running man nc.
Features

**TCP Connection**: Supports multiple clients connecting to a single server (1-to-many).

**User Name Requirement**: Clients must provide a non-empty name upon connecting.

**Connection Control**: Limits the number of simultaneous client connections to 10.

**Message Transmission**: Clients can send messages that are timestamped and prefixed with the sender's name.

**Previous Messages**: New clients receive all previous messages upon joining.

**Join/Leave Notifications**: All clients are notified when someone joins or leaves the chat.

**Error Handling**: Proper error handling for both server and client operations.

**Default Port**: If no port is specified, the application defaults to port 8989.

### Getting Started
- Prerequisites

Go installed on your machine (version 1.15 or higher recommended).

Basic understanding of Go programming language.

#### Installation

    Clone the repository:

```bash
git clone https://github.com/jowala/net-cat.git
cd net-cat
```
Ensure you have Go installed by running:

```bash
go version
```
Running the Server
To start the server, navigate to the project directory and run:

```bash
go run .
```
- This will start the server on the default port 8989. To specify a different port, run:

```bash
go run . <port>
```
- For example:

```bash
go run . 2525
```
- Running the Client
To connect a client to the server, use the following command in a new terminal window:

```bash
nc localhost <port>
```
Replace <port> with the port number your server is listening on (e.g., 8989 or 2525).
- Example Usage

    Start the server:

```bash
go run .
```
    Listening on port :8989

**Open a new terminal and connect a client**:

```bash
nc localhost 8989
```
***Enter your name when prompted***:

text
[ENTER YOUR NAME]: Alice

***Send messages***:

text
[2020-01-20 16:03:43][Alice]: Hello everyone!

**Connect another client**:

```bash
nc localhost 8989
```
- Enter your name for the second client:

text
[ENTER YOUR NAME]: Bob

- Observe messages from both clients.

**Exit the Chat**
To exit the chat, simply close the terminal window or use Ctrl+C in any terminal running a client.

### Contributing
We welcome contributions to improve this project! Here’s how you can help:

- Enhance Features: Suggest or implement new features that improve user experience.
- Improve Documentation: Help us make our documentation clearer and more comprehensive.
- Fix Bugs: Report any bugs you find or submit pull requests with fixes.
- Testing: Write unit tests for both server and client functionalities.

**How to Contribute**

    Fork the repository on GitHub.
    Create a new branch for your feature or bug fix:

```bash
git checkout -b feature/YourFeatureName
```
Make your changes and commit them:

```bash
git commit -m "Add new feature"
```
Push your changes to your forked repository:

```bash
git push origin feature/YourFeatureName
```
    Submit a pull request detailing your changes.

- Known Issues

    Clients may need to press Enter after sending a message to see new messages from other users due to buffering.


### Learning Outcomes
This project will help you learn about:

    Manipulation of structures in Go.
    Understanding TCP/UDP networking.
    Implementing TCP/UDP connections and sockets.
    Utilizing Go concurrency with goroutines and channels.
    Using mutexes for safe concurrent access.
    Working with IP addresses and ports.
