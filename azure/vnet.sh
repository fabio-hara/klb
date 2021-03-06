# Virtual Network related functions

fn azure_vnet_create(name, group, location, cidr, dnsservers) {
        fn join(list, sep) {
            out = ""

            for l in $list {
                out = $out+$l+$sep
            }
            out <= echo $out | sed "s/"+$sep+"$//g"
            return $out
        }

        dns <= join($dnsservers, ",")
	az network vnet create --name $name --resource-group $group --location $location --address-prefixes $cidr --dns-servers $dns
}

fn azure_vnet_delete(name, group) {
	az network vnet delete --resource-group $group --name $name
}

fn azure_vnet_set_route_table(vnet, subnet, group, routetable) {
	az network vnet subnet set --name $subnet --resource-group $group --vnet-name $vnet --route-table-name $routetable
}
