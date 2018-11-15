package logerror

import (
   "fmt"
)

// error handling
func CheckErr(err error) int {
    if err != nil {
        fmt.Printf("logerror: err: %v\n", err)
        return 1
    } else {
        return 0
    }
}


