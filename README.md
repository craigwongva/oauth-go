### Define your application to GitHub
1\. Sign in to GitHub and go to Settings, Oauth Applications.

2\. For "Authorization Callback URL", specify "http://12.13.14.15:8888/callmebacktomorrow" (this must match oauthproxy.go line 645 as described below).

### Run your application (a trivial Go web server)

3\. From ssh session one, run the trivial Go server. It's listening on :8888.
```
clientid=<your GitHub app client id>
clientsecret=<your GitHub app client secret>
cd ~/oauth-go
./bin/jquery $clientid $clientsecret
```

### Run oauth2_proxy

4\. Get oauth2_proxy
```
cd ~/oauth-go
go get github.com/bitly/oauth2_proxy
```
5\. Make sure the oauth2-proxy will forward the correct callback_uri.
  1. Force oauthproxy.go line 645 to "http://12.13.14.15:8888/callmebacktomorrow" (TODO: This manual override is necessary, but why?)
  1. go install github.com/bitly/oauth2_proxy
  1. Update the GitHub settings to that same string. This allows the browser access to a further GitHub authorization you-choose-‘em page.

6\. From ssh session two, run oauth2-proxy. It's listening on :4180.
```
cd ~/oauth-go
clientid=<your GitHub app client id>
clientsecret=<your GitHub app client secret>
./cli-oauth2_proxy $clientid $clientsecret
```

### Run test to see Github authentication in action
7\. From a browser, start the authentication sequence.
```
http://12.13.14.15:8887
```

### What's happening
nginx is listening on :8887, authenticates (TODO: authentication is redundant, actually superfluous, so remove this), forwards to oauth2_proxy, which forwards to GitHub, which returns a code to your app via the callback URL, your app sends the code back to GitHub and receives an access_token (TODO: where?), your app uses the access_token to make a GitHub API call.

If your app was only using GitHub for authentication and cared nothing about GitHub itself, then your app would then manage your application's authorizations (TODO: find AWS federation login S3 example).

### Bibliography
* http://developers.canal-plus.com/blog/install-nginx-reverse-proxy-with-github-oauth2/
* https://github.com/bitly/oauth2_proxy
* https://developer.github.com/v3/oauth/
* https://golang.org/pkg/net/http/
* https://golang.org/pkg/net/url/
* http://stackoverflow.com/questions/35663892/parse-string-into-map-golang
* https://gobyexample.com/command-line-arguments

