# Rabbit Tools

The object of this tools is to help a manager in your RabbitMQ blazingly fast

## Set config.env

In the same local of download, create the file `config.env` and fill with

```
RABBIT_URL=http://your_url/api/queues/your_virtual_host
RABBIT_USER=your_user
RABBIT_PASS=your_pass
```

Check if the plugin of rabbitmq-manager is available in your RabbitMQ

[Management Plugin](https://www.rabbitmq.com/management.html)

## How to use

Rabbit Tools for now only has two functions show List of Queue or Purge Queue

### List

```
./rabbit-tools list all
```

Result:
+--------------------------------+----------+
| QUEUE                          | MESSAGES |
+--------------------------------+----------+
| something_test1                |        5 |
| queue_test1234                 |        5 |
| queue_test1                    |        5 |
| test_something                 |        3 |
| something_queue                |        1 |
+--------------------------------+----------+
| TOTAL QUEUES                   |        5 |
| TOTAL MESSAGES                 |        19|
+--------------------------------+----------+

List all Queue

```
./rabbit-tools list something
```

Result:

+--------------------------------+----------+
| QUEUE                          | MESSAGES |
+--------------------------------+----------+
| something_test1                |        5 |
| test_something                 |        3 |
| something_queue                |        1 |
+--------------------------------+----------+
| TOTAL QUEUES                   |        3 |
| TOTAL MESSAGES                 |        9 |
+--------------------------------+----------+


List queues with **something** in the name

### Purge

```
./rabbit-tools purge all
```
Result:

+--------------------------------+----------+
| PURGE QUEUE                    | MESSAGES |
+--------------------------------+----------+
| something_test1                |        5 |
| queue_test1234                 |        5 |
| queue_test1                    |        5 |
| test_something                 |        3 |
| something_queue                |        1 |
+--------------------------------+----------+
| TOTAL PURGED MESSAGES          |        19|
+--------------------------------+----------+

Purge content all Queue

```
./rabbit-tools purge something
```
Result:

+--------------------------------+----------+
| PURGE QUEUE                    | MESSAGES |
+--------------------------------+----------+
| something_test1                |        5 |
| test_something                 |        3 |
| something_queue                |        1 |
+--------------------------------+----------+
| TOTAL PURGED MESSAGES          |        9 |
+--------------------------------+----------+

Purge content of all queues that has **something** in the name
