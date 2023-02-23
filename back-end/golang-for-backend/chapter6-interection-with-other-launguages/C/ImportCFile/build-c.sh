#!/bin/bash

# linux
#clang yjfunction.c -shared -fPIC -o libyjfunction.so

# Mac
clang yjfunction.c -shared -fPIC -o libyjfunction.dylib
