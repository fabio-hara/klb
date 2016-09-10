# security groups related functions

# Create a security group on EC2-Classic or EC2-VPC if argument
# vpcid is not empty.
fn aws_secgroup_create(name, desc, vpcid) {
	vpcarg = ()

	if $vpcid != "" {
		vpcarg = (
			"--vpc-id"
                        $vpcid
		)
	}

        grpid <= (
		aws ec2 create-security-group	--group-name $name
						--description $desc
						$vpcarg |
		jq ".GroupId" | xargs echo -n
	)

	return $grpid
}

fn aws_secgroup_delete(grpid) {
	aws ec2 delete-security-group --group-id $grpid
}
