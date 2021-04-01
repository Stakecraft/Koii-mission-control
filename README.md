# solana-prometheus

# Get the code
```bash
$ git clone https://github.com/PrathyushaLakkireddy/solana-prometheus
$ cd solana-prometheus
$ cp example.config.toml config.toml
```

# Configure the following variables in `config.toml`
- **[telegram]**
  - *tg_chat_id*

    Telegram chat ID to receive Telegram alerts, required for Telegram alerting.
    
  - *tg_bot_token*

    Telegram bot token, required for Telegram alerting. The bot should be added to the chat and should have send message permission.
- **[Email]**

  - *email_address*

    E-mail address to receive mail notifications, required for e-mail alerting.
- **[validator_details]**

   - *validator_name*
   
       Moniker of your validator,to get it displayed in alerts.

   - *pub_key*
  
      Public Node key of the validator, which will be used to get validator identity and other validator metrics like commision,validator status etc...

   - *vote_key*
   
      Vote key of validator, required for validator Identity.
    
- **[enable_alerts]**

   - *enable_telegram_alerts*

      Configure **yes** if you wish to get telegram alerts otherwise make it **no** .

   - *enable_email_alerts*

      Configure **yes** if you wish to get email alerts otherwise make it **no** .

- **[regular_status_alerts]**
   - *alert_timings*
   
      Array of timestamps for alerting about the validator health, i.e. whether it's voting or jailed. You can get alerts based on the time which can be configured.

- **[alerter_preferences]**

   - *account_balance_change_alerts*

       Configure **yes** if you wish to get account balance change alerts otherwise make it **no** .

   - *block_diff_alerts*

       If you want to recieve alerts when there is a gap between your validator block height and network height then make it **yes** otherwise **no**

   - *epoch_diff_alrets*

      If you want to recieve alerts when there is a gap between your validator epoch and network epoch then make it **yes** otherwise **no**

   - *delegation_alerts*

      Configure **yes** if you wish to get alerts about alters in account balance otherwise make it **no**

- **[alerting_threholds]**

   - *block_diff_threshold*

      An Integer value to recieve block difference alerts, e.g. a value of 2 would alert you if your validator falls 2 or more blocks behind the network's current block height.

    - *epoch_diff_threshold*
       
       An integer value to recieve epoch difference alerts, e.g. a value of 5 would alert you if your validator's epoch number and network's epoch difference is 5 or more.

    - *account_bal_threshold*

       An integer value to recieve account balance change alerts, e.g. if your account balance has dropped to given threshold value you will receive alerts.

- **[prometheus]**

    - *prometheus_address*
       
       Address of your lcd client (ex: http://localhost:9090)

    - *listen_address*
       
       Port in which prometheus will run, by default it will run on 9090 port

After populating config.toml, check if you are running prometheus server in local

## Grafana Dashboards

Solana MOnitoring Tool provides three dashboards

1. Validator Monitoring Metrics (These are the metrics which we have calculated and stored in prometheus)
2. System Metrics (These are the metrics related to the system configuration)
3. Summary (Which gives quick overview of validator and system metrics)

### 1. Validator monitoring metricss

The following list of metrics are displayed in this dashboard.

- **Validator Details**

  Validator public Key:
    Node key of validator

  Validator vote key:
     Vote Key of the validator

- **Validator Information**

    Solana node version:
       Current version of the solana

    Solana node health:
        current health of the node

    IP Address:
        Gossip address of node
    
    Vote Account: 
       Shows information about whether the validator is voting or jailed
    
    Validator Active stake:
        Displays Validator current active stake
        
    Commision:
      Validator's vote account commision

- **validator Health**

    Network Epoch:
       Displays network's epoch height 

    Validator Epoch:
       Displays validator' epoch height

    Epoch Difference:
       Difference between validator's and network's epoch

     Validator status:
        Shows validator's status whether the validator is voting or jailed

- **Validator Performance**

    Block Height - Network: 
       Displays the latest block height of a network
    
    Block Height - Validator:
       Displays the latest block height committed by the validator
    
    Height Difference:
       Displays Block Height Difference of network and validator's block height

    Solana Blocktime:
       Displays estimated production time of a confirmed block
    
    Vote Height - Network:
       Displays the latest vote height of network
    
    Vote Height - Validator:
       Displays the latest vote height committed by validator

    Vote Height Difference:
       Displays height difference of network and validator's vote height

    Account Balance:
       Displays Account Balance of validator in SOL's
    
    Solana current slot height:
        Displays Current slot height 
    
    Confirmed Blocktime - Network:
         Displays estimated production time of confirmed block of Network
    
    Confirmed Blocktime - Validator:
         Displays estimated production time of confirmed block committed by validator
        
    Block Time Difference:
        Difference between confirmation time of network and validator

    Current Epoch - vote credits:
        Displays current epoch vote credits of validator vote account

    Previous Epoch - Vote credits:
        Displays previous epoch vote credits of validator vote account
    
    Total valid slots:
       Displays number of leader valid slots per leader
    
    Total skipped slots:
       Displays number of leader skipped slots per leader

- **Validator Details**

   solana slot leader:
      Displays current slot leader
   
   Transaction Count:
      Displays current Transaction count from the ledger
   
   Validator last voted:
      Displays Most recent slot voted on by this vote account
   
   Solana confirmed slot height:
      Displays current slot height
    
   Current Active Validators:
       Displays the number of current active validators, i.e validators who are voting 
    
   Delinquent validators:
      Displays the number of delinquent validators, i.e validators who are jailed.

   Confirmed epoch last slot:
      Displays current epoch's last slot

   Validator Root slot:
      Displays Root slot per validator
   
    




      
