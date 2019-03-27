# Rocket.Chat.MQTT
Access RocketChat via MQTT

# Usage
The Following functions are supported:

## Private channel messaging (not implemented jey)
|Name|MQTT|Remark|Implemented|
|----|----|------|-----------|
|Write into private channel|Publish '<prefix>/group/{channel-name}'|Without autojoin|No|
|Receive from private channel|Subscribe '<prefix>/group/{channel-name}'|Without autojoin|No|

## Public channel messaging (not implemented jey)
|Name|MQTT|Remark|Implemented|
|----|----|------|-----------|
|Write into public channel|Publish '<prefix>/channel/{channel-name}'|Without autojoin|No|
|Receive from public channel|Subscribe '<prefix>/channel/{channel-name}'|Without autojoin|No|

## Direct messaging (not implemented jey)
|Name|MQTT|Remark|Implemented|
|----|----|------|-----------|
|Write to user|Subscribe '<prefix>/direct/{username}'||No|
|Receive from user|Subscribe '<prefix>/direct/{username}'||No|

## Quick start


## Configure
