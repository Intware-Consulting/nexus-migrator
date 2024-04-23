# nexus-migrator
<img src="docs/nexus-migrator-banner.jpeg" width="400">

The application can be used to `download` and `upload` raw Nexus repositories.

## Configuration
By default, the application searches for the configuration file in the same directory. The file should be called `nexus.conf`:
```json
{
    "nexus": "https://nexus.intware.eu",
    "user": "api",
    "pass": "api",
    "timeout": "120s",
    "level": "info"
}
```
If `nexus.conf` isn't found, the application will read environment variables:
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
To download the remote Nexus raw repository to a local directory:
```shell
./nexus-migrator download -s software -t /tmp
```

### Upload repository
Upload a local directory to the Nexus raw repository:
```shell
./nexus-migrator upload -s /tmp -t software
```

### Show help
```shell
./nexus-migrator --help
```







