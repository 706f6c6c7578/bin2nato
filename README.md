#### bin2nato - a binary file converter to NATO/HEX words

Purpose of this software is to convert small encrypted binary blobs to
NATO alphabet so that the sender can speak  the file to be transmitted
over the telephone and the receiving party  writes down the NATO
words as HEX values and then uses the second program to convert back
the HEX values to a binary file.

This approach should allow transmitting small binary blobs relatively fast
and  error-free internationally, compared to other solutions like mnemonicode
or the  PGP wordlist, which non-native English speakers may have problems with.

Usage for the sending party: **sender -i input file -o output file**.

**sender** outputs five word-pairs per line so that the listener has only to type
in five HEX values per line.

The listener simply types in the spoken NATO words as HEX values in his / her 
editor of choice and then uses **receiver -d -i input file -o output file**,  to
convert the HEX values back to the original binary blob.

Because both programs are the same and differ only in the wordlist one can
use them each on their own.

Example: 

Sender speaks: **Alfa-Zero, Foxtrot-Nine, Six-Eight, Three-Two, Bravo-Five**

Listener writes down: **A0 F9 68 32 B5**


