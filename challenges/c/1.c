#include <stdio.h>

#include "1.h"

int main() {
	int correct[7] = {1, 1, 2, 6, 24, 120, 720};

	for (int i = 0; i < 6; i++) {
		if (factorial(i) != correct[i]) {
			printf("\033[31;1mIncorrect: %d factorial should output %d, not %d.\n", i, correct[i], factorial(i));
		}
	}

	return 0;
}
