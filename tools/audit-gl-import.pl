#!/bin/perl

use strict;
use warnings;

use File::Find;

my $prefix = "github.com/go-gl/gl/";
my $suffix = "v4.1-core/gl";

my @goFiles = ();
my $count = 0;
my $found = 0;
find_goFiles(qw{cmd pkg});
foreach my $file (sort @goFiles) {
  $found++;
  $count += audit($file, $prefix, $suffix);
}

if ($found == 0) {
  die "ERROR: No go files found/audited";
} elsif ($count > 0) {
  die "ERROR: Unexpected go-gl version included";
}

sub audit {
  my ($file, $prefix, $suffix) = @_;
  my $count = 0;
  my $lNum = 0;
  open(my $IN, '<', $file);
  foreach my $line (<$IN>) {
    $lNum++;
    if ($line =~ m{"$prefix(.+)"}) {
      my $foundVer = $1;
      if ($foundVer ne $suffix) {
        print "Imported go-gl in $file\n";
        print "  - expected '$suffix'\n";
        print "  - found    '$foundVer' at line $lNum\n";
        $count++;
      }
    }
  }
  close($IN);
  return $count;
}

sub find_goFiles {
  my @folders = @_;
  find(\&wanted, @folders);
}

sub wanted {
  if ($File::Find::name =~ m{[.]go$}) {
    push(@goFiles, $File::Find::name);
  }
}
