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

### Flags


#### Env

```
--env=Prefix
```

You pass in flag env the prefix to use more than one configuration. Example:

```
./rabbit-tools list all --env=PROD
```

The configuration look like:

```
PROD_RABBIT_URL=http://your_url/api/queues/your_virtual_host
PROD_RABBIT_USER=your_user
PROD_RABBIT_PASS=your_pass
```


### List

```
./rabbit-tools list all
```

#### List *all* Queue

Result:
```
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
```

#### List by target

```
./rabbit-tools list something
```

Search and list all queues with **something** in the name

Result:
```
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
```


### Purge

#### Purge content of all Queue

**all** is a reserved word

```
./rabbit-tools purge all
```

Result:
```
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
```

#### Purge by target
```
./rabbit-tools purge something
```

Search and purge content of all queues that has **something** in the name

Result:
```
+--------------------------------+----------+
| PURGE QUEUE                    | MESSAGES |
+--------------------------------+----------+
| something_test1                |        5 |
| test_something                 |        3 |
| something_queue                |        1 |
+--------------------------------+----------+
| TOTAL PURGED MESSAGES          |        9 |
+--------------------------------+----------+
```
