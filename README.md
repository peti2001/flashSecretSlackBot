# Flash Secret SlackBot

Sometimes you might want to send secret to your colleague and you want to be sure the secret is not compromised and you
don't want to leave behind sensitive data.

For this purpuse the services like https://quickforget.com/ are excelent. 

This little Go code will help you to generate flash secret storage on https://quickforget.com/. After the first read or
after 24 hours your secret will be ereased.

#How to use it?
`cd flashSecretSlackBot`
`go build`

Deploy the executable to your server and create a slash command in slack. 

#Configuration
CLI parameters:
-port: The port where the service will be available
-path: The URL where the service will be availabble
-slack-token: A secret token from Slack. Leave it empty to skip the check.
