#!/bin/bash

for f in $(ls)
do
	x="$(echo $f | sed -r 's/(.*)_(.\..)(.*)/\2_\1\3/g')"
	mv $f $x
done	
