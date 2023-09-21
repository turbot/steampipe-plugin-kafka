![image](https://hub.steampipe.io/images/plugins/turbot/kafka-social-graphic.png)

# Kafka plugin for Steampipe

Use SQL to instantly query Kafka Topics, Brokers and more. Open source CLI. No DB required.

- **[Get started ->](https://hub.steampipe.io/plugins/turbot/kafka)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/kafka/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-kafka/issues)

## Quick start

### Install

Download and install the latest Kafka plugin:

```shell
steampipe plugin install kafka
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/kafka#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/kafka#configuration).

Configure your account details in `~/.steampipe/config/kafka.spc`:

```hcl
connection "kafka" {
  plugin = "kafka"

  # Authentication information
  # bootstrap_servers = ["localhost:9092"]
}
```

Run steampipe:

```shell
steampipe query
```

List your Kafka Topics:

```sql
select
  name,
  version,
  is_internal,
  jsonb_array_length(partitions) as number_of_partitions
from
  kafka_topic;
```

```
+--------------------+---------+-------------+----------------------+
| name               | version | is_internal | number_of_partitions |
+--------------------+---------+-------------+----------------------+
| mytopic            | 5       | false       | 1                    |
| __consumer_offsets | 5       | true        | 50                   |
+--------------------+---------+-------------+----------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-kafka.git
cd steampipe-plugin-kafka
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/kafka.spc
```

Try it!

```
steampipe query
> .inspect kafka
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-kafka/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Kafka Plugin](https://github.com/turbot/steampipe-plugin-kafka/labels/help%20wanted)
