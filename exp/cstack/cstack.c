#include <stdio.h>

#define nil 0

struct Stack {
  int val;
  struct Stack* rest;
};

int main() {
  struct Stack s = {0, nil};
  printf("hello\n");
}
