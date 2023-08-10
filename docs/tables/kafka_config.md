# Table: kafka_config

The configuration for the specified resources. The returned configuration includes default values and the Default is true can be used to distinguish them from user supplied values. Config entries where ReadOnly is true cannot be updated. The value of config entries where Sensitive is true is always nil so sensitive information is not disclosed. This operation is supported by brokers with version 0.11.0.0 or higher.

You **_must_** specify `resource_name` and `resource_type` in a `where` clause in order to use this table.

Below are the supported resource types -
- UnknownResource -> 0
- TopicResource -> 2
- BrokerResource -> 4
- BrokerLoggerResource -> 8

## Examples

### Basic info

```sql
select
  name,
  resource_name,
  resource_type,
  value,
  read_only
from
  kafka_config;
```

### Show configurations for all the topics

```sql
select
  c.name,
  resource_name,
  resource_type,
  value,
  read_only
from
  kafka_config as c,
  kafka_topic as t
where
  c.resource_name = t.name
  and c.resource_type = 2;
```

### Show configurations for all the brokers

```sql
select
  c.name,
  resource_name,
  resource_type,
  value,
  read_only
from
  kafka_config as c,
  kafka_broker as b
where
  c.resource_name = b.id
  and c.resource_type = 4;
```

### Show read only configurations for all the topics

```sql
select
  c.name,
  resource_name,
  resource_type,
  value,
  read_only
from
  kafka_config as c,
  kafka_topic as t
where
  c.resource_name = t.name
  and c.resource_type = 2
  and read_only;
```

### Show sensitive configurations for all the topics

```sql
select
  c.name,
  resource_name,
  resource_type,
  value,
  read_only
from
  kafka_config as c,
  kafka_topic as t
where
  c.resource_name = t.name
  and c.resource_type = 2
  and sensitive;
```