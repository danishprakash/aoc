#!/usr/bin/awk -f

BEGIN {
    # set field separator as none
    # turning it into paragraph-wise
    RS="";

    split("byr iyr eyr hgt hcl ecl pid", f)
    for (i in f) fs[f[i]] = ""
}

{
    c = 0
    for (i=1; i<=NF; i++) {
        split($i, k, ":")
        if (k[1] in fs) c++
    }
    if (c == 7) res++
}

END { print res }
