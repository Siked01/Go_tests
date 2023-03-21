package uniq

import "flag"

type flagsMain struct {
	count     bool
	double    bool
	unique    bool
	fields    int
	chars     int
	ignoreReg bool
}

func Init(opts *flagsMain) {

	flag.BoolVar(&opts.count, "c", false, "count the number of occurrences of a string in the input data. Print this number before the string separated by a space.")
	flag.BoolVar(&opts.double, "d", false, "output only those lines that are repeated in the input data.")
	flag.BoolVar(&opts.unique, "u", false, "output only those lines that are not repeated in the input data.")
	flag.IntVar(&opts.fields, "f", 0, "ignore the first num_fields of fields in a row.")
	flag.IntVar(&opts.chars, "s", 0, "ignore the first num_chars characters in the string.")
	flag.BoolVar(&opts.ignoreReg, "i", false, "ignore the case of letters.")
}
