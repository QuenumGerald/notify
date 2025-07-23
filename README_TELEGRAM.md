# Ignite Notify Plugin – Telegram Sink Guide


This guide explains how to set up and use the **Telegram sink** with the Ignite Notify Plugin to receive blockchain event notifications directly in your Telegram chats or groups.

---

## 1. Create a Telegram Bot

1. Open Telegram and search for [@BotFather](https://t.me/botfather).
2. Start a chat and send `/newbot`.
3. Follow the instructions to set a name and username for your bot.
4. Copy the **token** provided (e.g., `123456789:ABCdefGhIJKlmNoPQRstUvwxYZ`).

---

## 2. Get Your Chat ID

1. Start a chat with your bot or add it to a group.
2. Send any message to the bot.
3. Open this URL in your browser (replace `<TOKEN>`):
   ```
   https://api.telegram.org/bot<TOKEN>/getUpdates
   ```
4. In the JSON response, find `"chat":{"id":...}` and copy the `id` (may be negative for groups).

---

## 3. Construct the Webhook URL

Format:
```
https://api.telegram.org/bot<TOKEN>/sendMessage?chat_id=<CHAT_ID>
```
Example:
```
https://api.telegram.org/bot123456789:ABCdefGhIJKlmNoPQRstUvwxYZ/sendMessage?chat_id=123456789
```

---

## 4. Add a Telegram Subscription

Run:
```sh
ignite add \
  --name mytelegram \
  --node ws://localhost:26657 \
  --query "tm.event EXISTS" \
  --sink telegram \
  --webhook "https://api.telegram.org/bot123456789:ABCdefGhIJKlmNoPQRstUvwxYZ/sendMessage?chat_id=123456789"
```

---

## 5. Run the Notifier

```sh
ignite run
```
You will now receive notifications in your Telegram chat or group when the event matches your query.

---

## 6. Example YAML Config

Your `~/.ignite/notify.yaml` will have an entry like:
```yaml
- name: mytelegram
  node: ws://localhost:26657
  query: "tm.event EXISTS"
  sink: telegram
  webhook: https://api.telegram.org/bot123456789:ABCdefGhIJKlmNoPQRstUvwxYZ/sendMessage?chat_id=123456789
```

---

## Telegram Notification Formatting

All fields from the `result` section are sent in the Telegram notification as a flat `key: value` list (including nested fields).

- The prefixes `data.` and `data.value.` are automatically removed for clarity:
    - `data.value.height` becomes `height`
    - `data.value.step` becomes `step`
    - `data.type` becomes `type`
    - etc.
- Every field present in the event will be included automatically, even new fields added in the future.

Example output in Telegram:

```
height: 3776
step: RoundStepCommit
type: tendermint/event/RoundState
events.tm.event.0: NewRoundStep
query: tm.event EXISTS
```

---

## License
MIT
