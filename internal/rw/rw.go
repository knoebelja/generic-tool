package rw

import "os"

/*
rw handles reading and writing files of various formats.
*/

const perm os.FileMode = 0777

type read func(filepath string) (object interface{})

type write func(filepath string, object interface{})
