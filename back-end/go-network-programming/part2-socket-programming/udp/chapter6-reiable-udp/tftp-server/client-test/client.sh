#!/bin/bash

# tftp -i 127.0.0.1:7166 GET test

tftp <<EOF

mode binary
connect 127.0.0.1 7166
get server-resource.svg
quit

EOF