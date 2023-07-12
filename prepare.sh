#!/bin/bash

PROJECT_ROOT=$PWD

echo "compiling binaries ->"

echo ""
echo "### building foo..."
cd foo && make
cd $PROJECT_ROOT

echo ""
echo "### building bar..."
cd bar && make
cd $PROJECT_ROOT

echo ""
echo "### building zed..."
cd zed && make
cd $PROJECT_ROOT

echo ""
echo "### building etl..."
cd etl && make
cd $PROJECT_ROOT

echo ""
echo "<- compiling of binaries complete"