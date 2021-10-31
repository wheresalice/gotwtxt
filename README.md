# GoTwtxt
> A Go implementation of [twtxt](https://github.com/buckket/twtxt) - decentralised, minimalist microblogging service for hackers

[![Release](https://img.shields.io/github/release-pre/wheresalice/gotwtxt.svg?logo=github&style=flat&v=1)](https://github.com/wheresalice/gotwtxt/releases)
[![Build Status](https://img.shields.io/github/workflow/status/wheresalice/gotwtxt/run-go-tests?logo=github&v=1)](https://github.com/wheresalice/gotwtxt/actions)
[![Mergify Status](https://img.shields.io/endpoint.svg?url=https://gh.mergify.io/badges/wheresalice/gotwtxt&style=flat&v=1)](https://mergify.io)
[![Go](https://img.shields.io/github/go-mod/go-version/wheresalice/gotwtxt?v=1)](https://golang.org/)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/wheresalice/gotwtxt)


## Usage:

Send a tweet:

```shell
# pass the message on the cli
gotwtxt tweet hello world

# or take the first line from stdin
echo hello world | gotwtxt tweet
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



## License

[![License](https://img.shields.io/github/license/wheresalice/gotwtxt.svg?style=flat&v=1)](LICENSE)
