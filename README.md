# papertrail
Simple PapertrailApp Handler for Golang [apex/log](https://github.com/apex/log) package.

Before usage please register new [PapertrailApp](https://papertrailapp.com/) account.

## Usage

```golang
import "github.com/apex/log"
import "github.com/zerospiel/papertrail"

endpoint := "logs123.papertrailapp.com:55555"

h, err := papertrail.New(endpoint)
if err != nil {
    panic(err)
}

log.SetHandler(h)
ctx := log.WithField("key 1", "value 1")
ctx.Info("put any message here")
```