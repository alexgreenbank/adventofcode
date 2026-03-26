# AoC 2025 Day 2

AoC Link: https://adventofcode.com/2025/day/2

## Intro

A seemingly simple one here, we're just checking for repeating sets of digits in a range of values, but a more optimal solution gets us into some surprising mathematical weeds.

(Remember that this write up is based on only knowing what is in part 1 of the puzzle until we have solved that part, so I'm not planning ahead for the twist that is coming in part 2.)

## Part 1

### Naive approach 1

TODO: link to 2_orig.pl

We loop through the ranges checking each value as we go.

We can ignore odd length strings as they cannot be split evenly into two.

For each candidate value we split the string into a front and back half and compare them, adding the value to the p1 answer if they match.

### Naive approach 2

TODO: link to 2_new.pl

We still loop through each value in the range but we can just use a regex to do the matching for us.

The perl is simply:-

```
	for $n ( $n..$b ) {
		$p1+=$n if( $n =~ /^(.+)\1$/ );
	}
```

The regex `^(.+)\1$` means:
* `^` matches the start of the string
* `(.+)` matches a string of one or more digits and keeps a note of the string
* `\1` matches the string we matched in the first step
* `$` matches the end of the string

Without the `^` and `$` anchors we can match something but have stray characters at the beginning or the end, e.g. a regex of `^(.+)\1` would match `23234` as the `^(.+)` would match the `23` at the start of the string, then the `\1` matches the second `23` substring, but there is nothing to prevent the extra `4` character from stopping it match something we don't want to match.

### Efficiency/speed of naive approaches

Both of the naive approaches are going to be slow over large ranges as they test every single possible number in the range.

A range of `1000000000-9999999999` is going to take a long time to check one value at a time.

### Faster approach

We can use a mathematical approach to finding the numbers in the range with repeating sets of digits.

Consider the number `11`. We can see this is formed of the substring `1` repeated twice. For a two-digit number of the form `nn` (where `n` is the same digit) the number must have a factor of `11`.

Similarly for the number `1212` we can factorise this number as `101 * 12` (it doesn't matter that `12` is not prime, the important fact is that any number of the form `xyxy` will have a factor of `101`.

So, in order to find any 4-digit numbers between `a` and `b` that are formed of repeated pairs of digits we don't need to check every number, we could just choose to check only multiples of 101 once we have found the first one.

If we are checking the range `1210-1310` we can, by inspection, note that we are only going to find `1212` in that range since `1111` is below the start of the range and `1313` is beyond the end of the range.

Programatically we can start at the bottom end of the range (`1210`) and we see that this is not divisible by `101` as `1210 % 101 = 99`.

We can find the first number above the start of the range that is divisible by `101` by doing:
* Find the remainder if we try to divide by `101`: `1210 % 101 = 99`
* Take this away from the start of the range: `1210 - 99 = 1111`
* Add on the divisor: `1111 + 101 = 1212`

We can then check to see if this value (`1212`) is within the range, which it is.

Note that if the start of the range IS divisible by `101` then we don't need to do anything to it, we use this as our "first match" value.

We can do something similar with the upper bound of the range.

If the upper bound is divisible by `101` then we don't need to do anything with it, this is our "last match" value.

Otherwise we take off whatever remainder there is mod `101` and to get to the "last match" value that is lower than the top of the range, e.g.
* Find the remainder if we try to divide by `101`: `1310 % 101 = 98`
* Take this away from the end of the range: `1310 - 98 = 1212`

This is our "last match" value.

So for the range `1210-1310` we have a "first match" value of `1212` and a "last match" value of `1212` as expected.

Since the first and last match values are the same it means the only match is that value (which is in the range) and that's our answer for this part.

But that was just a simple case, what about some other cases we have to consider.

What if there is no match at all?

Consider the range `1213-1312`. We can see by inspection that there are no matches in the range as `1212` is before the range and the next possible match would be `1313` which is beyond the end of the range. How would our sums above detect this?

In computing the "first match" we would do:
* Find the remainder if we try to divide by `101`: `1213 % 101 = 1`
* Take this away from the start of the range: `1213 - 1 = 1212`
* Add on the divisor: `1212 + 101 = 1313`

This "first match" value of `1313` is beyond the end of the range, which we take as a sign that there are no matches in that range. We don't need to find the "last match" value as it will also be outside the range.

What else do we need to consider?

Well, not all ranges have the same number of digits in every number.

If we just try and process the range `9998-10101` without noticing the upper end of the range has a different number of digits we may work out that this range contains both `9999` and `10100`, both of which are divisible by `101` but only `9999` is formed of a repeated pair of digits. If we counted `10100` as matching via this method it would lead to the wrong answer.

My input had a few instances where the upper end of the range had one more digit than the lower end of the range, but no instances where the upper end had 2 or more digits than the lower end of the range, but there is nothing in the problem description to prevent this from happening.

Numbers with 2 digits can be checked for division by `11`.

Numbers with 3 digits cannot be formed by a pair of repeated sets of digits.

Numbers with 4 digits can be checked for division by `101`.

Numbers with 5 digits cannot be formed by a pair of repeated sets of digits.

Numbers with 6 digits can be checked for division by `1001`.

...etc...

So, we need to be careful and if the range we are given spans multiple digit lengths then we need to split these up and check them separately.

A range of `15-108642` would need to be checked for:
* Numbers with in the range [15,99] can be checked for division by `11`
* No need to check numbers in the range [100,999]
* Numbers with in the range [1000,9999] can be checked for division by `101`.
* No need to check numbers in the range [10000,99999]
* Numbers with in the range [100000,108642] can be checked for division by `1001`.

Taking the range `100000-108642` as an example, we perform the calculations for first and last matches as described above to find:
* first match = `100100`
* last match =  `108108`

We could iterate through these values adding `1001` each loop and summing as we go, but there is a faster way to do this.

We know that there are 9 values that match (`100100, 101101, 102102, 103103, 104104, 105105, 106106, 107107 and 108108`).

If we consider the first part of these numbers we have a sequence of `100` to `108`.

If we take `99` away from each of these values we have the sequence `1, 2, 3, 4, 5, 6, 7, 8, 9`.

The sum of a sequence like this (from `1` to `n`) is given by the formula `n(n+1)/2` (see Triangular numbers).

So the sum of a sequence from `x+1` to `x+n` is given by the formula `xn+(n(n+1)/2)`.

For `x=99` and `n=9` we get the sum to be `(99*9)+(9*10/2) = 936`.

Almost there, this is just the sum of the first half of each number (100, 101, ..., 108) whereas we want the sum for the numbers `100100, 101101, ..., 108108`.

We can get this by taking `936` and mulitplying this by `1001`: `936 * 1001 = 936936`.

### Putting part 1 together

One way to split the program up is as follows:
* A main loop that processes the input and splits it up into the individual ranges
* A function `doRange()` that deals with one range at a time, which may span multiple digit lengths
* A function `sumDivs()` that deals with a subset of a range (with the same digit lengths) and an individual divisor

So if we get a range of `10-23` we call `doRange(10, 23)` which in turn makes a single call to `sumDivs(10, 23, 11)`.

If we get a range of `15-108642` we call `doRange(15, 108642)` which in turn calls:
* `sumDivs(15, 99, 11)`
* `sumDivs(1000, 9999, 101)`
* `sumDivs(100000, 108642, 1001)`
* and returns the sum of these return values.

### Part 1 speeds

(Note that I had already written the two perl solutions that solved part2 as well. I'm not going to rewrite them with just the part1 solution, there's no point.)

These use my real inputs (which I don't make public) so the timings should be representative of most AoC-ers inputs.

Naive string manipulation in perl (`2_orig.pl`):
```
$ time ./2_orig.pl 2.inp
part1: 53420042388
part2: 69553832684

real    0m3.705s
user    0m3.702s
sys     0m0.003s
```

Naive regex wrangling in perl (`2_new.pl`):
```
$ time ./2_new.pl 2.inp
part1: 53420042388
part2: 69553832684

real    0m1.232s
user    0m1.229s
sys     0m0.003s
```

The above is slightly faster.

Part 1 only in Golang:
```
$ go run 2.go 2.inp
part1: 53420042388
dur=166.403µs
Alloc = 0 MiB   TotalAlloc = 0 MiB      Sys = 7 MiB     NumGC = 0
```

That'll do for part 1.

## Part 2

TODO
