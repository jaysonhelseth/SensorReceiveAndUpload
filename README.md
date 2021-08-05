# SensorReceiveAndUpload
Receive the Xbee (Zigbee) data and upload it to the cloud. Their will be the ability to see the data too.

## Collecting Data
There will be a Python script running in systemd that will collect the Xbee data and put it into a sqlite table. Systemd files are stored in `/etc/systemd/system` and have a name such as `runlistener.system`.
```
# Example of a service file
[Unit]
Description=Testing out a service

[Service]
User=jayson
WorkingDirectory=/home/jayson/projects/testService
ExecStart=/usr/bin/python3 test.py
Restart=always

[Install]
WantedBy=multi-user.target
```
Then run `sudo systemctl daemon-reload`. Now you can do the normal systemctl operations of `start, stop, status, enable, disable`. You can even use `journalctl`.

It will always update the same record over and over. Since this device will have an internet connection and correct time, it will add the time to the collected data. This table will not be the place for storing historical data. That will eventually be in the cloud repo for this system.

## Sending Data
There will be a cron task that runs every 2 minutes (maybe a longer wait to avoid too much traffic to the cloud) to push the data to the webhook I have in the cloud.

\* Cockroach labs is offering a free tier for life. That may be something worth looking into for saving historical data. Then we could look for trends or something.

## Displaying Data
This will be on a Raspberry Pi that has a small screen attached. The device will boot to a web page hosted on the device that will read from the sqlite table that stores the temperature information.

The current device has a small screen with xserver and no desktop installed. The .xinitrc file in the user account (autologin) looks like this:
```
#!/bin/sh

unclutter -idle 0.1 &

exec chromium-browser --kiosk http://localhost/index.php
```
