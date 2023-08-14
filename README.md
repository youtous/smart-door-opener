# Smart Door Opener

[![pipeline status](https://gitlab.com/youtous/smart-door-opener/badges/master/pipeline.svg)](https://gitlab.com/youtous/smart-door-opener/-/commits/master)
[![Docker image](https://img.shields.io/badge/image-registry.gitlab.com%2Fyoutous%2Fsmart--door--opener-e4f0fb?logo=docker)](https://gitlab.com/youtous/smart-door-opener/container_registry)
[![GitHub Repo stars](https://img.shields.io/github/stars/youtous/smart-door-opener?label=✨%20youtous%2Fsmart-door-opener&style=social)](https://github.com/youtous/smart-door-opener/)
[![Gitlab Repo](https://img.shields.io/badge/gitlab.com%2Fyoutous%2Fsmart--door--opener?label=✨%20youtous%2Fsmart-door-opener&style=social&logo=gitlab)](https://gitlab.com/youtous/smart-door-opener/)
[![Licence](https://img.shields.io/github/license/youtous/smart-door-opener)](https://github.com/youtous/smart-door-opener/blob/master/LICENSE)


This application open a door remotely using IFTTT and a FingerBot.
This cheap materials (~40$) enables remote opening of buildings for Airbnb rentals.

## How to?

* Buy a FingerBot and its associated gateway on aliexpress or other : https://www.aliexpress.com/w/wholesale-fingerbot.html.
* Configure Smart Home and link it with a scene on IFTTT : https://ifttt.com/solutions/smart-home
* The IFTTT webhook should be configured to trigger the SmartHome Scene.

## Usage

```text
╰─λ ./tmp/main server -h                                                                                                                                                          0 < 16:54:26
Usage: server server --ifttt-server="maker.ifttt.com" --ifttt-server-key=STRING --ifttt-web-hook-event-name=STRING --access-secret-code=STRING

Start the server.

Flags:
  -h, --help                                Show context-sensitive help.

      --ifttt-server="maker.ifttt.com"      IFTTT server for webhook ($SERVER_IFTTT_SERVER).
      --ifttt-server-key=STRING             IFTTT server key for webhook ($SERVER_IFTTT_SERVER_KEY).
      --ifttt-web-hook-event-name=STRING    IFTTT event name to trigger door opening ($SERVER_IFTTT_WEB_HOOK_EVENT_NAME).
      --access-secret-code=STRING           Secret code used to access the opening page ($SERVER_ACCESS_SECRET_CODE).
```

Host this application on your favorite web hosting server and enjoy!
Feel free to fork for translations.

## License

MIT