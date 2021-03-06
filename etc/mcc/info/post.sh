#!/bin/bash

ftp -n localhost <<EOF
quote user test
quote pass test
put 2024F72524FFBFF.xml
put 202600157800000.xml
put 202805001600000.xml
put 202B89B81600000.xml
put 202C05001600000.xml
put 202E05001600000.xml
put 20341500BF81FE0.xml
put 20383C480000000.xml
put 9B7B2059560E9D1.xml
put A029C2900D97591.xml
put A22E00027000000.xml
put A22F03504400001.xml
put A786492E70174C1.xml
put ADCC404C8400315.xml
put ADCCB8E29D80001.xml
put ADCDABC3C3C0001.xml
put BBAD5EE4A400191.xml
put BEEC0358DC00001.xml
put located.xml
put unlocated.xml
ren 2024F72524FFBFF.xml 2024F72524FFBFF.txt
ren 202600157800000.xml 202600157800000.txt
ren 202805001600000.xml 202805001600000.txt
ren 202B89B81600000.xml 202B89B81600000.txt
ren 202C05001600000.xml 202C05001600000.txt
ren 202E05001600000.xml 202E05001600000.txt
ren 20341500BF81FE0.xml 20341500BF81FE0.txt
ren 20383C480000000.xml 20383C480000000.txt
ren 9B7B2059560E9D1.xml 9B7B2059560E9D1.txt
ren A029C2900D97591.xml A029C2900D97591.txt
ren A22E00027000000.xml A22E00027000000.txt
ren A22F03504400001.xml A22F03504400001.txt
ren A786492E70174C1.xml A786492E70174C1.txt
ren ADCC404C8400315.xml ADCC404C8400315.txt
ren ADCCB8E29D80001.xml ADCCB8E29D80001.txt
ren ADCDABC3C3C0001.xml ADCDABC3C3C0001.txt
ren BBAD5EE4A400191.xml BBAD5EE4A400191.txt
ren BEEC0358DC00001.xml BEEC0358DC00001.txt
ren located.xml located.txt
ren unlocated.xml unlocated.txt
EOF
