package piscine

import("github.com/01-edu/z01")

// boucle forte

func PrintComb(){    
    for x := '0'; x<= '7'; x++ {
      for y := x+1; y<= '8'; y++ {
         for z := y+1; z<= '9'; z++ {
        z01.PrintRune(x) 
        z01.PrintRune(y)
        z01.PrintRune(z)
           if x < '7' {
             z01.PrintRune(',') // afficher des virgules
        z01.PrintRune(' ') // afficher des espace 
           }
    }}}
       z01.PrintRune('\n') // afficher en ligne 

}