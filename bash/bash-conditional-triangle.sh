#!/bin/bash -e

main() {
  read -r a
  read -r b
  read -r c

  if [ "$a" -eq "$b" ] && [ "$a" -eq "$c" ]; then
    echo "EQUILATERAL"
  elif [ "$a" -eq "$b" ] || [ "$a" -eq "$c" ] || [ "$b" -eq "$c" ]; then
    echo "ISOSCELES"
  else
    echo "SCALENE"
  fi
}

main
