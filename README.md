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
rabbit list all
```

List all Queue

```
rabbit list something
```

List queues with **something** in the name

### Purge

```
rabbit purge all
```

Purge content all Queue

```
rabbit purge something
```

Purge content of all queues that has **something** in the name
