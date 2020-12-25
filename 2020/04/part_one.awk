#!/usr/bin/awk -f

# turn into paragraph-wise mode
BEGIN { RS=""; }
{
    c = 0
    for (i=1; i<=NF; i++) { split($i, k, ":"); if (match(k[1], "byr|iyr|eyr|hgt|hcl|ecl|pid")) c++ }
    if (c == 7) res++
}
END { print res }
