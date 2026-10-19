FROM otel/opentelemetry-collector-contrib:0.123.0

COPY cloudrun/otel-collector-config.yml /etc/otelcol-contrib/config.yaml

ENTRYPOINT ["/otelcol-contrib"]
CMD ["--config=/etc/otelcol-contrib/config.yaml"]
