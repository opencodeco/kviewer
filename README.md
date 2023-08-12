# kviewer

The `kviewer` is a command-line tool built in Go to consume Kafka messages in a formatted and user-friendly manner.

# How to Use
To consume messages from a topic:

```bash
./bin/kviewer consume <my-topic> [number-of-messages-to-be-readed]
```

If the number of messages isn't specified, the kviewer will consume indefinitely until interrupted.

# Configuration
For instance, you can define the bootstrap server address using the environment variable `KVIEWER_BOOTSTRAP_SERVERS` with a default fallback to `localhost:29092`.

# How to Contribute
- Fork the project on GitHub.
- Clone the fork to your local machine.
- Create a branch for your new feature or bug fix.
- Make your changes and commit them to your branch.
- Push your branch to your fork on GitHub.
- Create a Pull Request from your fork to the original repository

All contributions are welcome! Please send PRs for improvements and bug fixes. If you have any questions, feel free to open an issue.
