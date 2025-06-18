# Conway's Game of Life

## Rules

1. **Any live cell with fewer than two live neighbors dies** (underpopulation)
2. **Any live cell with two or three live neighbors lives on** to the next generation (survival)
3. **Any live cell with more than three live neighbors dies** (overpopulation)
4. **Any dead cell with exactly three live neighbors becomes a live cell** (reproduction)

## Common Implementations

### In Memory matrix approach

This approach would require the software to store the whole grid in memory, using a matrix filled with zero's and one's to represent **alive** and **dead** cells. This implementation is simple, but is limited, concidering that the scope of this project is to accept a grid in the signed 64-bit range, this would require around 4 million terabytes of memory.

### Hashlife

Hashlife would work great for this, but from my research, it is not simple to implement and it's absolutely overengineering for a project like this.

### ✅ Sparce Matrix

This is the implementation on this project, since we can't represent every cell in the grid, let's only represent every living cell, this reduces the time and space complexity a lot. Basically whe iterate over all living cells and apply the rules for each one, instead of checking every single cell in the grid.


## Execution

> [!TIP]
> In order to have a predictable output each time, the output is ordered by X and Y.

To run this program, pass the Life 1.06 format file on the stdin and pipe the output to file (this helps if the result is too large)

```sh
go run main.go < sample.life > result.life
```


## Tests

To run the test, execute this command:
```sh
go test __tests__/main_test.go
```
