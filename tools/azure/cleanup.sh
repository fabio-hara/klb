#!/usr/bin/env nash

resgroups <= azure group list
resgroups <= split($resgroups, "\n")

for resgroup in $resgroups {
    filtered <= echo $resgroup | -grep "klb"
    echo "filtered["+$filtered+"]"
    splitted <= split($filtered, " ")
    size <= len($splitted)
    echo "size: " + $size
    for word in $splitted {
        echo "word[" + $word + "]"
    }
}
