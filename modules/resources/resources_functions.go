package resources

import (
  "math"
)

func CodeGenerator(key uint32) (string) {
  // you dont need to know a lot about math to understand this one
  // since this is basic :)

  var allowed_ascii_letter string = "LmOSNXMzH37oDuFaR2scZJfCy0prwbVPB8ht5k4qdxWgQlKAU96E1njTeGvIYi"
  var res string
  var weight uint32

  for i := 1; i <= 6; i++ {
    weight = uint32(math.Pow(62, float64(6-i)))
    res += string(allowed_ascii_letter[key/weight])
    key %= weight
  }

  return res
}
