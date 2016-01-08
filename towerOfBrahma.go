// towerOfBrahma
package Tower

import (
	"strconv"
	"fmt"
)

type Terror struct {
	num int
	msg string
}


type stack struct {
	name string
	numID int
	count int
	rings [100]int

}

type moves struct {
	ring int	//ring number to move
	from int	//from stack no. {1 to 3}
	to int		//to stack no. {1 to 3} but different from "from"
}

func DebugPrint(msg string) {
//	fmt.Println(msg)
}

func (s *stack) push(n int) Terror {
		reterr := Terror{0,"No Error"}
		DebugPrint("Push::Pushing "+strconv.Itoa(n)+" on the stack "+s.name)
		if s.count == 0 {
			s.rings[s.count] = n
			s.count++
//			fmt.Println("Push::the stack has ",s.count," items we are adding ",n)
			return (reterr)
		} else {
			if s.rings[s.count-1] > n {
				s.rings[s.count] = n
				s.count++
//				fmt.Println("Push::the stack has ",s.count,"items and the top item is ",s.rings[s.count-1]," we are adding ",n)
				return (reterr)
			} else {
//			fmt.Println("Push::the stack has ",s.count,"items and the top item is ",s.rings[s.count-1]," we are adding ",n)
			reterr.msg = "Push::the ring to be moved is bigger than the one on top"
//			fmt.Println(reterr.msg)
			return reterr
			}
	}
}
	
func (s *stack) pop() (int,Terror) {
	reterr := Terror{0,"Pop::No Error"}
	if s.count == 0 {
		reterror := Terror{1,"Pop::the stack is empty"}
		return 0,reterror
	} else {
		s.count--
		x := s.rings[s.count]
		DebugPrint("Pop::Popped "+strconv.Itoa(x)+" off the stack "+s.name+" s.count is"+strconv.Itoa(s.count))
		return x,reterr
	}
}
	
func print()  {
//	fmt.Println()
}
/*
func DumpStacks() {
	fmt.Print("A has "+strconv.Itoa(A.count)+" values are ")
	for i := 0;i < A.count;i++ {fmt.Print(A.rings[i]," ")}
	print()
	fmt.Print("B has "+strconv.Itoa(B.count)+" values are ")
	for i := 0;i < B.count;i++ {fmt.Print(B.rings[i]," ")}
	print()
	fmt.Print("T has "+strconv.Itoa(T.count)+" values are ")
	for i := 0;i < T.count;i++ {fmt.Print(T.rings[i]," ")}
	print()
}
*/
var A,B,T stack
var hmany int
var nmoves int
var allMoves [200]moves
var aMoves [200]int

func shiftOne(From, To *stack) {
	fmt.Println("---->SShiftOne::Moving one ring no.",From.rings[From.count - 1]," from stack ",From.name," to stack ",To.name)
	allMoves[nmoves].ring = From.rings[From.count - 1]
	allMoves[nmoves].from = From.numID
	allMoves[nmoves].to = To.numID
	aMoves[nmoves] = 100*allMoves[nmoves].ring+10*allMoves[nmoves].from+allMoves[nmoves].to
	nmoves++
//	fmt.Println("---->ShiftOne: nmoves is",nmoves)
	rN, err := From.pop()
	if err.num == 0 {
		err = To.push(rN)
		if err.num != 0 {
			panic(err.msg)
		}
	} else {panic(err.msg)}
//	DumpStacks()
}

func MoveRings(From *stack,HM int,To *stack,Using *stack) {
//	DumpStacks()
	fmt.Println("MoveRings::Moving ",HM," Rings from ",From.name," To ",To.name," Using ",Using.name)
	if HM >1 {MoveRings(From,HM-1,Using,To)}
	shiftOne(From,To)
	if HM>1 {MoveRings(Using,HM-1,To,From)}
}
/*
func PrintStacks() {
	for i:= 1; i < 5;i++ {fmt.Println("    |")}
	fmt.Print(A.name," |")
	if A.count > 0 {
		for i := 0; i < A.count; i++ { fmt.Print(" ",A.rings[i])}
		fmt.Println(">")
	}
	for i:= 1; i < 5;i++ {fmt.Println("    |")}
	fmt.Print(B.name," |")
	if B.count > 0 {
		for i := 0; i < B.count; i++ { fmt.Print(" ",B.rings[i])}
		fmt.Println(">")
	}
	for i:= 1; i < 5;i++ {fmt.Println("    |")}
	fmt.Print(T.name," |")
	if T.count > 0 {
		for i := 0; i < T.count; i++ { fmt.Print(" ",T.rings[i])}
		fmt.Println(">")
	}
	fmt.Println("PrintStack::So this is what stacks look like")
}
*/
func TowerOfBramha(nRings int, moves *[200]int) int{
	nmoves = 0
	A.name = "A"
	A.numID = 1
	B.name = "B"
	B.numID = 2
	T.name = "T"
	T.numID = 3
	hmany := nRings
	for i := hmany;i > 0;i-- {A.push(i)}
//	fmt.Println("Mail::We are starting %d rings",hmany)
//	PrintStacks()
	//Now we do a recursive function to move rings from A to C using B
	MoveRings(&A,hmany,&B,&T)
//	PrintStacks()
//	fmt.Println("Number of moves to move ",hmany," rings is ",nmoves)
//	fmt.Println("List of Moves")
	for i := 0; i < nmoves;i += 10 {
		for j := i; j < i + 10 && j < nmoves; j++ {
//		   fmt.Print(j,":",allMoves[j].ring," ",allMoves[j].from,allMoves[j].to,":-",aMoves[j],";")
		}
//	fmt.Println()
	}
	for j := 0; j < nmoves;j++ {
		moves[j] = aMoves[j]
//		fmt.Print(aMoves[j],";",moves[j])
	}
	return nmoves
}
