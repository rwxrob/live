# Domain model

Since any one instance of this framework is intended to support only a single live stream the use of global state can be used to unify similar variables provided by the different services.

```yaml
stream:
  title:               # main portion of title (combined with status, mood in actual title)
  tags:                # one word tags to help with finding stream
  status:              # afk, online, muted, etc.
  mood:                # emoji(s) and one word for current mood
  commands:            # commands to hilight in title

annoucement:
  text:
  scope:
  origin:
    service:
    sender:

message:
  tstamp:
  text:
  origin:
    service:
    sender:

community:
  members:
    - name:            # name unique to the overall streaming community
      services:
        - type:        # enum type of service (1-twitch,2-discord,3-youtube, etc.)
          user:        # username or account id on service

command:
  name:
  args:
  desc:
  timeout:

job:
  name:
  args:
  cron:
```
