//go:build !pro
// +build !pro

package main

func init() {
  features = append(features,
    "Non-Pro Feature #1",
    "Non-Pro Feature #2",
  )
}
func init() {
  features = append(features,
    "Non-Pro Feature #1(dup)",
    "Non-Pro Feature #2(dup)",
  )
}
