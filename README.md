beanstool [![Latest Stable Version](http://img.shields.io/github/release/tyba/beanstool.svg?style=flat)](https://github.com/tyba/beanstool/releases) [![Circle CI](https://img.shields.io/circleci/project/tyba/beanstool.svg?style=flat)](https://circleci.com/gh/tyba/beanstool)
==============================

Dependency free [beanstalkd](http://kr.github.io/beanstalkd/) admin tool.

Basically this is a rework of the wonderful [bsTools](https://github.com/jimbojsb/bstools) with some extra features and of course without need to install any dependency. Very useful in companion of the server in a small docker container.

##Installation
###Linux
```
wget https://github.com/tyba/beanstool/releases/download/v0.2.0/beanstool_v0.2.0_linux_amd64.tar.gz
tar -xvzf beanstool_v0.2.0_linux_amd64.tar.gz
cp beanstool_v0.2.0_linux_amd64/beanstool /usr/local/bin/
```
###Mac OS X
```
wget https://github.com/tyba/beanstool/releases/download/v0.2.0/beanstool_v0.2.0_darwin_amd64.tar.gz
tar -xvzf beanstool_v0.2.0_darwin_amd64.tar.gz
cp beanstool_v0.2.0_darwin_amd64/beanstool /usr/local/bin/
```


Browse the [`releases`](https://github.com/tyba/beanstool/releases) section to see other archives and versions.


##Usage

```sh
Usage:
  beanstool [OPTIONS] <command>

Help Options:
  -h, --help  Show this help message

Available commands:
  bury   bury existing jobs from ready state
  kick   kicks jobs from buried back into ready
  delete a job from a queue
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

##License

MIT, see [LICENSE](LICENSE)
