#!/bin/bash -e

read -r a
read -r b

sum=$((a+b))
diff=$((a-b))
product=$((a*b))
quotient=$((a/b))
echo "$sum"
echo "$diff"
echo "$product"
echo "$quotient"
