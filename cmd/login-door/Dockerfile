FROM jrei/systemd-centos:7

COPY LoginDoor.viper.yaml /etc/LoginDoor/LoginDoor.viper.yaml
COPY login-door /usr/local/bin/login-door

CMD ["login-door", "run"]
