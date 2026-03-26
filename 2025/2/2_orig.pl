#!/usr/bin/perl

use strict;
use warnings;

my $p1=0;
my $p2=0;

while( my $l = <> ) {
	chomp( $l );
	foreach my $k ( split( /,/, $l ) ) {
		if( $k =~ /^(\d+)-(\d+)$/ ) {
			my $a = $1;
			my $b = $2;
			my $n=$a;
			while( $n <= $b ) {
				my $z = length($n);
				# part1
				if( $z % 2 == 0 ) {
					# We have an even length string
					my $first = substr( $n, 0, $z/2 );
					my $last = substr( $n, $z/2, $z/2 );
					if( $first eq $last ) {	
						# print "invalid 1 $n\n";
						$p1 += int($n);
					}
				}
				# part2 we consider multiples
				my $mult=2;
				while( $z >= $mult ) {
					if( $z % $mult == 0 ) {
						# Try to split string up into $mult parts of len $ml
						my $ml = $z/$mult;
						my $first=substr( $n, 0, $ml );
						# print "\t$n $mult first=[$first]\n";
						my $ok=1;
						my $mi=1;
						while( $mi < $mult ) {
							my $check = substr( $n, ($ml*$mi), $ml );
							# print "\t\tcheck = $check\n";
							if( $check ne $first ) {
								$ok=0;
								last;
							}
							$mi++;
						}
						if( $ok ) {
							$p2 += int($n);
							# print "invalid 2 $n\n";
							last;
						}
					}
					$mult++;
				}
				if( $n =~ /^(.{1,})\1+$/ ) {
					# print "invalRE 2 $n\n";
				}
				$n++;
			}
		} else {
			die "UNHANDLED [$l]";
		}
	}
}

print "part1: $p1\n";
print "part2: $p2\n";
