package waifupicsgo

type Action interface {
	~string
	String() string
}

type SFWAction string
type NSFWAction string

const (
	SFWWaifu    SFWAction = "waifu"
	SFWNeko     SFWAction = "neko"
	SFWShinobu  SFWAction = "shinobu"
	SFWMegumin  SFWAction = "megumin"
	SFWBully    SFWAction = "bully"
	SFWCuddle   SFWAction = "cuddle"
	SFWCry      SFWAction = "cry"
	SFWHug      SFWAction = "hug"
	SFWAwoo     SFWAction = "awoo"
	SFWKiss     SFWAction = "kiss"
	SFWLick     SFWAction = "lick"
	SFWPat      SFWAction = "pat"
	SFWSmug     SFWAction = "smug"
	SFWBink     SFWAction = "bonk"
	SFWYeet     SFWAction = "yeet"
	SFWBlush    SFWAction = "blush"
	SFWSmile    SFWAction = "smile"
	SFWWave     SFWAction = "wave"
	SFWHighfive SFWAction = "highfive"
	SFWHandhold SFWAction = "handhold"
	SFWNom      SFWAction = "nom"
	SFWBite     SFWAction = "bite"
	SFWGlomp    SFWAction = "glomp"
	SFWSlap     SFWAction = "slap"
	SFWKill     SFWAction = "kill"
	SFWKick     SFWAction = "kick"
	SFWHappy    SFWAction = "happy"
	SFWWink     SFWAction = "wink"
	SFWPoke     SFWAction = "poke"
	SFWDance    SFWAction = "dance"
	SFWCringe   SFWAction = "cringe"
)

func (sfw SFWAction) String() string {
	return string(sfw)
}

const (
	NSFWWaifu   NSFWAction = "waifu"
	NSFWNeko    NSFWAction = "neko"
	NSFWTrap    NSFWAction = "trap"
	NSFWBlowjob NSFWAction = "blowjob"
)

func (nsfw NSFWAction) String() string {
	return string(nsfw)
}

type CategoryClient[A Action] struct {
	category string
}

var (
	SFW  CategoryClient[SFWAction]  = CategoryClient[SFWAction]{category: "sfw"}
	NSFW CategoryClient[NSFWAction] = CategoryClient[NSFWAction]{category: "nsfw"}
)

type Image struct {
	URL string `json:"url"`
}
