#!/bin/bash

go build -o bookings cmd/web/*.go && ./bookings -dbname=bookings -dbuser=postgres -dbpass=shub@123 -cache=false -production=false 
