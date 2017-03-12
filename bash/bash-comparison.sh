#!/bin/bash -e

main() {
  read -r x
  read -r y

  if [ "$x" -gt "$y" ]; then
    echo "X is greater than Y"
  elif [ "$x" -lt "$y" ]; then
    echo "X is less than Y"
  else
    echo "X is equal to Y"
  fi
}

main
