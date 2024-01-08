#include <stdio.h>

int Reverse(int n) {
    int r = 0;
    while (n > 0) {
        r = r * 10 + n % 10;
        n = n / 10;
    }
    return r;
}

int main() {
    printf("%d\n", Reverse(100));
    printf("%d\n", Reverse(101));
    printf("%d\n", Reverse(123));
    printf("%d\n", Reverse(1));
    printf("%d\n", Reverse(1234));

    return 0;
}
