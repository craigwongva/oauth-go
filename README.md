## Run the trivial Go web server

1. From ssh session one, run the trivial Go server. It's listening on :8888.
```
clientid=<your GitHub app client id>
clientsecret=<your GitHub app client secret>
cd ~/oauth-go
./bin/jquery $clientid $clientsecret
```

## Run oauth2_proxy

2\. Get oauth2_proxy
```
cd ~/oauth-go
go get github.com/bitly/oauth2_proxy
```
3\. Make sure the oauth2-proxy will forward the correct callback_uri.
  1. Force oauthproxy.go line 645 to "http://12.13.14.15:8888/callmebacktomorrow" 
  1. go install github.com/bitly/oauth2_proxy
  1. Update the GitHub settings to that same string. This allows the browser access to a further GitHub authorization you-choose-‘em page.

4\. From ssh session two, run oauth2-proxy. It's listening on :8887.
```
cd ~/oauth-go
clientid=<your GitHub app client id>
clientsecret=<your GitHub app client secret>
./cli-oauth2_proxy $clientid $clientsecret
```

## Run test to see Github authentication in action
5. From a browser, start the authentication sequence.
```
http://12.13.14.15:8887
```

## Bibliography
* http://developers.canal-plus.com/blog/install-nginx-reverse-proxy-with-github-oauth2/
* https://github.com/bitly/oauth2_proxy
* https://developer.github.com/v3/oauth/
* https://golang.org/pkg/net/http/
* https://golang.org/pkg/net/url/
* http://stackoverflow.com/questions/35663892/parse-string-into-map-golang
* https://gobyexample.com/command-line-arguments

