# Availset related functions

# azure_availset_create creates a new `availability set`.
# `name` is the availability set name
# `group` is the resource group name
# `location` is the azure region
fn azure_availset_create(name, group, location) {
	azure availset create --name $name --resource-group $group --location $location
}

# azure_availset_exists checks if the given
# availset exists inside the provided resource group.
# returns "true" if it exists, "false" otherwise.
fn azure_availset_exists(name, group) {
        azure availset list $group | -grep -w $name
        if $status == "0" {
                return "true"
        }
        return "false"
}

# azure_availset_delete deletes a exit `availability set`.
# `name` is the availability set name
# `group` is the resource group name
fn azure_availset_delete(name, group) {
	azure availset delete -q --name $name --resource-group $group
}
