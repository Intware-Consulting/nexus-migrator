# nexus-migrator
<img src="docs/nexus-migrator-banner.jpeg" width="400">


Application can be used to `download` and `upload` raw Nexus repositories.

## Configuration
By default application search for configuration file in the same directory. The file should be named `nexus.conf`:
```json
{
    "nexus": "https://nexus.intware.eu",
    "user": "api",
    "pass": "api",
    "timeout": "120s",
    "level": "info"
}
```
If `nexus.conf` isn't found, application will read environment variables:
```shell
NEXUS="https://nexus.intware.eu"
USER="api"
PASS="api"
TIMEOUT="120s"
```
Possible configuration properties:
| Property |                            Description                             |
| -------- | ------------------------------------------------------------------ |
| nexus    | Url to Nexus repository                                            |
| user     | API username                                                       |
| Password | API password                                                       |
| timeout  | HTTP timeout                                                       |
| level    | Log level: trace, debug, info, warn, error, fatal, panic, disabled |

## How to use it

### Download repository
In order remote raw Nexus repository to local directory:
```shell
./nexus-migrator download -s software -t /tmp
```

### Upload repository
In order to upload local directory to raw Nexus repository:
```shell
./nexus-migrator upload -s /tmp -t software
```

### Show help
```shell
./nexus-migrator --help
```







