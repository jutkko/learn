#!/bin/bash -e

main() {
  read -r char

  if [ "$char" == "y" ] || [ "$char" == "Y" ]; then
    echo "YES"
  else
    echo "NO"
  fi
}

main
