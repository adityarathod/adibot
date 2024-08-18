# Simple Discord bot interface

This is a simple bot that calls the web API exposed by `mlx_lm`.

## Setup

(This assumes you're running on a `*nix` based system. I'm not sure how this works on Windows.)

### Step 1: Build the bot

```bash
go mod download
go build cmd/bot/bot.go
```

### Step 2: Configure the bot

This requires a Discord application to be created [on their portal](https://discord.com/developers/applications).

Once you have set up everything necessary and you have a bot token, make a copy of `bot-config-example.json` and replace the `token` field.

#### Allowlisting configurations

Allowlists can be enabled to serve as guardrails against who can prompt the bot.

There are two main types of allowlists available: `userAllowlist` and `channelAllowlist`. Each can be enabled/disabled independently of each other, and take in an array of IDs (which can be obtained by right-clicking a user/channel with [developer mode enabled](https://support-dev.discord.com/hc/en-us/articles/360028717192-Where-can-I-find-my-Application-Team-Server-ID)).

In general, it's highly recommended to set this up only in a bot testing channel in a server that you control. User allowlists are up to your discretion, but also recommended in larger servers.

#### Ratelimiting configurations

To prevent your machine from getting overloaded (and the bot being too disruptive to existing conversations in channels it can talk in), there's a crude rate-limiter built in under `replyRateLimit` in the config.

This takes in a parameter, which is the proportion of messages to reply to. In the default config it's set to `0.5`, which means the bot will respond to approximately half of messages.

### Step 3: Invite it to your server

On the Discord developer [portal](https://discord.com/developers/applications), under your application's OAuth2 settings, generate a URL that provides the following scopes:

- `applications.commands`
- `bot`

Under the "Bot Permissions" section, enable the following permissions:

- `Send Messages`
- `Add Reactions`
- `Read Message History`
- `Use External Emojis`

Copy this link and open it in a browser tab to invite the bot to your server.

### Step 4: Start up everything

In a terminal, make sure [the model server](../model-training/README.md#step-5-run-a-server) is running.

In another terminal, run `./bot` to launch the server.

You should be able to go to any allowlisted channel and now interact with the bot!
