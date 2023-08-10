# Table: kafka_broker

The Kafka broker can be defined as one of the core components of the Kafka architecture. It is also widely known as the Kafka server and a Kafka node.

## Examples

### Basic info

```sql
select
  id,
  addr,
  connected,
  rack,
  response_size
from
  kafka_broker;
```

### List disconnected brokers

```sql
select
  id,
  addr,
  connected,
  rack,
  response_size
from
  kafka_broker
where
  not connected;
```

### List brokers where rack details are unknown

```sql
select
  id,
  addr,
  connected,
  rack,
  response_size
from
  kafka_broker
where
  rack is null;
```