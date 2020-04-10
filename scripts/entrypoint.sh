#!/bin/sh

export PORT=80

/bin/migration --action=up
/bin/server