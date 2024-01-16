#!/bin/bash

set -e

cd $HOME

echo "--------- Cloning koii-monitoring-tool -----------"

git clone https://github.com/Stakecraft/koii-mission-control.git

cd koii-mission-control

mkdir -p  ~/.koii-mc/config/

cp example.config.toml ~/.koii-mc/config/config.toml

cd $HOME

echo "------ Updatig config fields with exported values -------"

sed -i '/rpc_endpoint =/c\rpc_endpoint = "'"$RPC_ENDPOINT"'"' ~/.koii-mc/config/config.toml

sed -i '/network_rpc =/c\network_rpc = "'"$NETWORK_RPC"'"' ~/.koii-mc/config/config.toml

sed -i '/validator_name =/c\validator_name = "'"$VALIDATOR_NAME"'"'  ~/.koii-mc/config/config.toml

sed -i '/pub_key =/c\pub_key = "'"$PUB_KEY"'"'  ~/.koii-mc/config/config.toml

sed -i '/vote_key =/c\vote_key = "'"$VOTE_KEY"'"'  ~/.koii-mc/config/config.toml

if [ ! -z "${TELEGRAM_CHAT_ID}" ] && [ ! -z "${TELEGRAM_BOT_TOKEN}" ];
then 
    sed -i '/tg_chat_id =/c\tg_chat_id = '"$TELEGRAM_CHAT_ID"''  ~/.koii-mc/config/config.toml

    sed -i '/tg_bot_token =/c\tg_bot_token = "'"$TELEGRAM_BOT_TOKEN"'"'  ~/.koii-mc/config/config.toml

    sed -i '/enable_telegram_alerts =/c\enable_telegram_alerts = 'true''  ~/.koii-mc/config/config.toml
else
    echo "---- Telgram chat id and/or bot token are empty --------"
fi

echo "------ Building and running the code --------"

cd $HOME
cd koii-mission-control

go build -o koii-mc
mv koii-mc $HOME/go/bin

echo "--------checking for koii binary path and updates it in env--------"

if [ ! -z "${KOII_BINARY_PATH}" ];
then 
    KOII_BINARY="$KOII_BINARY_PATH"
else 
    KOII_BINARY="koii"
fi

echo "----------- Setup koii-mc service------------"

echo "[Unit]
Description=koii-mc
After=network-online.target

[Service]
User=$USER
Environment="KOII_BINARY_PATH=$KOII_BINARY"
ExecStart=$HOME/go/bin/koii-mc
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target" | sudo tee "/lib/systemd/system/koii_mc.service"

echo "--------------- Start Koii-Mission-Control service ----------------"


sudo systemctl daemon-reload

sudo systemctl enable koii_mc.service

sudo systemctl start koii_mc.service

echo "** Done **"
