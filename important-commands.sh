# creating a topic
./kafka-topics.sh --create --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic

# alter the number of partitions of a topic
./kafka-topics.sh  --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic --alter --partitions 3

# describe a topic
./kafka-topics.sh --describe --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic

# produce messages to a topic
./kafka-console-producer.sh --broker-list broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic

# consume messages from a topic
./kafka-console-consumer.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic --from-beginning

# delete a topic
./kafka-topics.sh --delete --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic

# list all topics
./kafka-topics.sh --list --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092

# Check message in each partition
./kafka-console-consumer.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic --partition 0 --from-beginning

# consume with offset
./kafka-console-consumer.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic --partition 0 --offset 0

# create a consumer group
./kafka-consumer-groups.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --group test-group --reset-offsets --to-earliest --execute --topic test-topic

# check the offset of a consumer group
./kafka-consumer-groups.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --group test-group --describe

# consume messages from a consumer group
./kafka-console-consumer.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --topic test-topic --group test-group

# describe as consumer group
./kafka-consumer-groups.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --group test-group --describe

# list the consumer groups
./kafka-consumer-groups.sh --bootstrap-server broker-1:19092,broker-2:19092,broker-3:19092 --list