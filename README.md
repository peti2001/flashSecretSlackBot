# Flash Secret SlackBot

Sometimes you might want to send secret to your colleague and you want to be sure the secret is not compromised and you
don't want to leave behind sensitive data.

For this purpose the services like https://quickforget.com/ are excellent.

This little Go code will help you to generate flash secret storage on https://quickforget.com/. After the first read or
after 24 hours your secret will be erased.

## How to use it?
```
git clone git@github.com:peti2001/flashSecretSlackBot.git
cd flashSecretSlackBot
go build
```

Deploy the executable to your server and create a slash command in slack. 

## CLI parameters
* `-h`: Show help
* `-port`: The port where the service will be available
* `-path`: The URL where the service will be availabble
* `-slack-token`: A secret token from Slack. Leave it empty to skip the check.

## Contribute
Feel free to implement other flash secret storage other than https://quickforget.com/
