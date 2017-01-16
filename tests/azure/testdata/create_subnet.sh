#!/usr/bin/env nash

import ../../azure/login
import ../../azure/nsg
import ../../azure/vnet
import ../../azure/subnet

subnet = $ARGS[1]
resgroup = $ARGS[2]
vnet = $ARGS[3]
addr = $ARGS[4]
nsg = $ARGS[5]
vnetaddr = $ARGS[6]
location = $ARGS[7]

azure_login()
azure_nsg_create($nsg, $resgroup, $location)
azure_vnet_create($vnet, $resgroup, $location, $vnetaddr)
azure_subnet_create($subnet, $resgroup, $vnet, $addr, $nsg)