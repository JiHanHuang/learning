#!/bin/bash

lpath=$(cd $(dirname 0); pwd)
cd $lpath

rm -rf build
mkdir build
cd build

cmake ..
if [ $? -eq 0 ]; then 
    make
    echo "Build finish!"
else
    echo "Build error!"
fi

