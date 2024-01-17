# USE THIS INSTRUCTION ASSUMING THAT YOU ALREADY HAVE GRAFANA AND PROMETHEUS INSTALLED 

## 1 - Install/Update GO

```sh
$ wget -c https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
$ sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
$ export PATH=$PATH:/usr/local/go/bin
$ go version

```

## 2 - Clone git and preconfigure your tool

```sh
$ git clone https://github.com/Stakecraft/koii-mission-control
$ cd koii-mission-control
$ cp example.config.toml config.toml
$ sudo nano config.toml

```
Note: you need to add your validator name, identity and vote account in the config.toml
In case prometheus in running on a different machine - **Do not forget to add your prometheus IP as well**. Instead of 
```sh
prometheus_address = "http://localhost:9090"
```
use
```sh
prometheus_address = "http://[YOUR_Prometheus_IP]:9090"
```
Edit the `config.toml` with your changes. Information about all the fields in `config.toml` can be found [here](./docs/config-desc.md)

## 3 - Allow connection from Prometheus and build your tool

```sh
$ export SOLANA_BINARY_PATH="/home/koii/.local/share/koii/install/active_release/bin/koii"
$ sudo ufw allow from YOUR_Prometheus_IP to any port 5678 comment "Koii Monitoring"
$ go build -o koii-mc
```

## 4 - Create system file and run it it

```sh
$ nano /etc/systemd/system/koii.service

[Unit]
Description=Koii-mc
After=network-online.target

[Service]
User=koii
Environment=KOII_BINARY_PATH="/home/koii/.local/share/koii/install/active_release/bin/koii"
WorkingDirectory=/home/koii/koii-mission-control/
ExecStart=/home/koii/koii-mission-control/koii-mc
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target

```
```sh
$ sudo systemctl daemon-reload
$ sudo systemctl restart koii_mc.service && sudo systemctl enable koii_mc.service && sudo journalctl -u koii_mc.service -f
```

## 5 - **Prometheus configuration**

you need to add the following configuration to `prometheus.yml`. You can find the prometheus file at `$HOME/prometheus.yml` .

```sh
 scrape_configs:

  - job_name: 'koii'

    static_configs:
    - targets: ['localhost:5678']

```

Restart the prometheus serivce

```sh 
$ sudo systemctl daemon-reload
$ sudo systemctl restart prometheus.service
```

## 6 - **Grafana Dashboards**

### Import the dashboards

- To import the dashboards click the **+** button present on left hand side of the dashboard. Click on import and paste the UID of the dashboards on the text field below **Import via grafana.com** and click on load. 

- Select the datasources and click on import.

UID of dashboards are as follows:

 - **14738**: Validator monitoring metrics dashboard.
 - **14739**: Summary dashboard.
 - **13445**: System monitoring metrics dashboard.

 While importing these dashboards if you face any issues at valueset i.e., unique identifier (uid) of a dashboard, then change it to empty and then click on import by selecting the datasources.


- *For more info about grafana dashboard imports you can refer https://grafana.com/docs/grafana/latest/reference/export_import/*
