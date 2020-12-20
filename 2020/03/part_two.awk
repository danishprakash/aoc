#!/usr/bin/awk -f

{
    split($0, a, "")
    if (a[(((NR - 1) * 1) % length($0)) + 1] == "#") { c1++ }
    if (a[(((NR - 1) * 3) % length($0)) + 1] == "#") { c2++ }
    if (a[(((NR - 1) * 5) % length($0)) + 1] == "#") { c3++ }
    if (a[(((NR - 1) * 7) % length($0)) + 1] == "#") { c4++ }
    if ((NR)%2 && a[(int((NR - 1) / 2) % length($0)) + 1] == "#") { c5++ }
}
END { print c1 * c2 * c3 * c4 * c5 }

