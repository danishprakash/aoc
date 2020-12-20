#!/usr/bin/awk -f

{
    split($0, a, "")
    if (a[(((NR - 1) * 3) % length($0)) + 1] == "#") { c++ }
}
END { print c }
