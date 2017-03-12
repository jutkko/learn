#!/bin/bash -e

main() {
  read -r n

  sum=0
  for _ in $(seq 1 "$n"); do
    read -r number
    sum="$((sum+number))"
  done

  printf "%0.3f\n" "$(echo "scale = 5; $sum/$n" | bc)"
}

main
