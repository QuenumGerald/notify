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

### For a Private Chat:

1. Start a chat with your bot.
2. Send any message to the bot.
3. Add [@userinfobot](https://t.me/userinfobot) or [@getidsbot](https://t.me/getidsbot) to your contacts and start a chat.
4. The bot will display your personal chat ID (a positive number like `123456789`).

### For a Group:

1. **IMPORTANT: Add your bot to the group first!**
2. Make sure someone sends a message in the group, mentioning the bot (e.g., "@yourbot hello").
3. Add [@getidsbot](https://t.me/getidsbot) to the group.
4. The bot will display the group chat ID (usually a negative number like `-123456789`).

### Alternative Method:

You can also get IDs via API:
```
https://api.telegram.org/bot<TOKEN>/getUpdates
```
This shows recent interactions with your bot, including chat IDs.

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

### Verify Your Webhook

Before setting up the Ignite subscription, test your webhook with:

```bash
curl -X POST "https://api.telegram.org/bot<TOKEN>/sendMessage?chat_id=<CHAT_ID>&text=Test+message"
```

If you receive the message in Telegram, your webhook is correctly configured!

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

## Troubleshooting

### Common Errors

1. **"Bad Request: chat not found"**
   - The bot has not been added to the group or chat
   - The chat_id is incorrect
   - Solution: Add the bot to the group first, then verify the chat_id

2. **"Forbidden: bots can't send messages to bots"**
   - You are trying to send messages to the bot itself
   - Solution: Use a human user's chat_id or a group chat_id, never the bot's own ID

3. **Bot not visible in group contacts**
   - Make sure you've interacted with the bot in a private chat first
   - For groups, the bot might need privacy mode disabled:
     1. Contact [@BotFather](https://t.me/botfather)
     2. Send `/setprivacy`
     3. Select your bot
     4. Choose "Disable"

4. **Messages not being received**
   - Verify the webhook with curl command (see "Verify Your Webhook" section)
   - Check if your blockchain node is producing events
   - Try with `--sink stdout` first to debug

---

## License
MIT
