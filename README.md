# Ignite Notify Plugin

A modular notification plugin for the Ignite CLI and Cosmos SDK ecosystem. It allows you to subscribe to blockchain events from a local node and receive real-time notifications via multiple channels (stdout, Slack, Telegram, ...).

---

## Features

- **Subscribe to Tendermint/Cosmos events** via custom queries
- **Multiple notification sinks**: stdout, Slack, Telegram (extensible)
- **YAML-based persistent configuration**
- **Automatic WebSocket reconnection**
- **Command-line management**: add, list, remove, run subscriptions
- **Modular internal architecture**: config, runner, sink, subscriber
- **Test coverage for all major components**

---

## Architecture

- **cmd/**: CLI entrypoints (`add`, `list`, `remove`, `run`, `autorun`)
- **internal/config**: YAML config management and subscription struct
- **internal/sink**: Sink interface and implementations (stdout, Slack, Telegram)
- **internal/subscriber**: WebSocket subscription logic, sink dispatch
- **internal/runner**: Orchestrates subscriptions and manages lifecycle

---

## Installation

```
git clone https://github.com/your-org/ignite-notify.git
cd ignite-notify
go build -o ignite-notify
```

---

## Usage

### Add a subscription
```bash
ignite-notify add --name mysub --node ws://localhost:26657 --query "tm.event EXISTS" --sink slack --webhook https://hooks.slack.com/services/XXX
```
Example (Telegram):
```bash
ignite-notify add --name mytelegram --node ws://localhost:26657 --query "tm.event EXISTS" --sink telegram --webhook "https://api.telegram.org/bot<token>/sendMessage?chat_id=<chat_id>"
```

### List subscriptions
```
ignite-notify ls
```

### Remove a subscription
```
ignite-notify rm --name mysub
```

### Run all subscriptions
```
ignite-notify run
```

---

## Configuration

Subscriptions are stored in `~/.ignite/notify.yaml` as a list of objects:
```yaml
- name: mysub
  node: ws://localhost:26657
  query: tm.event EXISTS
  sink: slack
  webhook: https://hooks.slack.com/services/XXX
```

- **sink**: One of `stdout`, `slack`, `telegram` (extensible)
- **webhook**: For Slack, use the webhook URL. For Telegram, use the full Bot API URL with token and chat_id.

---

## Extending

- Add new sinks by implementing the `Sink` interface in `internal/sink/sink.go`.
- Add new commands or flags in `cmd/` and register them in `app.go`.

---

## Development & Testing

### Troubleshooting

#### If you see `Unknown command path: ignite add`
- Make sure you have updated the plugin dispatcher in `app.go` to handle both `add` and `ignite add` (and same for other commands).
- Uninstall and reinstall the plugin:
  ```sh
  ignite app uninstall -g /home/nova/Documents/projects/Ignite/ignite-notify
  ignite app install -g /home/nova/Documents/projects/Ignite/ignite-notify
  ```
- If the problem persists, check that you are running the latest code and that the app is properly registered.

### Running Tests

All code is modular and covered by unit tests. Test files are present in each major package (`cmd/`, `internal/config/`, `internal/sink/`, `internal/runner/`).

Run all tests:
```
go test ./...
```

---

## Example: Telegram Sink

To receive notifications on Telegram, create a bot and use the following API URL as webhook:
```
https://api.telegram.org/bot<YOUR_TOKEN>/sendMessage?chat_id=<YOUR_CHAT_ID>
```

---

## Roadmap / TODO
- Add more sinks (Discord, email, ...)
- Improve error handling and reconnection strategies
- Add integration tests and CLI e2e tests
- Document advanced event queries

---

## License
MIT
