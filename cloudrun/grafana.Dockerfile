FROM grafana/grafana:12.0.2

COPY cloudrun/grafana-datasources.yml /etc/grafana/provisioning/datasources/datasources.yml
COPY observability/grafana/provisioning/dashboards/dashboards.yml /etc/grafana/provisioning/dashboards/dashboards.yml
COPY observability/grafana/dashboards /var/lib/grafana/dashboards
