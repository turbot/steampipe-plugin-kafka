# Table: kafka_topic_message

A message in Apache Kafka, associated with a specific topic. A topic in Kafka serves as a category or feed name to which records (messages) are published. Messages in Kafka are key-value pairs, and each is stored in a specific topic. As consumers process these messages, they read them from specific topics. Each message within a Kafka topic has a unique offset, which allows consumers to keep track of the messages they've already processed.

## Examples

### Basic info

```sql
select
  topic,
  timestamp,
  key,
  value,
  partition
from
  kafka_topic_message;
```

### List messages on a particular topic

```sql
select
  topic,
  timestamp,
  key,
  value,
  partition
from
  kafka_topic_message
where
  topic = 'mytopic';
```

### List messages which are older than 30 days

```sql
select
  topic,
  timestamp,
  key,
  value,
  partition
from
  kafka_topic_message
where
  timestamp <= (now() - interval '30' day);
```

### List messages of a particular order

```sql
select
  topic,
  timestamp,
  key,
  value,
  partition
from
  kafka_topic_message
where
  key = 'sampleOrder';
```
