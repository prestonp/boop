#! /bin/bash

set -x
rm -rf build

git clone git://github.com/prestonp/bloggo.git build

cd build

hugo

aws s3 sync ./public s3://blog.preston.io --region us-west-1
