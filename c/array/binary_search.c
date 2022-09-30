#include <stdio.h>
#include <stdlib.h>

/*
 * [left, right]
 * 
 * */

int binary_search(int nums[], int arrlen, int target);

// gcc binary_search.c -o bs
// ./bs <target>
int main(int argc, char *argv[]) {
  int nums[] = {1, 3, 12, 44, 65, 213};

  int arrlen = sizeof(nums) / sizeof(nums[0]);
  int target = atoi(argv[1]);

  printf("%d\n", binary_search(nums, arrlen, target));
  return 0;
}

int binary_search(int nums[], int arrlen, int target) {
  int low, high, middle;

  low = 0;
  high = arrlen;
  middle = (low + high) / 2;

  while(low <= high) {

    if (nums[middle] < target) {
      low = middle+1;
    } else if (nums[middle] > target){
      high = middle-1;
    } else {
      return nums[middle];
    }

    middle = (low + high) / 2;
  }

  return -1;
}

