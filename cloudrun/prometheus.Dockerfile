FROM prom/prometheus:v3.4.2

COPY cloudrun/prometheus.yml /etc/prometheus/prometheus.yml

ENTRYPOINT ["/bin/prometheus"]
CMD ["--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/tmp/prometheus"]
