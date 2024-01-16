### A - Install Grafana for Ubuntu
Download the latest .deb file and extract it by using the following commands

```sh
$ cd $HOME
$ sudo apt-get install -y adduser libfontconfig1
$ wget https://dl.grafana.com/oss/release/grafana_7.5.2_amd64.deb
$ sudo dpkg -i grafana_7.5.2_amd64.deb
```

Start the grafana server
```sh
$ sudo -S systemctl daemon-reload

$ sudo -S systemctl start grafana-server

Grafana will be running on port :3000 (ex:: https://localhost:3000)
```

### Install Prometheus

```sh
$ cd $HOME

$ wget https://github.com/prometheus/prometheus/releases/download/v2.22.1/prometheus-2.22.1.linux-amd64.tar.gz

$ tar -xvf prometheus-2.22.1.linux-amd64.tar.gz

$ sudo cp prometheus-2.22.1.linux-amd64/prometheus $GOBIN

$ sudo cp prometheus-2.22.1.linux-amd64/prometheus.yml $HOME
```
- Add the following in prometheus.yml using your editor of choices

```sh

  - job_name: 'koii'

    static_configs:
      - targets: ['localhost:5678']
    
```

Setup Prometheus System service

```bash
sudo nano /lib/systemd/system/prometheus.service
```
- Copy-paste the following:
   
```sh
[Unit]
Description=Prometheus
After=network-online.target

[Service]
Type=simple
ExecStart=/home/ubuntu/go/bin/prometheus --config.file=/home/ubuntu/prometheus.yml
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
- For the purpose of this guide it is assumed the `user` is `ubuntu`. If your user is   different please make the required changes above.
     
```sh 
$ sudo systemctl daemon-reload
$ sudo systemctl enable prometheus.service
$ sudo systemctl start prometheus.service
```

### Install node exporter

```sh
$ cd $HOME
$ curl -LO https://github.com/prometheus/node_exporter/releases/download/v1.2.2/node_exporter-1.2.2.linux-amd64.tar.gz
$ tar -xvf node_exporter-1.2.2.linux-amd64.tar.gz
$ sudo cp node_exporter-1.2.2.linux-amd64/node_exporter $GOBIN
```

Setup Node exporter service

```bash 
 sudo nano /lib/systemd/system/node_exporter.service
 ```

 Copy-paste the following:

 ```sh
 [Unit]
Description=Node_exporter
After=network-online.target

[Service]
Type=simple
ExecStart=/home/ubuntu/go/bin/node_exporter
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
For the purpose of this guide it is assumed the `user` is `ubuntu`. If your user is different please make the required changes above.

- **Note**:  Do not forget to setup node exporter configuration in prometheus.yml file.

Copy paste the following in prometheus.yml.

```sh
  - job_name: 'node_exporter'

      static_configs:
      - targets: ['localhost:9100']
```
Start Node exporter

```bash
$ sudo systemctl daemon-reload
$ sudo systemctl enable node_exporter.service
$ sudo systemctl start node_exporter.service
```

### Install koii client (If not already installed)
- Koii client binary can be used to get metrics of skip rate and block production details.
- So make sure to configure it before running monitoring tool.

- If you don't have prebuilt binaries of koii then you can follow this documentation.Follow this doc (https://docs.koii.network/run-a-node/k2-validators/system-setup)
- You can install this binary on your non validator node also.

- If you want to specify the koii binary path, then execute this command by providing koii executable path `export KOII_BINARY_PATH="path_to_koii_binary"` (ex : export KOII_BINARY_PATH="/usr/local/bin/koii")
- This provided path will be used in monitoring tool code and fetch the metrics accordingly.

Note : This is important to get metrics related to skip rate and block production details(leader slots, blocks produced etc). So please make sure to have, installed koii client binary.
