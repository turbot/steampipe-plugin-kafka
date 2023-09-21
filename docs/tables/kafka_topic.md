# Table: kafka_topic

Kafka topics are the categories used to organize messages. Each topic has a name that is unique across the entire Kafka cluster. Messages are sent to and read from specific topics. In other words, producers write data to topics, and consumers read data from topics.

Kafka topics are multi-subscriber. This means that a topic can have zero, one, or multiple consumers subscribing to that topic and the data written to it.

## Examples

### Basic info

```sql
select
  name,
  version,
  err,
  is_internal
from
  kafka_topic;
```

### List partitions of all the topics

```sql
select
  name,
  p ->> 'ID' as partition_id,
  p ->> 'Leader' as partition_leader,
  p ->> 'LeaderEpoch' as partition_leader_epoch,
  p -> 'OfflineReplicas' as partition_offline_replicas,
  p -> 'Replicas' as partition_replicas,
  p ->> 'Version' as partition_version
from
  kafka_topic,
  jsonb_array_elements(partitions) as p;
```

### List internal topics

```sql
select
  name,
  version,
  err,
  is_internal
from
  kafka_topic
where
  is_internal;
```
