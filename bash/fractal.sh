#!/bin/bash -e

WIDTH=100
N=
LINE="____________________________________________________________________________________________________"

print_base_n() {
  iter="$1"
  rows=$(echo "2 ^ (5-$iter)" | bc)
  ones=$(echo "2 ^ $iter-1" | bc)

  for _ in $(seq 1 "$rows"); do
    string="$LINE"
    for i in $(seq 1 "$ones"); do
      string="$(echo string | sed s/_/1/$((WIDTH/)))"
    done
  done
}

print_n() {
  iter="$1"
  rows=$(echo "2 ^ (5-$iter)" | bc)
  ones=$(echo "2 ^ $iter" | bc)

  for _ in $(seq 1 "$rows"); do
    string="$LINE"
    for offset in $(seq 1 "$ones"); do
      string="$(echo $string | sed s/_/1/$((WIDTH/2-offset)))"
    done

    printf "%s\n" "$string"
  done
}

main() {
  read -r N

  print_base_n "$N"
  # print_n "$N"
}

main
