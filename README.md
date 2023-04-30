# CelestiaTools
## Monitoring Celestia Validator and Bridge Status by using Grafana & Prometheus


## Scope

With this solution it will be possible to monitor both your Validator, the Celestia bridge and the hardware your validator and bridge runs on. This will help you improve the overall health of your Validator as well as offer you proper insight into what happens in the Celestia Network and with your machine at any point in time.  
Prometheus represents a monitoring solution for storing time series data like metrics. Grafana is another important complementary tool which allows users to visualize the data stored in Prometheus and many more sources like Telegraf, etc.  

In combination with an Alertmanager it becomes a viable solution to monitor and improve the overall uptime and performance of your Validator. Alertmanager can be used to define certain thresholds for which you will get alerted on Telegram, Discord, by SMS or many other mediums. Some example for possbile alerts: You can set your threshold for disk usage rate to 70% so that you will get alerted once that is reached in order to prepare for an upgrade for more disk space; low number of connected peers; validator is stuck (block height is not increasing over a certain period of time); validator is down; and much more.  

## Prerequisites
* Enable export of tendermint metrics --> please check ExposeTendermintMetrics in this repo: https://github.com/Chainode/CelestiaTools/blob/main/ExposeTendermintMetrics.md  
* Node Exporter  --> please check guide InstallNodeExporter in this repo: https://github.com/Chainode/CelestiaTools/blob/main/InstallNodeExporter.md
* Celestia Bridge Exporter --> please check guide ExposeCelestiaBridgeMetrics in this repo: https://github.com/Chainode/CelestiaTools/blob/main/ExposeCelestiaBridgeMetrics.md
* Prometheus  --> please check InstallPrometheus in this repo: https://github.com/Chainode/CelestiaTools/blob/main/InstallPrometheus.md  
* Grafana  --> please check guide InstallGrafana in this repo: https://github.com/Chainode/CelestiaTools/blob/main/InstallGrafana.md
* Grafana Pie Chart plugin  --> please check guide InstallGrafana in this repo: https://github.com/Chainode/CelestiaTools/blob/main/InstallGrafana.md

## Import Celestia Validator Dashboard into Grafana  
For this you will have to download the Celestia Validator Dashboard.json from this repo and then in Grafana go to *Dashboards* -> *Manage* -> *Import* -> *Upload JSON file* and select the Celestia Validator Dashboard.json to be uploaded.  
During the import, the Celestia Validator Dashboard.json Grafana dashboard will automatically search for Prometheus datasources that contain "Celestia", "celestia" or "cel" in their name.  
This can be changed once the dashboard was imported by going to "Dashboard Settings" -> "Variables" -> "datasource" variable -> here you will see a parameter "Instance name filter" where the regex condition for selecting the Prometheus source was defined.
As a next step you should make sure the right Prometheus connection has been selected and if necessary adapt the Dashboard. 


## Final Result 
![image](https://user-images.githubusercontent.com/53407923/235368283-385472a6-3ca1-45d0-a595-85dae4660344.png)
![image](https://user-images.githubusercontent.com/53407923/235368336-e7daf834-54fb-4e6b-b8c6-4771cb8c032f.png)


