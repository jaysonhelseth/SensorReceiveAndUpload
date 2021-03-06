# SensorReceiveAndUpload
Receive the Xbee (Zigbee) data and upload it to the cloud. There will be the ability to see the data too.

## Collecting Data
There will be a Go server running in systemd that will collect the Xbee data. Systemd files are stored in `/etc/systemd/system` and have a name such as `runlistener.service`.
```
# Example of a service file
[Unit]
Description=Sensor Receive

[Service]
User=pi
EnvironmentFile=/etc/default/sensor-receive
ExecStart=/home/pi/SensorReceiveAndUpload/PiSensorReceive
Restart=always

[Install]
WantedBy=multi-user.target
```
The EnvironmentFile is where you would store any environment variables that the program is looking for. Then run `sudo systemctl daemon-reload`. Now you can do the normal systemctl operations of `start, stop, status, enable, disable`. You can even use `journalctl`.

## Sending Data
There will be a go routine that will send the data every 3 minutes to a free tier of cockroachdb in the cloud.

## Displaying Data
This will be on a Raspberry Pi that has a small screen attached. The device will boot to a web page hosted on the device that will read from the sqlite table that stores the temperature information.

The current device has a small screen with xserver and no desktop installed. The .xinitrc file in the user account (autologin) looks like this:
```
#!/bin/sh

sudo systemctl start sensor-receive.service
unclutter -idle 0.1 &
exec chromium-browser --kiosk file:///home/pi/SensorReceiveAndUpload/ui/index.html
```
The sensor-receive service is started in the .xinitrc file so it has perfect timing. This should still allow the service to handle restarting itself if it fails.