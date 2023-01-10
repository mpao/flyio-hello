# Hello World on [fly.io](https://fly.io)


* [Howto](https://fly.io/docs/hands-on/)
* [Go example](https://fly.io/docs/languages-and-frameworks/golang/)

The `Dockerfile` is useless, use it just for local deployment, fly command take care of all the stuff.<br/> 
Here's the *TL:DR;* version to deploy a GO application.

```bash
$> CGO_ENABLED=0 go build -o hello main.go
$> flyctl launch
```
