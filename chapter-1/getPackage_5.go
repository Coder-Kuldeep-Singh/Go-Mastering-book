package main
import (
	"fmt"
	"github.com/mactsouk/go/simpleGitHub"
)

func main() {
	fmt.Println(simpleGitHub.AddTwo(5,6))
}


//remove package and single file of downloaded package 
// go clean -i -v -x github.com/mactsouk/go/simpleGitHub

//remove an entire downloaded package
// go clean -i -v -x github.com/mactsouk/go/simpleGitHub        //files removed from that package
// Then
// rm -rf ~/go/src/github.com/mactsouk/go/simpleGitHub //to remove entire package
