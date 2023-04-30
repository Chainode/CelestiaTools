# Expose Celestia Bridge Node Metrics  
 
### Download the compiled go binary celbridge_export
``` 
wget https://github.com/Chainode/CelestiaTools/blob/main/celbridge_export
```
In this repo you can also find the code of the compiled binary, in the file celbridge_export.go

The binary has the following flags:
``` 
--listen.port 8380 - with this you can specify the listen port and is relevant for the prometheus configuration to scrap the metrics. The used port 8380 is an example and if no port is specified, it will default to this value.
--endpoint http://localhost:26658 - with this flag you can specfiy to which bridge rpc address it should connect to. The used endpoint http://localhost:26658 is an example and if no endpoint is specified, it will default to this value.
--p2p.network blockspacerace  - with this flag you define the p2p network the bridge node is active on. The used p2p network blockspacerace is an example and if no p2p network is specified, it will default to this value.
```

### Create systemd file  
``` 
sudo nano /etc/systemd/system/celbridge_exporter.service  
```

### You can use the file below, adapt it to your user and group or create your own customized file based on this file and your installation
```   
[Unit]  
Description=Celestia Bridge Exporter  
After=network.target  
  
[Service]  
User=<your-user>  
Group=<your-user>  
Type=simple  
ExecStart=/home/<your-user>/celbridge_export --listen.port 8380 --endpoint http://localhost:26658 --p2p.network blockspacerace
  
[Install]  
WantedBy=multi-user.target  
```

### Start node exporter  
``` 
sudo systemctl start celbridge_exporter  
```
### Enable node exporter at startup  
``` 
sudo systemctl enable celbridge_exporter   
```
### Check the status of node exporter  
```
sudo systemctl status celbridge_exporter  
```
