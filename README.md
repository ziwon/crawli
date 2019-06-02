# crawli

A simple cralwer

<p align="center"><img src="nemo.png" width="200"></p>

## Getting started

```console
$ make build
$ ./bin/crawli
```

### Usage

```sh
Usage:
  crawli [flags]
  crawli [command]

Available Commands:
  collect     Collect data with given worksheet
  config      Create default config file
  help        Help about any command
  version     Print the version number of crawli

Flags:
  -h, --help   help for crawli
```

### Configuration

`crawli config` create `~/.crawli/config/config.toml` as a default configuration.

```
[database]
  host = "localhost"
  password = ""
  port = 5432
  user = "postgres"

[default]
  home = "$HOME/.crawli"

[workers]
  crontab = "0 3 * * *"
  max = 10
  min = 1
```

### Collect

Please describe a worksheet file like `~/.crawli/worksheets/coinmarketcap.yml` to collect data as follows:

#### Worksheet

```yaml
task:
  label: "coinmarketcap"
  url: "https://coinmarketcap.com/all/views/all/"
  allowedDomains:
    - "coinmarketcap.com"
  userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:62.0) Gecko/20100101 Firefox/62.0"
  delay: 0
  async: 0
  trigger: "#currencies-all tbody tr"
  columns:
    - column: "Name"
      selector: ".currency-name-container"
      type: "text"
      primary: true

    - column: "Symbol"
      selector: ".col-symbol"
      type: "text"

    - column: "Price (USD)"
      selector: "a.price"
      attr: "data-usd"
      type: "attr"

    - column: "Volume (USD)"
      selector: "a.volume"
      attr: "data-usd"
      type: "attr"

    - column: "Capacity (USD)"
      selector: ".market-cap"
      attr: "data-usd"
      type: "attr"

    - column: "Change (1h)"
      selector: ".percent-change[data-timespan=\"1h\"]"
      attr: "data-percentusd"
      type: "attr"

    - column: "Change (24h)"
      selector: ".percent-change[data-timespan=\"24h\"]"
      attr: "data-percentusd"
      type: "attr"

    - column: "Change (7d)"
      selector: ".percent-change[data-timespan=\"7d\"]"
      attr: "data-percentusd"
      type: "attr"
```

Then, run `collect` command:

```
./bin/crawli collect
open ~/.crawli/data/coinmarketcap.csv
```

<p align="center">
<img src="output.png"/>
</p>



### TODO

- save collected data to persistent database like postgres, mysql
- daily cronjob task as daemon
