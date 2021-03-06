#!/usr/bin/env nash

import ../../azure/all

routetable     = $ARGS[1]
name           = $ARGS[2]
resgroup       = $ARGS[3]
address        = $ARGS[4]
hoptype        = $ARGS[5]

route <= azure_route_table_route_new($name , $resgroup, $routetable, $address, $hoptype)

azure_route_table_route_create($route)