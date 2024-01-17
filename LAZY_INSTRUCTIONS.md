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
$ prometheus_address = "http://localhost:9090"
```
use
```sh
prometheus_address = "http://[YOUR_Prometheus_IP]:9090"
```


## 3 - Allow connection from Prometheus and build your tool

```sh
export SOLANA_BINARY_PATH="/home/koii/.local/share/koii/install/active_release/bin/koii"
sudo ufw allow from YOUR_Prometheus_IP to any port 5678 comment "Solana Monitoring"
go build -o koii-mc
```

## 4 - Create system file and run it it

```sh
nano /etc/systemd/system/koii.service

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
sudo systemctl daemon-reload
sudo systemctl restart koii_mc.service && sudo systemctl enable koii_mc.service && sudo journalctl -u koii_mc.service -f
```

- **Koii Client Binary**
- **Go 1.14.x+**
- **Grafana 7.x+**
- **Prometheus**
- **Node Exporter**

### Prerequisite Installation

 - Koii Client Binary Installation 

   Before installing prerequisites make sure to have koii client binary installed.
   - If you haven't installed it before, follow [this guide](https://docs.koii.network/run-a-node/k2-validators/system-setup) to install the prebuilt binaries of latest version.

   To learn more about koii client binary usage [click here](https://github.com/Stakecraft/koii-mission-control/blob/main/docs/prereq-manual.md#install-solana-client).

 - Install other prerequisites


**2) --> Manual installation**

To manually install the prerequisites please follow this [guide](./docs/prereq-manual.md).
 
## Install and configure the Koii Monitoring Tool

Manual installation

```bash
$ git clone https://github.com/Stakecraft/koii-mission-control
$ cd koii-mission-control
$ cp example.config.toml config.toml
```

**Note** : (OPTIONAL) If you wish to pass your config path from an ENV variable then you can use this command. `export CONFIG_PATH="/path/to/config"` (ex: `export CONFIG_PATH="/home/Desktop"`).

Edit the `config.toml` with your changes. Information about all the fields in `config.toml` can be found [here](./docs/config-desc.md)

Note : Before running this monitoring binary, you need to add the following configuration to `prometheus.yml`. You can find the prometheus file at `$HOME/prometheus.yml` .

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

If you wish to pass `koii binary path` then you export by following below step.
```sh
export KOII_BINARY_PATH="<koii-client-binary-path>" # Ex - export KOII_BINARY_PATH="/home/ubuntu/.local/share/koii/install/active_release/bin:$PATH"
```

- Build and run the monitoring binary

```sh
   $ go build -o koii-mc && ./koii-mc
```

- Run monitoring tool as a system service

Follow below steps to create a system service file and to start it.
Before running this make sure to export the `$KOII_BINARY_PATH`.

```sh
echo "[Unit]
Description=koii-mc
After=network-online.target

[Service]
User=$USER
Environment="KOII_BINARY_PATH=$KOII_BINARY_PATH"
ExecStart=$HOME/go/bin/koii-mc
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target" | sudo tee "/lib/systemd/system/koii_mc.service"
```
- Run the system service file
```sh
sudo systemctl daemon-reload

sudo systemctl enable koii_mc.service

sudo systemctl start koii_mc.service
````

Installation of the tool is completed let's configure the grafana dashboards.

### Grafana Dashboards

The repo provides three dashboards

1. Validator Monitoring Metrics - Displays the validator metrics which are calculated and stored in prometheus.
2. System Monitoring Metrics - Displays the metrics related to your validator server on which this tool is hosted on.
3. Summary - Displays a quick overview of validator monitoring metrics and system metrics.

Information of all the dashboards can be found [here](./docs/dashboard-desc.md).

## How to import these dashboards in your Grafana installation

### 1. Login to your Grafana dashboard
- Open your web browser and go to http://<your_ip>:3000/. `3000` is the default HTTP port that Grafana listens to, if you haven’t configured a different port.
- If you are a first time user type `admin` for the username and password in the login page.
- You can change the password after login.

### 2. Create Datasource

- Before importing the dashboards you have to create a `Prometheus` datasources.

- To create the datasoruce go to **Configuration** and select **Data Sources**.

- Click on **Add data source** and select `Prometheus` from Time series databases section.

- Replace the URL with http://localhost:9090. 

- Click on **Save & Test** .

### 3. Import the dashboards

- To import the dashboards click the **+** button present on left hand side of the dashboard. Click on import and paste the UID of the dashboards on the text field below **Import via grafana.com** and click on load. 

- Select the datasources and click on import.

UID of dashboards are as follows:

 - **14738**: Validator monitoring metrics dashboard.
 - **14739**: Summary dashboard.
 - **13445**: System monitoring metrics dashboard.

 While importing these dashboards if you face any issues at valueset i.e., unique identifier (uid) of a dashboard, then change it to empty and then click on import by selecting the datasources.


- *For more info about grafana dashboard imports you can refer https://grafana.com/docs/grafana/latest/reference/export_import/*
