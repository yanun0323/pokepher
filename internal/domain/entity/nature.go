package entity

type Nature string

const (
	NatureHardy   = "Hardy"   //	がんばりや	Attack	Attack	Spicy	Spicy
	NatureLonely  = "Lonely"  //	さみしがり	Attack	Defense	Spicy	Sour
	NatureBrave   = "Brave"   //	ゆうかん	Attack	Speed	Spicy	Sweet
	NatureAdamant = "Adamant" //	いじっぱり	Attack	Sp.Attack	Spicy	Dry
	NatureNaughty = "Naughty" //	やんちゃ	Attack	Sp.Defense	Spicy	Bitter
	NatureBold    = "Bold"    //	ずぶとい	Defense	Attack	Sour	Spicy
	NatureDocile  = "Docile"  //	すなお	Defense	Defense	Sour	Sour
	NatureRelaxed = "Relaxed" //	のんき	Defense	Speed	Sour	Sweet
	NatureImpish  = "Impish"  //	わんぱく	Defense	Sp.Attack	Sour	Dry
	NatureLax     = "Lax"     //	のうてんき	Defense	Sp.Defense	Sour	Bitter
	NatureTimid   = "Timid"   //	おくびょう	Speed	Attack	Sweet	Spicy
	NatureHasty   = "Hasty"   //	せっかち	Speed	Defense	Sweet	Sour
	NatureSerious = "Serious" //	まじめ	Speed	Speed	Sweet	Sweet
	NatureJolly   = "Jolly"   //	ようき	Speed	Sp.Attack	Sweet	Dry
	NatureNaive   = "Naive"   //	むじゃき	Speed	Sp.Defense	Sweet	Bitter
	NatureModest  = "Modest"  //	ひかえめ	Sp.Attack	Attack	Dry	Spicy
	NatureMild    = "Mild"    //	おっとり	Sp.Attack	Defense	Dry	Sour
	NatureQuiet   = "Quiet"   //	れいせい	Sp.Attack	Speed	Dry	Sweet
	NatureBashful = "Bashful" //	てれや	Sp.Attack	Sp.Attack	Dry	Dry
	NatureRash    = "Rash"    //	うっかりや	Sp.Attack	Sp.Defense	Dry	Bitter
	NatureCalm    = "Calm"    //	おだやか	Sp.Defense	Attack	Bitter	Spicy
	NatureGentle  = "Gentle"  //	おとなしい	Sp.Defense	Defense	Bitter	Sour
	NatureSassy   = "Sassy"   //	なまいき	Sp.Defense	Speed	Bitter	Sweet
	NatureCareful = "Careful" //	しんちょう	Sp.Defense	Sp.Attack	Bitter	Dry
	NatureQuirky  = "Quirky"  //	きまぐれ	Sp.Defense	Sp.Defense	Bitter	Bitter

)
