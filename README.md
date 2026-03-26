# adventofcode
Solutions for Advent of Code

A bunch of solutions to the puzzles at https://adventofcode.com/

Note that the [AoC creator has asked that public repos of solutions do not contain the puzzle text or inputs](https://adventofcode.com/2025/about#faq_copying). I keep these in a private repo. Please don't ask for my inputs.

## Goals

I have got all 524 stars already but the code is in various private repos and in various states. I would like to upload clean solutions for all days in Golang, Perl and C.

Depending on how busy I am in December I would first aim to get the stars using whatever language is most suitable. This would often be Perl. After that I can move on to creating a solution in one of the other languages (usually Golang next). The approach to solving the problem may change as I understand more about the problem or recognise the underlying "more optimal" algorithm.

I am also attempting to have all years complete in under 1s of execution time (per year) if possible. These should be generic solutions (e.g. capable of solving any inputs that an AoC participant may receive) and not optimised for my specific input.

## Progress

| Year | Golang | Perl | C | TotTime |
| ---- | --- | --- | --- | --- |
| 2015 | 0/50 | 0/50 | 0/50 | . |
| 2016 | 0/50 | 0/50 | 0/50 | . |
| 2017 | 0/50 | 0/50 | 0/50 | . |
| 2018 | 0/50 | 0/50 | 0/50 | . |
| 2019 | 0/50 | 0/50 | 0/50 | . |
| 2020 | 0/50 | 0/50 | 0/50 | . |
| 2021 | 0/50 | 0/50 | 0/50 | . |
| 2022 | 0/50 | 0/50 | 0/50 | . |
| 2023 | 0/50 | 0/50 | 0/50 | . |
| 2024 | 0/50 | 0/50 | 0/50 | . |
| 2025 | 3/24 | 4/24 | 0/24 | . |

## Overall ethos

I'm not aiming for perfect code. As mentioned above the majority of the Perl code was implemented quickly to simply get the stars so I could get on with my day.

With all implementations I am purely relying on built-in functionality/libraries. I do not want to have to rely on third party code that I have not written.

Obviously I have not written a Perl interpreter or Golang compiler of my own, but the line has to be drawn somewhere. I can write code to implement a hash-table if need be (and I do have to do this if I want one in C) but I can rely on Perl or Golangs implementations.

What I don't want to do is solve an AoC problem by importing some magic library, pass a few bits of my input to it and extract the result. I (personally) want to understand how it all works, and I know how hash tables, recursion, etc works so I don't need to go over this again each time myself.

My perl code won't have any tests, nor will my C code, but my Golang code probably will.

Sometimes I'll also write a solution in python or some other language (shell/awk/sed/...).

## Per language notes

### Perl

Where possible I use `use strict;` and `use warnings;` in my perl programs.

Sometimes I skip them if I have a particularly short implementation but my goal here is not to code golf my way to the shortest solution (as they are often unreadable).

### Golang

For Golang solutions I'm trying to implement understandable code with some appropriate tests. They aren't as polished as I would like, and I don't intend to produce separate modules/libraries that can be reused between puzzles. If I need to borrow code from a different puzzle I'll simply copy-and-paste it. I want my solutions to be standalone implementations.

### C

Sometimes it's fun to remind yourself just how much higher level langauges (perl, Golang) give you that just aren't in base languages like C.
