package main   

import "fmt" 

func main(){ 
	var  (x int 
		y string 
		z bool
		)     
	/* 
	?  pour assigner plusieur var d'un seul coup on aurai pu faire var ( liste des nom de variable + type )
	*/
	
	
	x= 42
	y= "BlueSheep"
	z= true
	//? y := 24   c'ette asignation ne fonctionne que dans une fonction

	// fmt.Printf("\nMon chiffre favoris est %v et vive H2G2",x)
	// fmt.Printf("\nMon pseudo est : %v", y)
	// fmt.Printf("\nJ'ai plus de 18 ans : %v", z)
	fmt.Printf("\nMon chiffre favoris est %v, je peux concaténé mon pseudo %v et mon bool %v comme ceci",x, y, z) //? on defini l'odre nous même
}
/*
!		%v me permet d'aficher ma variable dans le text
!		%n me permet de faire un saut de ligne mais il y a une autre méthode
*/
