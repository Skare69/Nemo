package main

import (
    "fmt"
    "log"
    /*"os"*/
    "io/ioutil"
    "flag"
    "time"
    /*"math/big"
    crand "crypto/rand"*/
    "math/rand"
    "strings"
    /*"encoding/json"*/
)

// Declare command line arguments
var input_files = flag.String("i", "language.dict", "the dictionary input file(s)") // TODO enable support for more than one input file
var separator = flag.String("l", "\n", "the separator used to split the input file")
var fill = flag.String("f", "0,1,2,3,4,5,6,7,8,9,_,!,?", "comma separated list of random fill characters")
var fill_length_min = flag.Int("fill-length-min", 1, "minimum length of the random fill insert")
var fill_length_max = flag.Int("fill-length-max", 3, "maximum length of the random fill insert")
var fill_before = flag.Bool("fill-before", false, "whether to insert a fill character before or after a word") // TODO replace this with an integer: 0 none, 1 before, 2 after, 3 both
var min_pass_length = flag.Int("min", 30, "minimum passphrase length")
var number_gen_passp = flag.Int("n", 10, "number of passphrases to generate")
var replace_chars = flag.String("r", "u\",ue,a\",ae,o\",oe,A\",Ae,O\",Oe,U\",Ue", "comma separated list of old new strings to replace in every dictionary word that is used for an actual passphrase")
var distinct = flag.Bool("d", true, "whether the generated passphrases have to be distinct from each other.")

func main() {
    // Read the command line arguments. 
    flag.Parse()
    
    flags := flag.Args()
    // Check for unrecognized command line arguments and notify the user
    if len(flags) != 0 { 
        fmt.Println("Unrecognized arguments: ", flags)
    }
    
    // Seed the RNG
    rand.Seed(time.Now().UnixNano()) // TODO only pseudo random, better use crypto/rand
    
    // Read the dictionary files. 
    // TODO read more than one file from input_files
    dict, err := readDictionary(*input_files, *separator)
    if err != nil {
        panic(err)
    }
    if len(dict) == 0 {
        log.Fatal("Input dictionary is empty. Exiting.")
    }
    
    output := []string{}
    // Repeat until enough passphrases are generated
    for i := 0; i < *number_gen_passp; i++ {
        new_passphrase := generatePassphrase(dict, strings.Split(*fill, ","))
        if checkUniqueness(output, new_passphrase) {
            output = append(output, new_passphrase)
        } else {
            i--
        }
    }
    
    // Write the passphrases to the console
    for key, _ := range output {
        fmt.Println(output[key])
    }
}

// Open and read the dictionary file and push its contents into a string array
func readDictionary(dict_file string, sep string) (dict[]string, err error) {
    // TODO maybe better check the filesize first ...
    b, err := ioutil.ReadFile(dict_file)
    if err != nil {
        return nil, err
    }
    dict = strings.Split(string(b), sep)
    return
}

// Check if a given passphrase is already present in the output array
func checkUniqueness(output []string, new_passphrase string) (is_unique bool) {
    // TODO properly check for uniqueness
    if !*distinct {
        is_unique = true
        return
    }
    is_unique = true
    for _, value := range output {
        if new_passphrase == value {
            is_unique = false
        }
    }
    return
}

// Generate the actual passphrase
func generatePassphrase(words []string, fills []string) (gen_passphrase string) {
    if len(words) == 0 {
        log.Fatal("Dictionary import is empty. Exiting.")
    }
    // Create the passphrase string: 
    gen_passphrase = ""
    for len(gen_passphrase) < *min_pass_length { // Check if the passphrase's length is equal or greater than the minimum length. 
        // If fill_before = true insert y random characters (before the word)
        if *fill_before {
            for i := 0; i < getRandomInt(*fill_length_min, *fill_length_max); i++ {
                gen_passphrase += getRandomWord(fills)
            }
        }
        
        // Add the next word to the passphrase string. 
        gen_passphrase += getRandomWord(words)
        
        // If fill_before = false insert y random characters (after the word)
        if !*fill_before {
            for i := 0; i < getRandomInt(*fill_length_min, *fill_length_max); i++ {
                gen_passphrase += getRandomWord(fills)
            }
        }
    }
    return
}

// Replace all chars in raw with replace_with (oldnew ...string)
func replaceSpecialChars(raw string, replace_with string) (fixed string) {
    r := strings.NewReplacer(strings.Split(replace_with, ",")...)
    fixed = r.Replace(raw)
    return
}

// Return a random word from dict
func getRandomWord(words []string) string {
    if len(words) == 0 {
        return ""
    }
    // math.rand implementation
    //cr, _ := crand.Int(crand.Reader, big.NewInt(int64(len(i))))
        // TODO: how to convert *big.Int to int64 / int ?
//    rand.Seed(time.Now().UnixNano())
    r := rand.Intn(len(words))
    
    if len(*replace_chars) != 0 {
        words[r] = replaceSpecialChars(words[r], *replace_chars)
    }
    
    // crypto.rand implementation (sadly this returns an int64 which I can't use to access my array index :(
    /*r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
    r = r
    fmt.Println(r)*/
    
    return words[r]
}

// Return a random integer value r with min <= r <= max
func getRandomInt(min int, max int) (r int) {
    r = 0
    if (min == 0 && max == 0) || min >= max {
        return
    }
    r = rand.Intn(max-min)
    r += min
    return
}
