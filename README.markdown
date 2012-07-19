# Nemo

Nemo is a simple dictionary-based passphrase generator written in Go. It generates "secure" (by the means of length and complexity) passphrases that are still easy to memorize.  

# Usage

The only requirement to run Nemo successfully is to provide a valid dictionary file. Everything else can be taken from the default values. 

The possible command line options are: 

````bash
-d                  whether the generated passphrases have to be distinct from  true
                    each other
-f                  comma separated list of random fill characters              "0,1,2,3,4,5,6,7,8,9,_,!,?"
-fill-before        whether to insert a fill character before or after a word   false
-fill-length-min    minimum length of the random fill insert                    1
-fill-length-max    maximum length of the random fill insert                    3
-i                  the dictionary input file(s)                                "language.dict"
-l                  the separator used to split the input file                  "\n"
-min                minimum passphrase length                                   30
-n                  number of passphrases to generate                           10
-r                  comma separated list of old new characters to replace in    "u\",ue,a\",ae,o\",oe,A\",Ae,O\",Oe,U\",Ue"
                    every dictionary word that is used for an actual passphrase
````

# Where can I find good wordlists / dictionaries?
Take a look here: http://www.aircrack-ng.org/doku.php?id=faq#where_can_i_find_good_wordlists. 

# inspired by

This tool is inspired by http://xkcd.com/936/.