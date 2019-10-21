# How to install
It is created for me, however you can use it

#### Create user
create user on server  
`useradd -d /home/myfdb -m myfdb`  

create folder on server for bin files  
`mkdir -p /opt/myfdb/bin`  

create folder on server for config for system.d  
`mkdir -p /etc/myfdb.d/`  

copy to server bin  
`scp bin/app root@mysuperserver.com:/opt/myfdb/bin/`  
copy to server system.d config  
`scp install/myfdb.service root@mysuperserver.com:/etc/systemd/system/`  
`scp install/myfdb.cfg root@mysuperserver.com:/etc/myfdb.d/`  

enable systemd unit myfdb.service on server  
`systemctl enable myfdb`  
start systemd unit myfdb.service on server  
`systemctl start myfdb`  

check status systemd unit myfdb.service on server  
`systemctl status myfdb`  

view logs for systemd unit myfdb.service on server  
`journalctl -eu myfdb`  

data files are in '/home/myfdb/db'  





Ilya  Scherbina  
mail: sch@myfantasy.ru  
telegram: @super_botan
phone: 