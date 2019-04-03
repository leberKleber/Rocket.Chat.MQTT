# Rocket.Chat.MQTT
Access RocketChat via MQTT

![TRAVIS-CI](https://api.travis-ci.org/leberKleber/Rocket.Chat.MQTT.svg?branch=master)



# Usage
The Following functions are supported:

## Private channel messaging (not implemented yet)
|Name|MQTT|Remark|Implemented|
|----|----|------|-----------|
|Write into private channel|Publish '<prefix>/group/{channel-name}'|Without autojoin|No|
|Receive from private channel|Subscribe '<prefix>/group/{channel-name}'|Without autojoin|No|

## Public channel messaging (not implemented yet)
|Name|MQTT|Remark|Implemented|
|----|----|------|-----------|
|Write into public channel|Publish '<prefix>/channel/{channel-name}'|Without autojoin|No|
|Receive from public channel|Subscribe '<prefix>/channel/{channel-name}'|Without autojoin|No|

## Direct messaging (not implemented yet)
|Name|MQTT|Remark|Implemented|
|----|----|------|-----------|
|Write to user|Subscribe '<prefix>/direct/{username}'| |No|
|Receive from user|Subscribe '<prefix>/direct/{username}'| |No|

## Quick start

## Configure
Configurations must be applied via environment variables:

|Name|Description|Example|
|----|-----------|-------|
|ROCKET_CHAT_WS_URL|Url to RocketChat websocket| wss://chat.rocket.net/websocket|
|ROCKET_CHAT_USERNAME|Username to login|leberKleber|
|ROCKET_CHAT_PASSWORD_HASH|SHA-256 hashed password|4e738ca5563c06cfd0018299933d58db1dd8bf97f6973dc99bf6cdc64b5550bd |
|MQTT_BROKER_URL|URL to mqtt broker |127.0.0.1:1883|
|MQTT_CLIENT_ID|MQTT clientID |customClient4711|
