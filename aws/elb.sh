# ELB related functions

fn aws_elb_create(name, listeners, subnetid, secgrps) {
	dnsName <= (
		aws elb create-load-balancer
					--load-balancer-name $name
					--listeners $listeners
					--subnets $subnetid
					--security-groups $secgrps |
		jq ".DNSName" |
		xargs echo -n
	)

	return $dnsName
}

fn aws_elb_delete(name) {
	aws elb delete-load-balancer --load-balancer-name $name >[1=]
}
