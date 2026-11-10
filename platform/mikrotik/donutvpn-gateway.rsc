# Sends the network of a MikroTik (RouterOS 7) through donutvpn-core in router mode, running on
# another Linux box or OpenWrt router, or as a container on this router (docs/router-mode.md).
#
#   :global donutvpnGateway 172.17.0.2
#   /import donutvpn-gateway.rsc
#
# donutvpnGateway is the IPv4 address of the box running `donutvpn-core run --router`. It must be on
# a subnet of its own (a container's veth, a VLAN, another port), not on the LAN: on the LAN its
# replies reach the devices directly, and RouterOS drops the connections it sees half of.
#
# What it does, all marked with the comment "donutvpn" (donutvpn-gateway-remove.rsc takes it out):
# - a routing table "donutvpn" whose default route is the box; while the box does not answer pings,
#   the LAN goes out directly as before
# - connections from the WAN (to this router, and port forwards) are marked once at their start;
#   their replies keep the main table, so they still work
# - the LAN's other traffic to the internet - TCP, UDP and ICMP - takes the "donutvpn" table
# - the router's DNS asks the box, so names are resolved through the proxy, not by the ISP
# - fasttrack is turned off: a fasttracked packet skips the routing mark

:global donutvpnGateway
:if ([:typeof $donutvpnGateway] = "nothing" || [:typeof $donutvpnGateway] = "nil") do={
  :error "set the box's address first: :global donutvpnGateway 172.17.0.2"
}
:local gw [:toip $donutvpnGateway]
:if ([:typeof $gw] != "ip") do={ :error "donutvpnGateway is not an IPv4 address: $donutvpnGateway" }

/routing table
:if ([:len [find name=donutvpn]] = 0) do={ add name=donutvpn fib comment=donutvpn }

/ip route
remove [find comment=donutvpn]
add dst-address=0.0.0.0/0 gateway=$gw routing-table=donutvpn check-gateway=ping comment=donutvpn

/ip firewall address-list
remove [find list=donutvpn-direct]
:foreach net in={10.0.0.0/8; 172.16.0.0/12; 192.168.0.0/16; 100.64.0.0/10; 169.254.0.0/16; 224.0.0.0/4; 255.255.255.255/32} do={
  add list=donutvpn-direct address=$net comment=donutvpn
}

/ip firewall mangle
remove [find comment=donutvpn]
add chain=prerouting in-interface-list=WAN connection-state=new action=mark-connection \
    new-connection-mark=donutvpn-inbound passthrough=yes comment=donutvpn
add chain=prerouting in-interface-list=LAN src-address=("!" . $gw) dst-address-list=!donutvpn-direct \
    connection-mark=no-mark action=mark-routing new-routing-mark=donutvpn passthrough=yes comment=donutvpn

:foreach rule in=[/ip firewall filter find where action=fasttrack-connection and !disabled] do={
  /ip firewall filter set $rule disabled=yes comment=("donutvpn: fasttrack off; " . [/ip firewall filter get $rule comment])
}

:local servers [/ip dns get servers]
:if ([:len [/ip firewall address-list find where list=donutvpn-dns-was]] = 0) do={
  # 0.0.0.0 stands for no servers of its own (only the dynamic ones), for the removal to give back
  /ip firewall address-list add list=donutvpn-dns-was address=0.0.0.0 comment=donutvpn
  :foreach s in=$servers do={ /ip firewall address-list add list=donutvpn-dns-was address=$s comment=donutvpn }
}
/ip dns set servers=$gw

:put "donutvpn: the LAN goes through $gw (remove with /import donutvpn-gateway-remove.rsc)"
