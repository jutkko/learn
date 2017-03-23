#!/bin/bash -e

grade() {
  scores="$1"

  eng="$(echo "$scores" | awk -F" " '{print $2}')"
  math="$(echo "$scores" | awk -F" " '{print $3}')"
  science="$(echo "$scores" | awk -F" " '{print $4}')"
  average="$(echo "scale = 5; ($eng + $math + $science)/3" | bc)"

  if [ "$(echo "$average >= 80" | bc)" -eq 1 ]; then
    echo "$scores : A"
  elif [ "$(echo "$average >= 60" | bc)" -eq 1 ]; then
    echo "$scores : B"
  elif [ "$(echo "$average >= 50" | bc)" -eq 1 ]; then
    echo "$scores : C"
  else
    echo "$scores : FAIL"
  fi
}

main() {
  while read -r line; do
    grade "$line"
  done
}

main
