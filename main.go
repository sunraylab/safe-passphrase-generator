package main

import (
 "fmt" 
  
 "github.com/tyler-smith/go-bip39"
)

func main(){
  // Generate a mnemonic for memorization or user-friendly seeds
  entropy, _ := bip39.NewEntropy(128)
  mnemonic, _ := bip39.NewMnemonic(entropy)

  // Display mnemonic and keys
  fmt.Println("Mnemonic: ", mnemonic)
}
