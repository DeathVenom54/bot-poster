# bot-poster

This tool allows you to send a message to your Discord server through
a bot. So if you need to post as a different persona for making
your server more interesting, we've got you covered!

## Required parameters

This app requires these to send a message:

- Bot token ([how to get one](https://www.writebots.com/discord-bot-token/))
- [Channel ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID)
- Your message (optional)
- An attachment (image/gif URL)

## How to format messages

### Emoji

1. Type the emoji you want to use in Discord
2. Add a backslash (` \ `) before it and send it. If it is a default emoji, you will see a unicode emoji like `😃`. If it is a custom emoji, you will see the emoji name and ID like this: `<:Patreon:882723233042935848>`
3. Copy-paste the text you get as it is, it should work.

### Mentioning members

If you want to mention a user, [get their ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID) and insert `u{user ID}` in your message. An example: `u289677961349824512`.

### Mentioning roles

If you want to mention a role, [get its ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID) and insert `r{role ID}` in your message. An example: `r880746460801536000`.

### Mentioning channels

If you want to mention a channel, [get its ID](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID) and insert `c{channel ID}` in your message. An example: `c824880567618109480`.

### Formatted time

To send nice formatted time, [check out this website](https://r.3v.fi/discord-timestamps/), select your time and format and copy-paste the generated code in your message.
