#!/usr/bin/gawk -f

# turn into paragraph-wise mode
BEGIN { RS=""; }
{
    c = 0
    for (i=1; i<=NF; i++) {
        split($i, item, ":"); k=item[1]; v=item[2];
        if (match(k, "byr|iyr|eyr|hgt|hcl|ecl|pid")) {
            switch(k) {
                case "byr": if (length(v) == 4 && v >= 1920 && v <= 2002) c++; break
                case "iyr": if (length(v) == 4 && v >= 2010 && v <= 2020) c++; break
                case "eyr": if (length(v) == 4 && v >= 2020 && v <= 2030) c++; break
                case "ecl": if (match(v, "amb|blu|brn|gry|grn|hzl|oth")) c++ ; break
                case "hcl": if (match(v, "^#[0-9a-f]{6}")) c++ ; break
                case "pid": if (length(v) == 9 && match(v, "[0-9].*")) c++ ; break # regex intervals is a mess with awk :/
                case "hgt": if (match(v, "cm")) { gsub(/[[:alpha:]]/, "", v); if (v >= 150 && v <= 193) c++ }
                    else if (match(v, "in")) { gsub(/[[:alpha:]]/, "", v); if (v >= 59 && v <= 76) c++ }
                    break
            }
        }
    }
    if (c == 7) res++
}
END { print res }
