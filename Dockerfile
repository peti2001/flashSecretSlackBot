FROM peti2001/go-base-with-certificates

ADD ./flashSecretSlackBot /opt/slackbot/

CMD ["/opt/slackbot/flashSecretSlackBot"]