# Cornix-tv-channel

[![License: GPLv3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://github.com/gibigo/cornix-tv-channel/blob/master/LICENSE) 
[![testing](https://github.com/gibigo/cornix-tv-channel/actions/workflows/testing.yml/badge.svg)](https://github.com/gibigo/cornix-tv-channel/actions/workflows/testing.yml)
[![release](https://github.com/gibigo/cornix-tv-channel/actions/workflows/release.yml/badge.svg)](https://github.com/gibigo/cornix-tv-channel/actions/workflows/release.yml)
[![docker](https://github.com/gibigo/cornix-tv-channel/actions/workflows/docker.yml/badge.svg)](https://github.com/gibigo/cornix-tv-channel/actions/workflows/docker.yml)

Tradingview forwarder optimized for cornix.

## 💡 About

This project is a highly customizable webhook server which is optimized to use with cornix. It has an API where you can set up different channels and set a strategy for the entry, take-profit, and stoploss price.

## 🚀 Getting started
### ⚡️ Requirements
In order to use ctvc (cornix-tv-channel) you need 
- a server (preferably linux)
- public IP  
- either port 80 or 443 available

### 📱 Prerequisites

- create a telegram bot, see [here](https://core.telegram.org/bots#6-botfather)

## 🧑‍💻 Deployment 
### 🐳 Docker

Docker can be used for quick and easy deployment. 

To install `docker` and `docker-compose`, take a look at [this](https://docs.docker.com/engine/install/) and [this](https://docs.docker.com/compose/install/).

**Create the docker-compose.yml file**

```yml
---
version: '3.7'
services:
  ctvc:
    image: jon4hz/cornix-tv-channel
    restart: unless-stopped
    hostname: ctvc
    container_name: ctvc
    ports:
      - '80:3000'
    volumes:
      - ./data:/data
    environment:
        TELEGRAM_TOKEN: 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw # you telegram bot token
        REGISTRATION: true # enables user registration
        LOG_LEVEL: info # set the log level, (debug,info,warn)
        DATABASE_DEBUG: false # enable debug mode for the database
```

**Excute the service**   
`docker-compose up -d`

**Stop the service**  
`docker-compose down`

### 📦 Binary

Every github release has binaries attached for the following operating systems:
- Linux
- MacOS
- Windows

To run ctvc as a binary follow those steps:
1. Download the [release](https://github.com/gibigo/cornix-tv-channel/releases) for your system and unzip the executable
2. Place it in a folder where you can run it. A subfolder is created for the database
3. Set either the required [enviroinment variables](https://github.com/gibigo/cornix-tv-channel/tree/master#-environment-variables) or write a [config file](https://github.com/gibigo/cornix-tv-channel/tree/master#-config-file)
4. Run the executable

### 👩‍💻 Build locally 

Since the code is fully open source, you can also build the binary yourself.  

As a prerequisite, please configure your go development environment and enable go modules. 

1. Clone the repository  
`git clone https://github.com/gibigo/cornix-tv-channel`
2. Change direcory  
`cd cornix-tv-channel`
3. Download all go modules  
`go mod download`
4. Build the binary  
`go build .`
5. Set either the required [enviroinment variables](https://github.com/gibigo/cornix-tv-channel/tree/master#-environment-variables) or write a [config file](https://github.com/gibigo/cornix-tv-channel/tree/master#-config-file)
6. Run the binary  
`./cornix-tv-channel`

### 🌱 Environment variables

| Variable       | Required? | Description                  |
|----------------|-----------|------------------------------|
| TELEGRAM_TOKEN | yes       | The token from your telegram |
| REGISTRATION   | yes       | Whether user registration is enabled or not. <br>If you don't want other users on your server, set this to false after creating your own user | 
| LOG_LEVEL      | no        | Default: "info", the log level of the service. <br>Other options: "debug", "warn"
| DATABASE_DEBUG | no        | Default: "false", enables the debug mode for the database  


### 📜 Config file

If you prefer to use a config file instead of environment variables, you can create one at `config/config.yml`

#### Example

```yml
---
registration: true
database:
  debug: false
logging:
  logLevel: info
telegram:
  token: 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
```

### 🔒 Secure connection

In order to have a secure connection to the server a reverse proxy with a letsencrypt certificate (e.g. [traefik](https://doc.traefik.io/traefik/),[caddy](https://caddyserver.com/docs/) or [swag](https://docs.linuxserver.io/general/swag)) is recommended. 

## 👨‍💼 Usage

Once ctvc is up and running, you can take a look at the swagger documention which is located at [http://yoururl/swagger/index.html](http://yoururl/swagger/index.html).   
You will need these API endpoints to configure your channel.  


1. Create your user 
2. Create a channel
3. Create a strategy for the new channel
4. Ensure that the previously create telegram bot is in your channel
5. Point the webhook url on trading view to [http://yoururl/webhook/v1](http://yoururl/webhook/v1) or [https://yoururl/webhook/v1](https://yoururl/webhook/v1) if you use a reverse proxy.  
6. Set the message to  
```json
{
    "user": "<your_username>",
    "uuid": "<your_users_uuid>",
    "ticker": "{{ticker}}",
    "price": {{close}},
    "exchange": "{{exchange}}",
    "direction": "<long/short>",
    "tgChannel": <yourTelegramsChannelID>
}
```
and replace all values in "<>"


## 📜 Licensing
This project is released under the GPLv3-License found in the [LICENSE](https://github.com/gibigo/cornix-tv-channel/blob/master/LICENSE) file.