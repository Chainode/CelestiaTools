# Enable tendermint metrics

To enable the export of the Tendermint metrics, you need to set the flag for Prometheus to *true* in the *[instrumentation]* section inside of your *~/.celestia-app/config/config.toml* file  
[instrumentation]  
prometheus = true  
prometheus_listen_addr = ":26660"

![image](https://user-images.githubusercontent.com/53407923/136295121-f6eaec5e-76ec-4333-8893-d265846f9cec.png)

