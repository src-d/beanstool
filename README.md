beanstool [![Latest Stable Version](http://img.shields.io/github/release/tyba/beanstool.svg?style=flat)](https://github.com/tyba/beanstool/releases)
==============================

Dependency free [beanstalkd](http://kr.github.io/beanstalkd/) admin tool.

Basically this is a rework of the wonderful [bsTools](https://github.com/jimbojsb/bstools) with some extra features and of course without need to install any dependency. Very useful in companion of the server in a small docker container.

Installation
------------

```
wget https://github.com/tyba/beanstool/releases/download/v0.1.0/beanstool_v0.1.0_linux_amd64.tar.gz
tar -xvzf beanstool_v0.1.0_linux_amd64.tar.gz
cp beanstool_v0.1.0_linux_amd6/beanstool /usr/local/bin/
```

browse the [`releases`](https://github.com/tyba/beanstool/releases) section to see other archs and versions


Usage
-----

```sh
Usage:
  beanstool [OPTIONS] <command>

Help Options:
  -h, --help  Show this help message

Available commands:
  bury   bury existing jobs from ready state
  kick   kicks jobs from buried back into ready
  peek   peeks a job from a queue
  put    put a job into a tube
  stats  print stats on all tubes
  tail   tails a tube and prints his content
```

As example this is the output of the command `./beanstool stats`:

```
+---------+----------+----------+----------+----------+----------+---------+-------+
| Name    | Buried   | Delayed  | Ready    | Reserved | Urgent   | Waiting | Total |
+---------+----------+----------+----------+----------+----------+---------+-------+
| foo     | 20       | 0        | 5        | 0        | 0        | 0       | 28    |
| default | 0        | 0        | 0        | 0        | 0        | 0       | 0     |
+---------+----------+----------+----------+----------+----------+---------+-------+
```

License
-------

MIT, see [LICENSE](LICENSE)
