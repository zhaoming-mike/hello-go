#!/bin/bash
# author:zhaoming23@gmail.com
echo "git add *.go"
git add *.go

if [ -z $1 ];then
	echo "git commit -m 'no message'"
	git commit -m "no message"
else
	echo "git commit -m '$*'"
	git commit -m "$*"
fi
echo "git push..."
git push
