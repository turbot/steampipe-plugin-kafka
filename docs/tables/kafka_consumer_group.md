# Table: kafka_consumer_group

A consumer group is a set of consumers which cooperate to consume data from some topics. The partitions of all the topics are divided among the consumers in the group. As new group members arrive and old members leave, the partitions are re-assigned so that each member receives a proportional share of the partitions. This is known as rebalancing the group.

## Examples

### Basic info

```sql
select
  group_id,
  version,
  error_code,
  state,
  protocol_type,
  protocol
from
  kafka_consumer_group;
```

### List dead groups

```sql
select
  group_id,
  version,
  error_code,
  state,
  protocol_type,
  protocol
from
  kafka_consumer_group
where
  state = 'Dead';
```

### List groups with consumer protocol type

```sql
select
  group_id,
  version,
  error_code,
  state,
  protocol_type,
  protocol
from
  kafka_consumer_group
where
  protocol_type = 'consumer';
```

### Show member details of all the groups

```sql
select
  group_id,
  state,
  protocol_type,
  protocol,
  jsonb_pretty(members) as members
from
  kafka_consumer_group;
```
