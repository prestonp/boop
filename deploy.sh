#! /bin/bash

rm -rf build

git clone git@github.com:prestonp/bloggo.git build

cd build

hugo
