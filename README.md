# Accompanying code for Gophercon Russia 2019 talk about Bitmap Indexes

## Code structure

| Directory   | Description                                              |
| ----------- | -------------------------------------------------------- |
| simple      | Simple implementation for a Bitmap Index in Go.          |
| biggerbatch | Similar to previous one, but using 64 bit batches.       |
| simplesimd  | Implementation in assembly for scalar and SIMD versions. |
| roar        | Implementation using roaring bitmaps Go modules.         |
| pilosa      | Implementation using pilosa DB.                          |


## Pilosa

1. Download and run the pilosa
2. Run the provided program
   ```
   $ go run pilosa.go 
   2019/03/30 20:41:12 filling the data...
   2019/03/30 20:41:47 finished filling the data
   2019/03/30 20:41:47 got 2796 columns
   ```