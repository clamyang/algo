#include <stdio.h>
#include <stdlib.h>

void remove_target(int nums[], int arrlen, int target);
int remove_target_hard(int nums[], int arrlen, int target);

int main(int argc, char *argv[]) {
  int nums[] = {1, 2, 3, 4, 2, 5, 2};

  int length = sizeof(nums) / sizeof(nums[0]);
  
  //remove_target(nums, length, 2);
  int end_idx = remove_target_hard(nums, length, 2);

  for (int i = 0; i < end_idx; i++) {
    printf("%d\n", nums[i]);
  }

  return 0;
}

// easy soluation
void remove_target(int nums[], int arrlen, int target) {
  for (int i = 0; i < arrlen; i++) {
    if (nums[i] != target) {
      continue;
    }

    for (int j = i; j <= arrlen-1; j++) {
      nums[j] = nums[j+1];
    }

    nums[arrlen-1]=0;
  }
}

// hard soluation
int remove_target_hard(int nums[], int arrlen, int target) {
  // slow fast pointer
  // slow point to element that equals to target
  // fast point to element that not equals to target
  int slow, fast;

  slow = 0;
  for (fast = 0; fast < arrlen; fast++){
    if (nums[fast] != target) {
      nums[slow] = nums[fast];
      slow++;
    } 
  }

  return slow;
}
