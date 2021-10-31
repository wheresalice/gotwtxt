# GoTwtxt

A Go implementation of [twtxt](https://github.com/buckket/twtxt) - decentralised, minimalist microblogging service for hackers

## Usage:

Send a tweet:

```shell
gotwtxt tweet hello world
```

Follow somebody:

```shell
gotwtxt follow wheresalice https://envs.net/~wheresalice/twtxt.txt
```

Unfollow somebody:

```shell
gotwtxt unfollow wheresalice
```

See who you are following:

```shell
gotwtxt following
```

Fetch the timeline:

```shell
gotwtxt timeline
```

## Publishing

As with all good unix commands, it's up to you how you want to publish your file

Personally I run this on a shared unix system with a webserver running, but you could also write an alias which runs the tweet command followed by running scp or whatever else you'd like