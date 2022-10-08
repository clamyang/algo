#include <stdio.h>
#include <stdlib.h>

int min_size_subarr(int nums[], int length, int target);

int main(int argc, int * argv[]) {
  int nums[] = {1, 2, 5, 2};
  int target = 8;

  int length = sizeof(nums) / sizeof(nums[0]);


  printf("%d\n", min_size_subarr(nums, length, target));
  return 0;
}

int min_size_subarr(int nums[], int length, int target) {
  int min_val = 0;

  for (int i = 0; i < length; i++) {
    int sum = nums[i];

    for (int j = i + 1; j < length; j++) {
      sum += nums[j];
      if (sum >= target) {
	int tmp = j - i + 1;
	if (min_val == 0 || tmp < min_val) {
	  min_val = tmp;
	}
        break;
      }
    }
  }
  return min_val;
}
