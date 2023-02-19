# Domain model

Since any one instance of this framework is intended to support only a single live stream the use of global state can be used to unify similar variables provided by the different services.

```yaml
stream:
  project:             # usually title
  status:
  title:
  privacy:             # video only, audio only, both, none, etc.
annoucement:
message:
  service:
  text:
  tstamp:
community:
  service:
  members:
command:
  name:
  args:
  desc:
  cron:
```
