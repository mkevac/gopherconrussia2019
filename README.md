# Accompanying code for Gophercon Russia 2019 talk about Bitmap Indexes

## Code structure

| Directory   | Description                                              |
| ----------- | -------------------------------------------------------- |
| simple      | Simple implementation for a Bitmap Index in Go.          |
| biggerbatch | Similar to previous one, but using 64 bit batches.       |
| simplesimd  | Implementation in assembly for scalar and SIMD versions. |
| roar        | Implementation using roaring bitmaps Go modules.         |
| pilosa      | Implementation using pilosa DB.                          |

## Simple

```
name                                        time/op
SimpleBitmapIndex-8                         8.96µs ± 0%
SimpleBitmapIndexInlined-8                  9.03µs ± 0%
SimpleBitmapIndexNoBoundsCheck-8            8.02µs ± 0%
SimpleBitmapIndexInlinedAndNoBoundsCheck-8  8.07µs ± 1%
```

## Biggerbatch

```
name                                             time/op
BiggerBatchBitmapIndex-8                         1.14µs ± 0%
BiggerBatchBitmapIndexInlined-8                  1.39µs ± 0%
BiggerBatchBitmapIndexNoBoundsCheck-8            1.01µs ± 0%
BiggerBatchBitmapIndexInlinedAndNoBoundsCheck-8  1.39µs ± 0%
```

## Simplesimd

```
name                             time/op
SimpleSIMDBitmapIndex-8           154ns ± 0%
SimpleScalarFasterBitmapIndex-8  1.04µs ± 0%
SimpleScalarBitmapIndex-8        1.24µs ± 1%
```

## Roaring bitmaps

```
name                   time/op
RoaringBitmapIndex-8   10.9µs ±17%
CRoaringBitmapIndex-8  10.6µs ±16%

name                   alloc/op
RoaringBitmapIndex-8   13.2kB ± 1%
CRoaringBitmapIndex-8   16.0B ± 0%

name                   allocs/op
RoaringBitmapIndex-8     12.0 ± 0%
CRoaringBitmapIndex-8    2.00 ± 0%
```

## Pilosa

1. Download and run the pilosa
2. Run the provided program
   ```
   $ go run pilosa.go 
   2019/03/30 20:41:12 filling the data...
   2019/03/30 20:41:47 finished filling the data
   2019/03/30 20:41:47 got 2796 columns
   ```