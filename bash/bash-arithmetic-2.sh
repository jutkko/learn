#!/bin/bash -e

main() {
  read -r expression

  raw=$(printf "scale = 5; %s\n" "$expression" | bc)
  printf "%0.3f\n" "$raw"
}

main
