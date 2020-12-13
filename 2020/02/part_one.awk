#! /usr/bin/awk -f

BEGIN { FS = "[-\ :]" }
      {
          split($5, a, "")
          c = 0; for (i = 0; i < length(a); i = i + 1) { if (a[i] == $3) { c = c + 1 }}
          if (!(c == 0 || c < $1 || c > $2)) { n = n + 1 }
     }
END  { print n }
