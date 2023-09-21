---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/kafka.svg"
brand_color: "#5CA2A2"
display_name: "Kafka"
short_name: "kafka"
description: "Steampipe plugin for querying Kafka Topics, Brokers and other resources."
og_description: Query Kafka with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/kafka-social-graphic.png"
---

# Kafka + Steampipe

[Kafka](https://kafka.apache.org/) is an open-source distributed event streaming platform used by thousands of companies for high-performance data pipelines, streaming analytics, data integration, and mission-critical applications.

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/kafka/tables)**

## Quick start

### Install

Download and install the latest Kafka plugin:

```sh
steampipe plugin install kafka
```

### Credentials

| Item        | Description                                                                              |
| ----------- | ---------------------------------------------------------------------------------------- | --- |
| Credentials | Kafka plugin requires `bootstrap_servers` for all requests.                              |
| Permissions | NA                                                                                       |
| Radius      | Each connection represents one Kafka cluster.                                            |     |
| Resolution  | Credentials explicitly set in a steampipe config file (`~/.steampipe/config/kafka.spc`). |

### Configuration

Installing the latest Kafka plugin will create a config file (`~/.steampipe/config/kafka.spc`) with a single connection named `kafka`:

Configure your account details in `~/.steampipe/config/kafka.spc`:

```hcl
connection "kafka" {
  plugin = "kafka"

  # bootstrap_servers - A list of host:port addresses that will be used to discover the full set of alive brokers. Required.
  # bootstrap_servers = ["localhost:9092"]
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-kafka
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
