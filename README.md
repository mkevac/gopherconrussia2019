# Accompanying code for Gophercon Russia 2019 talk about Bitmap Indexes

## Slides

Gophercon Russia 2019 [PDF](Gophercon Russia 2019.pdf)

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
name                                         time/op
SimpleBitmapIndex-12                         10.8µs ± 0%
SimpleBitmapIndexInlined-12                  8.88µs ± 0%
SimpleBitmapIndexInlinedAndNoBoundsCheck-12  8.33µs ± 0%
```

## Biggerbatch

```
name                                              time/op
BiggerBatchBitmapIndex-12                         1.18µs ± 0%
BiggerBatchBitmapIndexInlined-12                  1.31µs ± 0%
BiggerBatchBitmapIndexNoBoundsCheck-12            1.06µs ± 0%
BiggerBatchBitmapIndexInlinedAndNoBoundsCheck-12  1.12µs ± 0%
```

## Simplesimd

```
name                              time/op
SimpleSIMDBitmapIndex-12           160ns ± 0%
SimpleScalarFasterBitmapIndex-12  1.06µs ± 1%
SimpleScalarBitmapIndex-12        1.24µs ± 0%
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
