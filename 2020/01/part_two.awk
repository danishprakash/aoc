#!/usr/bin/awk -f

    { a[NR] = $0 }
END { for (i = 1; i < length(a) - 1; i = i + 1) {
            for (j = 1; j < length(a) - 1; j = j+1) {
                for (k = 1; k < length(a) - 1; k = k+1) {
                    if (a[i] + a[j] + a[k] == 2020) {
                        print a[i] * a[j] * a[k]
                        exit
                    }
                }
            }
        }
    }
