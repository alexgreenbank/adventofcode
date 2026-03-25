#!/usr/bin/perl

use strict;

my $c=50;

my $p1=0;
my $p2=0;

while( my $l = <> ) {
	chomp( $l );
	# print "Doing: $l starting at $c\n";
	if( $l =~ /^([LR])(\d+)$/ ) {
		my $dir=$1;
		my $n=$2;
		if( $n >= 100 ) {
			# Simplify part 2
			$p2 += int($n/100);
			# print "Added several because of $n\n";
			$n %= 100;
		}
		# Part 2 just needs to check going across boundaries
		if( $dir eq "L" ) {
			if( $c != 0 && ($c - $n) < 0 ) {
				# print "Gone left over 0\n";
				$p2++;
			}
			$c -= $n;
			$c %= 100;
		} else {
			# Right
			if( ($c + $n) > 100 ) {
				# print "Gone right over 0\n";
				$p2++;
			}
			$c += $n;
			$c %= 100;
		}
	
		# print "Ends at $c\n";
		# Check for resting on 0
		if( $c == 0 ) {
			# print "Left on 0\n";
			$p1++;
			$p2++;
		}
	} else {
		die "UNHANDLED: [$l]";
	}

}

print "part1: $p1\n";
print "part2: $p2\n";
