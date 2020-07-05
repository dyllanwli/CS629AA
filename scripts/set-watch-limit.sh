# !/bin/bash

echo fs.inotify.max_user_watches=424288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p