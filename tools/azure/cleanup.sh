#!/usr/bin/env nash

resgroups <= azure group list --json
length <= echo $resgroups | jq ". | length"

echo $length
#for resgroup in $resgroups {
    #echo $resgroup
#}
