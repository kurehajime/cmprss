package cmprss

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCmprss(t *testing.T) {
	buf := &bytes.Buffer{}

	//Less is more.
	input := "Less is more."
	ans := "LssIsMre."
	Cmprss(strings.NewReader(input), buf)
	if buf.String() != ans {
		t.Errorf("Cmprss(x) =\n%s\n, want \n%s", buf.String(), ans)
		return
	}
	//Exclsion:len(str) <= 3
	buf = &bytes.Buffer{}
	input = "Cat"
	ans = "Cat"
	Cmprss(strings.NewReader(input), buf)
	if buf.String() != ans {
		t.Errorf("Cmprss(x) =\n%s\n, want \n%s", buf.String(), ans)
		return
	}
	//Exclsion:DoubleLwrCse
	buf = &bytes.Buffer{}
	input = "meat meet moot"
	ans = "MeatMeetMoot"
	Cmprss(strings.NewReader(input), buf)
	if buf.String() != ans {
		t.Errorf("Cmprss(x) =\n%s\n, want \n%s", buf.String(), ans)
		return
	}
	//Exclsion:FllUpprCse.
	buf = &bytes.Buffer{}
	input = "JPEG"
	ans = "JPEG"
	Cmprss(strings.NewReader(input), buf)
	if buf.String() != ans {
		t.Errorf("Cmprss(x) =\n%s\n, want \n%s", buf.String(), ans)
		return
	}
}

//ThnkDffrnt. https://www.youtube.com/watch?v=nmwXdGm89Tk
func Example_thnkdffrnt() {
	buf := &bytes.Buffer{}

	input := `Here’s to the crazy ones. The misfits. The rebels. The troublemakers. The round pegs in the square holes.
The ones who see things differently. They’re not fond of rules. And they have no respect for the status quo.
You can quote them, disagree with them, glorify or vilify them.
About the only thing you can’t do is ignore them. Because they change things. They invent. They imagine.
They heal. They explore. They create. They inspire. They push the human race forward.
Maybe they have to be crazy.
How else can you stare at an empty canvas and see a work of art? Or sit in silence and
hear a song that’s never been written? Or gaze at a red planet and see a laboratory on wheels?
We make tools for these kinds of people.
While some see them as the crazy ones, we see genius. Because the people who are crazy enough
to think they can change the world, are the ones who do.`
	Cmprss(strings.NewReader(input), buf)
	fmt.Println(buf.String())
	//Output:Hre’sToTheCrzyOns.TheMsfts.TheRbls.TheTroublmkrs.TheRoundPgsInTheSquareHls.
	//TheOnsWhoSeeThngsDffrntly.Thy’reNotFndOfRls.AndThyHveNoRspctForTheSttsQuo.
	//YouCanQuoteThm,DsgreeWthThm,GlrfyOrVlfyThm.
	//AboutTheOnlyThngYouCn’tDoIsIgnreThm.BcauseThyChngeThngs.ThyInvnt.ThyImgne.
	//ThyHeal.ThyExplre.ThyCreate.ThyInspre.ThyPshTheHmnRceFrwrd.
	//MybeThyHveToBeCrzy.
	//HowElseCanYouStreAtAnEmptyCnvsAndSeeAWrkOfArt?OrSitInSlnceAnd
	//HearASngTht’sNvrBeenWrttn?OrGzeAtARedPlntAndSeeALbrtryOnWheels?
	//WeMkeToolsForThseKndsOfPeople.
	//WhleSmeSeeThmAsTheCrzyOns,WeSeeGnius.BcauseThePeopleWhoAreCrzyEnough
	//ToThnkThyCanChngeTheWrld,AreTheOnsWhoDo.
}
