package helper

/*  This file has a static list of cnames which are alphanumeric to corresponding integer ids which
is used by the analytics team. We are using it here to allow conversion of cnames from string to int. */
var CompanyIdsMap = map[string]int64{
	"test1":                     287891538242857427,
	"myoloids-dd":               314234186476768442,
	"hritest123":                401254022864182212,
	"apple":                     494275072483087688,
	"freshdesk":                 990554223172357267,
	"folsomdemob":               702036245060478052,
	"jellyvision":               996793288220532551,
	"cleartripadmin@mindtickle": 363974818936274668,
	"smartlinkb":                969280419316831138,
	"bmsb":                      825398749699523836,
	"mt_test":                   890802478743261579,
	"mtacademy":                 451415051714996401,
	"qualtrics":                 343575950453901740,
	"thoughtspot":               348890974717222612,
	"peopleinteractiveb":        908852177951353344,
	"ie9one":                    337221132758310912,
	"admin_safari_3":            221821124771399409,
	"admin_safari_2":            254436026562855826,
	"admin_safari_1":            743174857105817221,
	"ieee":                      678134801662394246,
	"admin_safari_4":            839911197124284483,
	"testmm":                    378291206731399321,
	"kgdemo":                    814271571139603736,
	"unitrends":                 839769906272631337,
	"ultimateb":                 375010918885781031,
	"olacabsb":                  963193368936388227,
	"admin_chrome_3":            112273575025670002,
	"testkaran2":                140092853672364533,
	"admin_firefox_4":           683930049281589857,
	"admin_firefox_1":           834736655027892896,
	"ie9three":                  998359188841059741,
	"admin_firefox_3":           239663041679100356,
	"admin_firefox_2":           152235760024539677,
	"brainwash":                 207754608577691370,
	"knowlarity":                763842683588196612,
	"testkaran":                 273206927296452558,
	"nvrc":                      594546940074952675,
	"couchbaselearning":         882180622883350078,
	"rocketfuel":                338402558254437924,
	"softwaretesting":           751208959731302389,
	"commonfloor":               803912487647438976,
	"ie9two":                    205473343335545754,
	"fultonfinancial":           333902871628937316,
	"admin_chrome_2":            966543311700961498,
	"admin_chrome_4":            943399964325863252,
	"salesreadiness":            651673917597080285,
	"deepak":                    406179202557222521,
	"admin_ie_2":                958711146867608006,
	"emagine":                   316506558514958339,
	"therubiconproject":         399371581363453823,
	"oauth":                     346175401419301624,
	"quikr":                     172238902985464747,
	"testkaran4":                613821670466141577,
	"sagarnew":                  507917227572190022,
	"_default":                  731120944638013766,
	"appnexus":                  336020228103327367,
	"appdynamicsb":              474588761189442554,
	"cloudera":                  232835838059094170,
	"dmschoolsb":                995504024996246381,
	"benaam":                    662051280733022572,
	"pathshala":                 445102591172843526,
	"aricent":                   545391690164871223,
	"enabliumdemo":              785920062437184805,
	"admin_ie11_4":              553771629758067480,
	"nestle":                    701875211448685033,
	"mani":                      664667779847961005,
	"admin_ff_2":                688951457350177672,
	"testkaran3":                967216693418801724,
	"urbanladder":               954197724013275498,
	"ie8one":                    969488796812136709,
	"kpn":                       208225594765341353,
	"aricentdemo":               726019995920510682,
	"cisco":                     593550852352733882,
	"manishm":                   448490142842783519,
	"freechargeb":               202430144840331693,
	"admin_ie_4":                747891768985100205,
	"ie9four":                   389567140054345175,
	"coolgurukul":               391016550754557939,
	"nagesh":                    140861891603749996,
	"mttestadmin2":              365675746314878901,
	"appdynamics":               625084114870385143,
	"snapdelb":                  572022229562402552,
	"design":                    379897922436969033,
	"mttest1":                   711301732354080226,
	"dabur":                     121750294444005604,
	"education":                 945749868907308876,
	"admin_ie_1":                455381583201598051,
	"avalarab":                  443253468575469122,
	"admin_ie_3":                481481136270461195,
	"cleartrip":                 738529649050033242,
	"startup":                   198025932752459591,
	"shankartesting":            220468745480772758,
	"puma":                      873898907215701340,
	"someemailtest":             287851623411467852,
	"nttest":                    481610685555735126,
	"comstor":                   650122455352324774,
	"robotiqb":                  880204908373876479,
	"housingb":                  883263706286152235,
	"htcampus":                  629585471132944044,
	"q1q1":                      575444721229406511,
	"admin_chrome_1":            831083176252184844,
	"template":                  998564296813147154,
	"avalara1":                  137098118647895403,
	"admin_ie11_1":              851444800761613254,
	"sftest":                    141025704803818007,
	"admin_ie11_3":              639471255199294124,
	"admin_ie11_2":              197309625332784878,
	"rubiconprojectb":           863505456844014799,
	"vfbrands":                  428333592517730441,
	"demotgb":                   524030381843188969,
	"testm":                     140325368400111558,
	"admin_ie10_1":              912203390807738690,
	"admin_ie10_2":              421005713661236749,
	"admin_ie10_3":              263158554997161213,
	"admin_ie10_4":              757880487958072596,
	"bbtheatresb":               360683118471782559,
	"demosales":                 697451902025995519,
	"snapdealb":                 581377358542222393,
	"ddoloids":                  677194838511488122,
	"haygroup":                  249738577227393551,
	"whitbread":                 407563189120759989,
	"yahooharsh":                556912182705868035,
	"vivek":                     939185181536941610,
	"zooba":                     209398479710394282,
	"avalara":                   865048902023352667,
	"iostest":                   846477919265883979,
	"ie8four":                   287772568058512698,
	"htcampusb":                 894289740790484587,
	"mttestadmin":               757541206162930321,
	"admin_ff_3":                880205506142899363,
	"admin_ipad_4":              764189592221395104,
	"admin_ff_1":                676428794208593313,
	"admin_ipad_1":              149737985030601421,
	"admin_ipad_3":              682359735035381073,
	"admin_ipad_2":              142094529239773500,
	"rocketfueluniversity":      299230381295100738,
	"harshal":                   734617892017340790,
	"iossapp":                   887983269630433753,
	"mbrdi":                     372722867451569086,
	"mvf":                       430331407661431963,
	"admin_ff_4":                393262623010288350,
	"couchbasechannels":         807210989322942309,
	"ntnew":                     799982345728701910,
	"ie8three":                  159353597596405455,
	"bookmyshowb":               600373210557268344,
	"ie8two":                    126799992414851359,
	"quickheal":                 420394964085634655,
	"rubiconproject":            530352392526781743}

var OrgsMap = map[string]int64{
	"test1":                     287891538242857427,
	"myoloids-dd":               314234186476768442,
	"hritest123":                401254022864182212,
	"apple":                     494275072483087688,
	"freshdesk":                 990554223172357267,
	"folsomdemob":               702036245060478052,
	"jellyvision":               996793288220532551,
	"cleartripadmin@mindtickle": 363974818936274668,
	"smartlinkb":                969280419316831138,
	"bmsb":                      825398749699523836,
	"mt_test":                   890802478743261579,
	"mtacademy":                 451415051714996401,
	"qualtrics":                 343575950453901740,
	"thoughtspot":               348890974717222612,
	"peopleinteractiveb":        908852177951353344,
	"ie9one":                    337221132758310912,
	"admin_safari_3":            221821124771399409,
	"admin_safari_2":            254436026562855826,
	"admin_safari_1":            743174857105817221,
	"ieee":                      678134801662394246,
	"admin_safari_4":            839911197124284483,
	"testmm":                    378291206731399321,
	"kgdemo":                    814271571139603736,
	"unitrends":                 839769906272631337,
	"ultimateb":                 375010918885781031,
	"olacabsb":                  963193368936388227,
	"admin_chrome_3":            112273575025670002,
	"testkaran2":                140092853672364533,
	"admin_firefox_4":           683930049281589857,
	"admin_firefox_1":           834736655027892896,
	"ie9three":                  998359188841059741,
	"admin_firefox_3":           239663041679100356,
	"admin_firefox_2":           152235760024539677,
	"brainwash":                 207754608577691370,
	"knowlarity":                763842683588196612,
	"testkaran":                 273206927296452558,
	"nvrc":                      594546940074952675,
	"couchbaselearning":         882180622883350078,
	"rocketfuel":                338402558254437924,
	"softwaretesting":           751208959731302389,
	"commonfloor":               803912487647438976,
	"ie9two":                    205473343335545754,
	"fultonfinancial":           333902871628937316,
	"admin_chrome_2":            966543311700961498,
	"admin_chrome_4":            943399964325863252,
	"salesreadiness":            651673917597080285,
	"deepak":                    406179202557222521,
	"admin_ie_2":                958711146867608006,
	"emagine":                   316506558514958339,
	"therubiconproject":         399371581363453823,
	"oauth":                     346175401419301624,
	"quikr":                     172238902985464747,
	"testkaran4":                613821670466141577,
	"sagarnew":                  507917227572190022,
	"_default":                  731120944638013766,
	"appnexus":                  336020228103327367,
	"appdynamicsb":              474588761189442554,
	"cloudera":                  232835838059094170,
	"dmschoolsb":                995504024996246381,
	"benaam":                    662051280733022572,
	"pathshala":                 445102591172843526,
	"aricent":                   545391690164871223,
	"enabliumdemo":              785920062437184805,
	"admin_ie11_4":              553771629758067480,
	"nestle":                    701875211448685033,
	"mani":                      664667779847961005,
	"admin_ff_2":                688951457350177672,
	"testkaran3":                967216693418801724,
	"urbanladder":               954197724013275498,
	"ie8one":                    969488796812136709,
	"kpn":                       208225594765341353,
	"aricentdemo":               726019995920510682,
	"cisco":                     593550852352733882,
	"manishm":                   448490142842783519,
	"freechargeb":               202430144840331693,
	"admin_ie_4":                747891768985100205,
	"ie9four":                   389567140054345175,
	"coolgurukul":               391016550754557939,
	"nagesh":                    140861891603749996,
	"mttestadmin2":              365675746314878901,
	"appdynamics":               625084114870385143,
	"snapdelb":                  572022229562402552,
	"design":                    379897922436969033,
	"mttest1":                   711301732354080226,
	"dabur":                     121750294444005604,
	"education":                 945749868907308876,
	"admin_ie_1":                455381583201598051,
	"avalarab":                  443253468575469122,
	"admin_ie_3":                481481136270461195,
	"cleartrip":                 738529649050033242,
	"startup":                   198025932752459591,
	"shankartesting":            220468745480772758,
	"puma":                      873898907215701340,
	"someemailtest":             287851623411467852,
	"nttest":                    481610685555735126,
	"comstor":                   650122455352324774,
	"robotiqb":                  880204908373876479,
	"housingb":                  883263706286152235,
	"htcampus":                  629585471132944044,
	"q1q1":                      575444721229406511,
	"admin_chrome_1":            831083176252184844,
	"template":                  998564296813147154,
	"avalara1":                  137098118647895403,
	"admin_ie11_1":              851444800761613254,
	"sftest":                    141025704803818007,
	"admin_ie11_3":              639471255199294124,
	"admin_ie11_2":              197309625332784878,
	"rubiconprojectb":           863505456844014799,
	"vfbrands":                  428333592517730441,
	"demotgb":                   524030381843188969,
	"testm":                     140325368400111558,
	"admin_ie10_1":              912203390807738690,
	"admin_ie10_2":              421005713661236749,
	"admin_ie10_3":              263158554997161213,
	"admin_ie10_4":              757880487958072596,
	"bbtheatresb":               360683118471782559,
	"demosales":                 697451902025995519,
	"snapdealb":                 581377358542222393,
	"ddoloids":                  677194838511488122,
	"haygroup":                  249738577227393551,
	"whitbread":                 407563189120759989,
	"yahooharsh":                556912182705868035,
	"vivek":                     939185181536941610,
	"zooba":                     209398479710394282,
	"avalara":                   865048902023352667,
	"iostest":                   846477919265883979,
	"ie8four":                   287772568058512698,
	"htcampusb":                 894289740790484587,
	"mttestadmin":               757541206162930321,
	"admin_ff_3":                880205506142899363,
	"admin_ipad_4":              764189592221395104,
	"admin_ff_1":                676428794208593313,
	"admin_ipad_1":              149737985030601421,
	"admin_ipad_3":              682359735035381073,
	"admin_ipad_2":              142094529239773500,
	"rocketfueluniversity":      299230381295100738,
	"harshal":                   734617892017340790,
	"iossapp":                   887983269630433753,
	"mbrdi":                     372722867451569086,
	"mvf":                       430331407661431963,
	"admin_ff_4":                393262623010288350,
	"couchbasechannels":         807210989322942309,
	"ntnew":                     799982345728701910,
	"ie8three":                  159353597596405455,
	"bookmyshowb":               600373210557268344,
	"ie8two":                    126799992414851359,
	"quickheal":                 420394964085634655,
	"rubiconproject":            530352392526781743}