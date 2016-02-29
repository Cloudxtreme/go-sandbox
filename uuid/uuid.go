package main

/*
#include <stdio.h>
#include <uuid/uuid.h>

extern char uuid_gen() {
	// typedef unsigned char uuid_t[16];
	uuid_t uuid;

	// generate
	uuid_generate_random(uuid);

	// unparse (to string)
	char uuid_str[37];
	uuid_unparse_lower(uuid, uuid_str);

	return printf("%s\n", uuid_str);
	// return uuid_str;
}
*/
import "C"

func main() {
	// C.uuid_gen(

	C.uuid_gen()
	// fmt.Printf("%v", vec)
}
