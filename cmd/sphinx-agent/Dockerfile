FROM jrei/systemd-centos:7

COPY SphinxAgent.viper.yaml /etc/SphinxAgent/SphinxAgent.viper.yaml
COPY sphinx-agent /usr/local/bin/sphinx-agent

CMD ["sphinx-agent", "run"]
