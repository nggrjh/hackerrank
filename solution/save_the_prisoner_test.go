package solution

import (
	"testing"
)

func Test_saveThePrisoner(t *testing.T) {
	t.Parallel()
	type args struct {
		prisoner int32
		candy    int32
		chair    int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0.0": {
			args: args{
				prisoner: 4,
				candy:    6,
				chair:    2,
			},
			want: 3,
		},
		"case_0.1": {
			args: args{
				prisoner: 352926151,
				candy:    380324688,
				chair:    94730870,
			},
			want: 122129406,
		},
		"case_0.2": {
			args: args{
				prisoner: 94431605,
				candy:    679262176,
				chair:    5284458,
			},
			want: 23525398,
		},
		"case_0.3": {
			args: args{
				prisoner: 208526924,
				candy:    756265725,
				chair:    150817879,
			},
			want: 72975907,
		},
		"case_0.4": {
			args: args{
				prisoner: 962975336,
				candy:    972576181,
				chair:    396355184,
			},
			want: 405956028,
		},
		"case_0.5": {
			args: args{
				prisoner: 464237185,
				candy:    937820284,
				chair:    255816794,
			},
			want: 265162707,
		},
		"case_0.6": {
			args: args{
				prisoner: 649320641,
				candy:    742902564,
				chair:    647542323,
			},
			want: 91803604,
		},
		"case_0.7": {
			args: args{
				prisoner: 174512033,
				candy:    711706897,
				chair:    68977959,
			},
			want: 82636723,
		},
		"case_0.8": {
			args: args{
				prisoner: 107283902,
				candy:    156676511,
				chair:    67149447,
			},
			want: 9258153,
		},
		"case_0.9": {
			args: args{
				prisoner: 104513201,
				candy:    298399273,
				chair:    96292548,
			},
			want: 81152217,
		},
		"case_0.10": {
			args: args{
				prisoner: 113378824,
				candy:    274011418,
				chair:    98103763,
			},
			want: 31978708,
		},
		"case_0.11": {
			args: args{
				prisoner: 934815799,
				candy:    991959468,
				chair:    212396037,
			},
			want: 269539705,
		},
		"case_0.12": {
			args: args{
				prisoner: 88325121,
				candy:    435452998,
				chair:    24617705,
			},
			want: 18445097,
		},
		"case_0.13": {
			args: args{
				prisoner: 984873585,
				candy:    997634055,
				chair:    103050157,
			},
			want: 115810626,
		},
		"case_0.14": {
			args: args{
				prisoner: 344218387,
				candy:    715364875,
				chair:    90658180,
			},
			want: 117586280,
		},
		"case_0.15": {
			args: args{
				prisoner: 556065259,
				candy:    615709967,
				chair:    156417592,
			},
			want: 216062299,
		},
		"case_0.16": {
			args: args{
				prisoner: 57109959,
				candy:    451440582,
				chair:    4188603,
			},
			want: 55859471,
		},
		"case_0.17": {
			args: args{
				prisoner: 353972922,
				candy:    573651462,
				chair:    244520504,
			},
			want: 110226121,
		},
		"case_0.18": {
			args: args{
				prisoner: 946486979,
				candy:    973168361,
				chair:    647886035,
			},
			want: 674567416,
		},
		"case_0.19": {
			args: args{
				prisoner: 368127406,
				candy:    680428368,
				chair:    105517295,
			},
			want: 49690850,
		},
		"case_0.20": {
			args: args{
				prisoner: 884634320,
				candy:    965112925,
				chair:    119757238,
			},
			want: 200235842,
		},
		"case_0.21": {
			args: args{
				prisoner: 422557970,
				candy:    744431033,
				chair:    28932300,
			},
			want: 350805362,
		},
		"case_0.22": {
			args: args{
				prisoner: 634954007,
				candy:    957414623,
				chair:    341366379,
			},
			want: 28872987,
		},
		"case_0.23": {
			args: args{
				prisoner: 190265362,
				candy:    445253893,
				chair:    184632922,
			},
			want: 59090728,
		},
		"case_0.24": {
			args: args{
				prisoner: 293315518,
				candy:    479153471,
				chair:    120684020,
			},
			want: 13206454,
		},
		"case_0.25": {
			args: args{
				prisoner: 72410472,
				candy:    80198999,
				chair:    984948,
			},
			want: 8773474,
		},
		"case_0.26": {
			args: args{
				prisoner: 784893322,
				candy:    849791807,
				chair:    360911386,
			},
			want: 425809870,
		},
		"case_0.27": {
			args: args{
				prisoner: 846191883,
				candy:    880790237,
				chair:    510053756,
			},
			want: 544652109,
		},
		"case_0.28": {
			args: args{
				prisoner: 297201660,
				candy:    812278057,
				chair:    198376746,
			},
			want: 119049822,
		},
		"case_0.29": {
			args: args{
				prisoner: 945087694,
				candy:    999220046,
				chair:    465061514,
			},
			want: 519193865,
		},
		"case_0.30": {
			args: args{
				prisoner: 786859800,
				candy:    831171414,
				chair:    378370933,
			},
			want: 422682546,
		},
		"case_0.31": {
			args: args{
				prisoner: 528402029,
				candy:    859318899,
				chair:    224559950,
			},
			want: 27074790,
		},
		"case_0.32": {
			args: args{
				prisoner: 522640531,
				candy:    755821672,
				chair:    28838424,
			},
			want: 262019564,
		},
		"case_0.33": {
			args: args{
				prisoner: 864006909,
				candy:    879474138,
				chair:    806467486,
			},
			want: 821934714,
		},
		"case_0.34": {
			args: args{
				prisoner: 613544440,
				candy:    943850900,
				chair:    258190755,
			},
			want: 588497214,
		},
		"case_0.35": {
			args: args{
				prisoner: 734228459,
				candy:    928771704,
				chair:    265979283,
			},
			want: 460522527,
		},
		"case_0.36": {
			args: args{
				prisoner: 542812502,
				candy:    597832172,
				chair:    330877768,
			},
			want: 385897437,
		},
		"case_0.37": {
			args: args{
				prisoner: 541133561,
				candy:    748691081,
				chair:    126348492,
			},
			want: 333906011,
		},
		"case_0.38": {
			args: args{
				prisoner: 51187317,
				candy:    524866691,
				chair:    1143193,
			},
			want: 14136713,
		},
		"case_0.39": {
			args: args{
				prisoner: 885290374,
				candy:    990676850,
				chair:    373368385,
			},
			want: 478754860,
		},
		"case_0.40": {
			args: args{
				prisoner: 147955933,
				candy:    450823700,
				chair:    138416059,
			},
			want: 145371959,
		},
		"case_0.41": {
			args: args{
				prisoner: 100046465,
				candy:    587967212,
				chair:    32555061,
			},
			want: 20243482,
		},
		"case_0.42": {
			args: args{
				prisoner: 233926824,
				candy:    996957988,
				chair:    31809378,
			},
			want: 93060069,
		},
		"case_0.43": {
			args: args{
				prisoner: 785405778,
				candy:    835771758,
				chair:    689182705,
			},
			want: 739548684,
		},
		"case_0.44": {
			args: args{
				prisoner: 444389615,
				candy:    870657324,
				chair:    72775880,
			},
			want: 54653973,
		},
		"case_0.45": {
			args: args{
				prisoner: 702580369,
				candy:    895325888,
				chair:    345053199,
			},
			want: 537798717,
		},
		"case_0.46": {
			args: args{
				prisoner: 968559651,
				candy:    974760313,
				chair:    326732084,
			},
			want: 332932745,
		},
		"case_0.47": {
			args: args{
				prisoner: 299437419,
				candy:    514618345,
				chair:    254272806,
			},
			want: 170016312,
		},
		"case_0.48": {
			args: args{
				prisoner: 901702945,
				candy:    930227426,
				chair:    727030891,
			},
			want: 755555371,
		},
		"case_0.49": {
			args: args{
				prisoner: 721843209,
				candy:    736359383,
				chair:    225283784,
			},
			want: 239799957,
		},
		"case_0.50": {
			args: args{
				prisoner: 833018320,
				candy:    883439261,
				chair:    806599246,
			},
			want: 24001866,
		},
		"case_0.51": {
			args: args{
				prisoner: 267346244,
				candy:    307857609,
				chair:    46989880,
			},
			want: 87501244,
		},
		"case_0.52": {
			args: args{
				prisoner: 299901304,
				candy:    892163356,
				chair:    210218436,
			},
			want: 202677879,
		},
		"case_0.53": {
			args: args{
				prisoner: 565637506,
				candy:    795821687,
				chair:    158300457,
			},
			want: 388484637,
		},
		"case_0.54": {
			args: args{
				prisoner: 107336562,
				candy:    781910357,
				chair:    54488503,
			},
			want: 85042925,
		},
		"case_0.55": {
			args: args{
				prisoner: 513281286,
				candy:    916998022,
				chair:    254269425,
			},
			want: 144704874,
		},
		"case_0.56": {
			args: args{
				prisoner: 413431205,
				candy:    611661371,
				chair:    188998419,
			},
			want: 387228584,
		},
		"case_0.57": {
			args: args{
				prisoner: 740163288,
				candy:    938647312,
				chair:    571382392,
			},
			want: 29703127,
		},
		"case_0.58": {
			args: args{
				prisoner: 44702121,
				candy:    296589002,
				chair:    43470596,
			},
			want: 27144750,
		},
		"case_0.59": {
			args: args{
				prisoner: 771733011,
				candy:    919327188,
				chair:    317638907,
			},
			want: 465233083,
		},
		"case_0.60": {
			args: args{
				prisoner: 718860003,
				candy:    815844769,
				chair:    495144331,
			},
			want: 592129096,
		},
		"case_0.61": {
			args: args{
				prisoner: 377975600,
				candy:    438513404,
				chair:    111085209,
			},
			want: 171623012,
		},
		"case_0.62": {
			args: args{
				prisoner: 424965480,
				candy:    928959619,
				chair:    20776133,
			},
			want: 99804791,
		},
		"case_0.63": {
			args: args{
				prisoner: 234986523,
				candy:    732169039,
				chair:    205952749,
			},
			want: 233162218,
		},
		"case_0.64": {
			args: args{
				prisoner: 377078343,
				candy:    981597420,
				chair:    219264561,
			},
			want: 69626951,
		},
		"case_0.65": {
			args: args{
				prisoner: 612269027,
				candy:    791414524,
				chair:    580170106,
			},
			want: 147046575,
		},
		"case_0.66": {
			args: args{
				prisoner: 232336090,
				candy:    616084008,
				chair:    81392003,
			},
			want: 467740,
		},
		"case_0.67": {
			args: args{
				prisoner: 75059328,
				candy:    274029861,
				chair:    53524881,
			},
			want: 27317429,
		},
		"case_0.68": {
			args: args{
				prisoner: 239121359,
				candy:    646856043,
				chair:    141349731,
			},
			want: 70841696,
		},
		"case_0.69": {
			args: args{
				prisoner: 923193147,
				candy:    943368157,
				chair:    206717532,
			},
			want: 226892541,
		},
		"case_0.70": {
			args: args{
				prisoner: 12565064,
				candy:    536852817,
				chair:    11557940,
			},
			want: 8113004,
		},
		"case_0.71": {
			args: args{
				prisoner: 360225746,
				candy:    970375527,
				chair:    284883902,
			},
			want: 174582190,
		},
		"case_0.72": {
			args: args{
				prisoner: 213705306,
				candy:    380885426,
				chair:    14495860,
			},
			want: 181675979,
		},
		"case_0.73": {
			args: args{
				prisoner: 659446919,
				candy:    699401237,
				chair:    73837176,
			},
			want: 113791493,
		},
		"case_0.74": {
			args: args{
				prisoner: 335372713,
				candy:    785363124,
				chair:    7610828,
			},
			want: 122228525,
		},
		"case_0.75": {
			args: args{
				prisoner: 538388654,
				candy:    859196325,
				chair:    110284314,
			},
			want: 431091984,
		},
		"case_0.76": {
			args: args{
				prisoner: 118558760,
				candy:    713483972,
				chair:    83950807,
			},
			want: 86082218,
		},
		"case_0.77": {
			args: args{
				prisoner: 896959032,
				candy:    961368580,
				chair:    8848444,
			},
			want: 73257991,
		},
		"case_0.78": {
			args: args{
				prisoner: 25543240,
				candy:    521123082,
				chair:    2472730,
			},
			want: 12731011,
		},
		"case_0.79": {
			args: args{
				prisoner: 258530682,
				candy:    935834352,
				chair:    194732411,
			},
			want: 96444034,
		},
		"case_0.80": {
			args: args{
				prisoner: 465248213,
				candy:    679599042,
				chair:    334563499,
			},
			want: 83666114,
		},
		"case_0.81": {
			args: args{
				prisoner: 331230504,
				candy:    637771661,
				chair:    76778140,
			},
			want: 52088792,
		},
		"case_0.82": {
			args: args{
				prisoner: 976340152,
				candy:    988657462,
				chair:    243958260,
			},
			want: 256275569,
		},
		"case_0.83": {
			args: args{
				prisoner: 552994811,
				candy:    922743535,
				chair:    540135280,
			},
			want: 356889192,
		},
		"case_0.84": {
			args: args{
				prisoner: 626831986,
				candy:    839281366,
				chair:    24222267,
			},
			want: 236671646,
		},
		"case_0.85": {
			args: args{
				prisoner: 157704591,
				candy:    253731033,
				chair:    59023773,
			},
			want: 155050214,
		},
		"case_0.86": {
			args: args{
				prisoner: 806377559,
				candy:    859228114,
				chair:    238044289,
			},
			want: 290894843,
		},
		"case_0.87": {
			args: args{
				prisoner: 838798445,
				candy:    996851254,
				chair:    268459446,
			},
			want: 426512254,
		},
		"case_0.88": {
			args: args{
				prisoner: 847646888,
				candy:    928001503,
				chair:    755190846,
			},
			want: 835545460,
		},
		"case_0.89": {
			args: args{
				prisoner: 877625009,
				candy:    999275937,
				chair:    874127074,
			},
			want: 118152992,
		},
		"case_0.90": {
			args: args{
				prisoner: 847949466,
				candy:    893343194,
				chair:    10497830,
			},
			want: 55891557,
		},
		"case_0.91": {
			args: args{
				prisoner: 35029316,
				candy:    784675240,
				chair:    8200127,
			},
			want: 22230414,
		},
		"case_0.92": {
			args: args{
				prisoner: 111807455,
				candy:    690309882,
				chair:    23190325,
			},
			want: 42655476,
		},
		"case_0.93": {
			args: args{
				prisoner: 355765714,
				candy:    554560093,
				chair:    311565654,
			},
			want: 154594318,
		},
		"case_0.94": {
			args: args{
				prisoner: 1890615,
				candy:    614626804,
				chair:    976253,
			},
			want: 1153181,
		},
		"case_0.95": {
			args: args{
				prisoner: 132293206,
				candy:    165429783,
				chair:    65360680,
			},
			want: 98497256,
		},
		"case_0.96": {
			args: args{
				prisoner: 506726160,
				candy:    934170235,
				chair:    455394293,
			},
			want: 376112207,
		},
		"case_0.97": {
			args: args{
				prisoner: 210041918,
				candy:    328800789,
				chair:    159203369,
			},
			want: 67920321,
		},
		"case_0.98": {
			args: args{
				prisoner: 499999999,
				candy:    999999997,
				chair:    2,
			},
			want: 499999999,
		},
		"case_0.99": {
			args: args{
				prisoner: 499999999,
				candy:    999999998,
				chair:    2,
			},
			want: 1,
		},
		"case_0.100": {
			args: args{
				prisoner: 999999999,
				candy:    999999999,
				chair:    1,
			},
			want: 999999999,
		},
		"case_4.1": {
			args: args{
				prisoner: 5,
				candy:    2,
				chair:    1,
			},
			want: 2,
		},
		"case_4.2": {
			args: args{
				prisoner: 5,
				candy:    2,
				chair:    2,
			},
			want: 3,
		},
		"case_10.1": {
			args: args{
				prisoner: 46934,
				candy:    543563655,
				chair:    46743,
			},
			want: 20809,
		},
		"case_10.2": {
			args: args{
				prisoner: 530,
				candy:    533048047,
				chair:    529,
			},
			want: 15,
		},
		"case_10.3": {
			args: args{
				prisoner: 436776012,
				candy:    436776012,
				chair:    436776011,
			},
			want: 436776010,
		},
		"case_10.4": {
			args: args{
				prisoner: 999999999,
				candy:    999999998,
				chair:    999983945,
			},
			want: 999983943,
		},
		"case_10.5": {
			args: args{
				prisoner: 59,
				candy:    78693934,
				chair:    2,
			},
			want: 30,
		},
		"case_10.6": {
			args: args{
				prisoner: 49,
				candy:    897910613,
				chair:    48,
			},
			want: 17,
		},
		"case_10.7": {
			args: args{
				prisoner: 124,
				candy:    738996353,
				chair:    2,
			},
			want: 2,
		},
		"case_10.8": {
			args: args{
				prisoner: 999999999,
				candy:    871077789,
				chair:    999999998,
			},
			want: 871077787,
		},
		"case_10.9": {
			args: args{
				prisoner: 2,
				candy:    576581,
				chair:    1,
			},
			want: 1,
		},
		"case_10.10": {
			args: args{
				prisoner: 526663404,
				candy:    801992888,
				chair:    526663403,
			},
			want: 275329482,
		},
		"case_10.11": {
			args: args{
				prisoner: 999999998,
				candy:    999999998,
				chair:    1,
			},
			want: 999999998,
		},
		"case_10.12": {
			args: args{
				prisoner: 126,
				candy:    859530642,
				chair:    126,
			},
			want: 95,
		},
		"case_10.13": {
			args: args{
				prisoner: 1000000000,
				candy:    999999999,
				chair:    1000000000,
			},
			want: 999999998,
		},
		"case_10.14": {
			args: args{
				prisoner: 107,
				candy:    425601402,
				chair:    2,
			},
			want: 22,
		},
		"case_10.15": {
			args: args{
				prisoner: 381,
				candy:    695699141,
				chair:    380,
			},
			want: 378,
		},
		"case_10.16": {
			args: args{
				prisoner: 11,
				candy:    32020900,
				chair:    6,
			},
			want: 4,
		},
		"case_10.17": {
			args: args{
				prisoner: 468840391,
				candy:    468840391,
				chair:    1,
			},
			want: 468840391,
		},
		"case_10.18": {
			args: args{
				prisoner: 999999999,
				candy:    29010,
				chair:    1,
			},
			want: 29010,
		},
		"case_10.19": {
			args: args{
				prisoner: 31,
				candy:    238250965,
				chair:    2,
			},
			want: 1,
		},
		"case_10.20": {
			args: args{
				prisoner: 6,
				candy:    923562791,
				chair:    1,
			},
			want: 5,
		},
		"case_10.21": {
			args: args{
				prisoner: 39,
				candy:    558119524,
				chair:    38,
			},
			want: 38,
		},
		"case_10.22": {
			args: args{
				prisoner: 121,
				candy:    652798629,
				chair:    1,
			},
			want: 120,
		},
		"case_10.23": {
			args: args{
				prisoner: 94,
				candy:    105224796,
				chair:    94,
			},
			want: 67,
		},
		"case_10.24": {
			args: args{
				prisoner: 9,
				candy:    903414482,
				chair:    5,
			},
			want: 3,
		},
		"case_10.25": {
			args: args{
				prisoner: 1718761,
				candy:    828441828,
				chair:    1718761,
			},
			want: 1717786,
		},
		"case_10.26": {
			args: args{
				prisoner: 4970962,
				candy:    984250547,
				chair:    4970961,
			},
			want: 69,
		},
		"case_10.27": {
			args: args{
				prisoner: 19,
				candy:    235344290,
				chair:    2,
			},
			want: 12,
		},
		"case_10.28": {
			args: args{
				prisoner: 514824323,
				candy:    514824324,
				chair:    514824323,
			},
			want: 514824323,
		},
		"case_10.29": {
			args: args{
				prisoner: 181,
				candy:    511813156,
				chair:    180,
			},
			want: 178,
		},
		"case_10.30": {
			args: args{
				prisoner: 66,
				candy:    810757794,
				chair:    2,
			},
			want: 1,
		},
		"case_10.31": {
			args: args{
				prisoner: 154,
				candy:    935852917,
				chair:    154,
			},
			want: 152,
		},
		"case_10.32": {
			args: args{
				prisoner: 1000000000,
				candy:    999999999,
				chair:    999974361,
			},
			want: 999974359,
		},
		"case_10.33": {
			args: args{
				prisoner: 21,
				candy:    603073253,
				chair:    20,
			},
			want: 18,
		},
		"case_10.34": {
			args: args{
				prisoner: 29,
				candy:    834017184,
				chair:    28,
			},
			want: 5,
		},
		"case_10.35": {
			args: args{
				prisoner: 195446094,
				candy:    586338283,
				chair:    195446093,
			},
			want: 195446093,
		},
		"case_10.36": {
			args: args{
				prisoner: 93995,
				candy:    173193482,
				chair:    93995,
			},
			want: 54691,
		},
		"case_10.37": {
			args: args{
				prisoner: 101,
				candy:    143467773,
				chair:    101,
			},
			want: 100,
		},
		"case_10.38": {
			args: args{
				prisoner: 134,
				candy:    677010612,
				chair:    134,
			},
			want: 133,
		},
		"case_10.39": {
			args: args{
				prisoner: 99,
				candy:    741806010,
				chair:    2,
			},
			want: 1,
		},
		"case_10.40": {
			args: args{
				prisoner: 75,
				candy:    129103876,
				chair:    2,
			},
			want: 2,
		},
		"case_10.41": {
			args: args{
				prisoner: 689371544,
				candy:    689371544,
				chair:    689370115,
			},
			want: 689370114,
		},
		"case_10.42": {
			args: args{
				prisoner: 28410362,
				candy:    340924345,
				chair:    22721112,
			},
			want: 22721112,
		},
		"case_10.43": {
			args: args{
				prisoner: 170,
				candy:    780653100,
				chair:    170,
			},
			want: 9,
		},
		"case_10.44": {
			args: args{
				prisoner: 193,
				candy:    945602138,
				chair:    192,
			},
			want: 180,
		},
		"case_10.45": {
			args: args{
				prisoner: 96,
				candy:    23494832,
				chair:    95,
			},
			want: 78,
		},
		"case_10.46": {
			args: args{
				prisoner: 944675683,
				candy:    944675683,
				chair:    20312,
			},
			want: 20311,
		},
		"case_10.47": {
			args: args{
				prisoner: 27,
				candy:    546238476,
				chair:    26,
			},
			want: 16,
		},
		"case_10.48": {
			args: args{
				prisoner: 76195990,
				candy:    223258463,
				chair:    489,
			},
			want: 70866971,
		},
		"case_10.49": {
			args: args{
				prisoner: 999999999,
				candy:    269208525,
				chair:    1,
			},
			want: 269208525,
		},
		"case_10.50": {
			args: args{
				prisoner: 108,
				candy:    280122192,
				chair:    108,
			},
			want: 107,
		},
		"case_10.51": {
			args: args{
				prisoner: 16,
				candy:    995404080,
				chair:    15,
			},
			want: 14,
		},
		"case_10.52": {
			args: args{
				prisoner: 50158215,
				candy:    451423257,
				chair:    50151646,
			},
			want: 50150967,
		},
		"case_10.53": {
			args: args{
				prisoner: 4,
				candy:    467711281,
				chair:    4,
			},
			want: 4,
		},
		"case_10.54": {
			args: args{
				prisoner: 145,
				candy:    71654651,
				chair:    144,
			},
			want: 144,
		},
		"case_10.55": {
			args: args{
				prisoner: 1000000000,
				candy:    2985,
				chair:    1,
			},
			want: 2985,
		},
		"case_10.56": {
			args: args{
				prisoner: 990301380,
				candy:    1,
				chair:    990271854,
			},
			want: 990271854,
		},
		"case_10.57": {
			args: args{
				prisoner: 999999999,
				candy:    6413,
				chair:    21476,
			},
			want: 27888,
		},
		"case_10.58": {
			args: args{
				prisoner: 2,
				candy:    468939243,
				chair:    1,
			},
			want: 1,
		},
		"case_10.59": {
			args: args{
				prisoner: 399,
				candy:    592025825,
				chair:    398,
			},
			want: 396,
		},
		"case_10.60": {
			args: args{
				prisoner: 8,
				candy:    666688807,
				chair:    8,
			},
			want: 6,
		},
		"case_10.61": {
			args: args{
				prisoner: 7,
				candy:    633100633,
				chair:    2,
			},
			want: 5,
		},
		"case_10.62": {
			args: args{
				prisoner: 12,
				candy:    124444631,
				chair:    7,
			},
			want: 5,
		},
		"case_10.63": {
			args: args{
				prisoner: 8,
				candy:    347412080,
				chair:    5,
			},
			want: 4,
		},
		"case_10.64": {
			args: args{
				prisoner: 999999999,
				candy:    1,
				chair:    999999998,
			},
			want: 999999998,
		},
		"case_10.65": {
			args: args{
				prisoner: 42774012,
				candy:    765544482,
				chair:    2,
			},
			want: 38386279,
		},
		"case_10.66": {
			args: args{
				prisoner: 18,
				candy:    359622755,
				chair:    18,
			},
			want: 16,
		},
		"case_10.67": {
			args: args{
				prisoner: 198,
				candy:    964246139,
				chair:    2,
			},
			want: 198,
		},
		"case_10.68": {
			args: args{
				prisoner: 999999999,
				candy:    999999998,
				chair:    999999999,
			},
			want: 999999997,
		},
		"case_10.69": {
			args: args{
				prisoner: 10,
				candy:    143638249,
				chair:    1,
			},
			want: 9,
		},
		"case_10.70": {
			args: args{
				prisoner: 1946080,
				candy:    978878239,
				chair:    2,
			},
			want: 1946080,
		},
		"case_10.71": {
			args: args{
				prisoner: 1000000000,
				candy:    999976501,
				chair:    999990588,
			},
			want: 999967088,
		},
		"case_10.72": {
			args: args{
				prisoner: 999999999,
				candy:    999978713,
				chair:    28209,
			},
			want: 6922,
		},
		"case_10.73": {
			args: args{
				prisoner: 433677591,
				candy:    433663369,
				chair:    206662538,
			},
			want: 206648315,
		},
		"case_10.74": {
			args: args{
				prisoner: 999999999,
				candy:    1,
				chair:    1,
			},
			want: 1,
		},
		"case_10.75": {
			args: args{
				prisoner: 2325,
				candy:    562408200,
				chair:    2,
			},
			want: 1,
		},
		"case_10.76": {
			args: args{
				prisoner: 172,
				candy:    456632596,
				chair:    171,
			},
			want: 114,
		},
		"case_10.77": {
			args: args{
				prisoner: 19,
				candy:    563815520,
				chair:    11,
			},
			want: 11,
		},
		"case_10.78": {
			args: args{
				prisoner: 34339,
				candy:    656699084,
				chair:    101,
			},
			want: 148,
		},
		"case_10.79": {
			args: args{
				prisoner: 1000000000,
				candy:    999997154,
				chair:    999999999,
			},
			want: 999997152,
		},
		"case_10.80": {
			args: args{
				prisoner: 1000000000,
				candy:    1000000000,
				chair:    90143095,
			},
			want: 90143094,
		},
		"case_10.81": {
			args: args{
				prisoner: 2,
				candy:    213164653,
				chair:    1,
			},
			want: 1,
		},
		"case_10.82": {
			args: args{
				prisoner: 134,
				candy:    644278309,
				chair:    113,
			},
			want: 123,
		},
		"case_10.83": {
			args: args{
				prisoner: 1000000000,
				candy:    640282835,
				chair:    2,
			},
			want: 640282836,
		},
		"case_10.84": {
			args: args{
				prisoner: 1000000000,
				candy:    1000000000,
				chair:    999999999,
			},
			want: 999999998,
		},
		"case_10.85": {
			args: args{
				prisoner: 999999999,
				candy:    999999999,
				chair:    999999999,
			},
			want: 999999998,
		},
		"case_10.86": {
			args: args{
				prisoner: 999999999,
				candy:    11132,
				chair:    999999998,
			},
			want: 11130,
		},
		"case_10.87": {
			args: args{
				prisoner: 197,
				candy:    190791557,
				chair:    197,
			},
			want: 11,
		},
		"case_10.88": {
			args: args{
				prisoner: 1000000000,
				candy:    1000000000,
				chair:    1000000000,
			},
			want: 999999999,
		},
		"case_10.89": {
			args: args{
				prisoner: 46,
				candy:    56740430,
				chair:    45,
			},
			want: 26,
		},
		"case_10.90": {
			args: args{
				prisoner: 40,
				candy:    277585960,
				chair:    1,
			},
			want: 40,
		},
		"case_10.91": {
			args: args{
				prisoner: 56,
				candy:    306549319,
				chair:    56,
			},
			want: 54,
		},
		"case_10.92": {
			args: args{
				prisoner: 62,
				candy:    803079454,
				chair:    43,
			},
			want: 6,
		},
		"case_10.93": {
			args: args{
				prisoner: 184,
				candy:    834149464,
				chair:    184,
			},
			want: 183,
		},
		"case_10.94": {
			args: args{
				prisoner: 9,
				candy:    526219551,
				chair:    9,
			},
			want: 8,
		},
		"case_10.95": {
			args: args{
				prisoner: 999999999,
				candy:    999999999,
				chair:    583101931,
			},
			want: 583101930,
		},
		"case_10.96": {
			args: args{
				prisoner: 1000000000,
				candy:    999999999,
				chair:    999999999,
			},
			want: 999999997,
		},
		"case_10.97": {
			args: args{
				prisoner: 176,
				candy:    719643761,
				chair:    1,
			},
			want: 1,
		},
		"case_10.98": {
			args: args{
				prisoner: 1000000000,
				candy:    999999999,
				chair:    328966243,
			},
			want: 328966241,
		},
		"case_10.99": {
			args: args{
				prisoner: 65,
				candy:    980609890,
				chair:    37,
			},
			want: 36,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			if got := saveThePrisoner(tt.args.prisoner, tt.args.candy, tt.args.chair); got != tt.want {
				t.Errorf("saveThePrisoner() = %v, want %v", got, tt.want)
			}
		})
	}
}
