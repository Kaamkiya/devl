#include <iostream>
#include <vector>

#include "1.h"

int main() {
	std::vector<int> correct {1, 1, 2, 6, 24, 120, 720, 5040, 40320};

	for (int i = 0; i < 7; i++) {
		if (solution::factorial(i) != correct[i]) {
			std::cout << "\033[31;1mIncorrect. " << i << " factorial should output " << correct[i] << ", not " << solution::factorial(i) << "\n";
		}
	}

	return 0;
}
