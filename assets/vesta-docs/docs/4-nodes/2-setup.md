---
sidebar_position: 2
---
# Installing Vestad

## Pre-Requisites
See [Set-up](../1_install.md).

### Firewall Configuration
Configure UFW to only accept traffic on ports we use.

```bash
ufw limit ssh/tcp comment 'Rate limit for openssh server'
ufw default deny incoming
ufw default allow outgoing
ufw allow 26656/tcp comment 'Vesta - Cosmos SDK/Tendermint P2P'
ufw allow 26657/tcp comment 'Vesta - Cosmos SDK/Tendermint P2P'
ufw enable
```

:::tip

Perform the next follow steps as your `vesta` user with 'sudo' permissions

:::

### Creating a Service
You may want the daemon to run without you needing to supervise it. To turn the executable into a service follow these steps.

First create the service file `/etc/systemd/system/vestad.service`

```sh
sudo nano /etc/systemd/system/vestad.service
```

Copy and paste the follow into the service file: (you may need to edit it if you've set a custom home directory location)

```conf
[Unit]
Description=Vesta Validator
After=network.target

[Service]
Group=vesta
User=vesta
WorkingDirectory=/home/vesta
ExecStart=/home/vesta/go/bin/vestad start
Restart=on-failure
RestartSec=3
LimitNOFILE=8192

[Install]
WantedBy=multi-user.target
```

Update systemd and enable the service file.
```sh
sudo systemctl daemon-reload
sudo systemctl enable vestad.service
```

## Building from Source

Replace `<VERSION>` with the current running version.
```sh
git clone https://github.com/VestaProtocol/vesta.git
cd vesta
git fetch -a
git checkout <VERSION>

make install
```

From there you will be able to use `vestad`, ex:
```sh
vestad version
```