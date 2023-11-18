# RANGEMAP
This tool allows you to scan ip and port in given range

# USAGE
./rangemap -dr github.com:50-80 this will scan ports between 50 and 80

./rangemap -r 127.0.0.1-10:80 this will scan 80 port on ips between 127.0.0.1 and 127.0.0.10

-dr arguament does not allows you to use range on address or domain
-r arguament allows you tou use range on just ip, you can't event type domain

Also you can use -h arguament to get help
