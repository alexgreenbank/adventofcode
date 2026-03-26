#!/usr/bin/perl

use warnings;

foreach $k ( split( /,/, <> ) ) {
	( $n, $b ) = split( /-/, $k );
	for $n ( $n..$b ) {
		$p1+=$n if( $n =~ /^(.+)\1$/ );
		$p2+=$n if( $n =~ /^(.+)\1+$/ );
	}
}

print "part1: $p1\n";
print "part2: $p2\n";
