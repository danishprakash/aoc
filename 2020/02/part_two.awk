#! /usr/bin/awk -f

BEGIN { FS = "[- :]" }
      {
          split($5, a, "")
          for (i = 0; i < length(a); i = i + 1) {if (a[i] == $3); c = c + 1}
          if ((a[$1] == $3 && a[$2] != $3) || (a[$1] != $3 && a[$2] == $3)) { n = n + 1 }
      }
END   { print n }
