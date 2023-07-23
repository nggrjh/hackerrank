package solution

import (
	"testing"
)

func Test_saveThePrisoner(t *testing.T) {
	t.Parallel()
	type args struct {
		prisoner int32
		candy    int32
		start    int32
	}
	tests := map[string]struct {
		args args
		want int32
	}{
		"case_0.0": {
			args: args{
				prisoner: 4,
				candy:    6,
				start:    2,
			},
			want: 3,
		},
		// "case_0.1": {
		// 	args: args{
		// 		prisoner: 352926151,
		// 		candy:    380324688,
		// 		start:    94730870,
		// 	},
		// 	want: 122129406,
		// },
		"case_0.2": {
			args: args{
				prisoner: 94431605,
				candy:    679262176,
				start:    5284458,
			},
			want: 23525398,
		},
		// "case_0.3": {
		// 	args: args{
		// 		prisoner: 208526924,
		// 		candy:    756265725,
		// 		start:    150817879,
		// 	},
		// 	want: 72975907,
		// },
		"case_0.4": {
			args: args{
				prisoner: 962975336,
				candy:    972576181,
				start:    396355184,
			},
			want: 405956028,
		},
		"case_0.5": {
			args: args{
				prisoner: 464237185,
				candy:    937820284,
				start:    255816794,
			},
			want: 265162707,
		},
		// "case_0.6": {
		// 	args: args{
		// 		prisoner: 649320641,
		// 		candy:    742902564,
		// 		start:    647542323,
		// 	},
		// 	want: 91803604,
		// },
		"case_0.7": {
			args: args{
				prisoner: 174512033,
				candy:    711706897,
				start:    68977959,
			},
			want: 82636723,
		},
		// "case_0.8": {
		// 	args: args{
		// 		prisoner: 107283902,
		// 		candy:    156676511,
		// 		start:    67149447,
		// 	},
		// 	want: 9258153,
		// },
		// "case_0.9": {
		// 	args: args{
		// 		prisoner: 104513201,
		// 		candy:    298399273,
		// 		start:    96292548,
		// 	},
		// 	want: 81152217,
		// },
		// "case_0.10": {
		// 	args: args{
		// 		prisoner: 113378824,
		// 		candy:    274011418,
		// 		start:    98103763,
		// 	},
		// 	want: 31978708,
		// },
		"case_0.11": {
			args: args{
				prisoner: 934815799,
				candy:    991959468,
				start:    212396037,
			},
			want: 269539705,
		},
		// "case_0.12": {
		// 	args: args{
		// 		prisoner: 88325121,
		// 		candy:    435452998,
		// 		start:    24617705,
		// 	},
		// 	want: 18445097,
		// },
		"case_0.13": {
			args: args{
				prisoner: 984873585,
				candy:    997634055,
				start:    103050157,
			},
			want: 115810626,
		},
		"case_0.14": {
			args: args{
				prisoner: 344218387,
				candy:    715364875,
				start:    90658180,
			},
			want: 117586280,
		},
		"case_0.15": {
			args: args{
				prisoner: 556065259,
				candy:    615709967,
				start:    156417592,
			},
			want: 216062299,
		},
		"case_0.16": {
			args: args{
				prisoner: 57109959,
				candy:    451440582,
				start:    4188603,
			},
			want: 55859471,
		},
		// "case_0.17": {
		// 	args: args{
		// 		prisoner: 353972922,
		// 		candy:    573651462,
		// 		start:    244520504,
		// 	},
		// 	want: 110226121,
		// },
		"case_0.18": {
			args: args{
				prisoner: 946486979,
				candy:    973168361,
				start:    647886035,
			},
			want: 674567416,
		},
		// "case_0.19": {
		// 	args: args{
		// 		prisoner: 368127406,
		// 		candy:    680428368,
		// 		start:    105517295,
		// 	},
		// 	want: 49690850,
		// },
		"case_0.20": {
			args: args{
				prisoner: 884634320,
				candy:    965112925,
				start:    119757238,
			},
			want: 200235842,
		},
		"case_0.21": {
			args: args{
				prisoner: 422557970,
				candy:    744431033,
				start:    28932300,
			},
			want: 350805362,
		},
		// "case_0.22": {
		// 	args: args{
		// 		prisoner: 634954007,
		// 		candy:    957414623,
		// 		start:    341366379,
		// 	},
		// 	want: 28872987,
		// },
		// "case_0.23": {
		// 	args: args{
		// 		prisoner: 190265362,
		// 		candy:    445253893,
		// 		start:    184632922,
		// 	},
		// 	want: 59090728,
		// },
		// "case_0.24": {
		// 	args: args{
		// 		prisoner: 293315518,
		// 		candy:    479153471,
		// 		start:    120684020,
		// 	},
		// 	want: 13206454,
		// },
		"case_0.25": {
			args: args{
				prisoner: 72410472,
				candy:    80198999,
				start:    984948,
			},
			want: 8773474,
		},
		"case_0.26": {
			args: args{
				prisoner: 784893322,
				candy:    849791807,
				start:    360911386,
			},
			want: 425809870,
		},
		"case_0.27": {
			args: args{
				prisoner: 846191883,
				candy:    880790237,
				start:    510053756,
			},
			want: 544652109,
		},
		// "case_0.28": {
		// 	args: args{
		// 		prisoner: 297201660,
		// 		candy:    812278057,
		// 		start:    198376746,
		// 	},
		// 	want: 119049822,
		// },
		"case_0.29": {
			args: args{
				prisoner: 945087694,
				candy:    999220046,
				start:    465061514,
			},
			want: 519193865,
		},
		"case_0.30": {
			args: args{
				prisoner: 786859800,
				candy:    831171414,
				start:    378370933,
			},
			want: 422682546,
		},
		// "case_0.31": {
		// 	args: args{
		// 		prisoner: 528402029,
		// 		candy:    859318899,
		// 		start:    224559950,
		// 	},
		// 	want: 27074790,
		// },
		"case_0.32": {
			args: args{
				prisoner: 522640531,
				candy:    755821672,
				start:    28838424,
			},
			want: 262019564,
		},
		"case_0.33": {
			args: args{
				prisoner: 864006909,
				candy:    879474138,
				start:    806467486,
			},
			want: 821934714,
		},
		"case_0.34": {
			args: args{
				prisoner: 613544440,
				candy:    943850900,
				start:    258190755,
			},
			want: 588497214,
		},
		"case_0.35": {
			args: args{
				prisoner: 734228459,
				candy:    928771704,
				start:    265979283,
			},
			want: 460522527,
		},
		"case_0.36": {
			args: args{
				prisoner: 542812502,
				candy:    597832172,
				start:    330877768,
			},
			want: 385897437,
		},
		"case_0.37": {
			args: args{
				prisoner: 541133561,
				candy:    748691081,
				start:    126348492,
			},
			want: 333906011,
		},
		"case_0.38": {
			args: args{
				prisoner: 51187317,
				candy:    524866691,
				start:    1143193,
			},
			want: 14136713,
		},
		"case_0.39": {
			args: args{
				prisoner: 885290374,
				candy:    990676850,
				start:    373368385,
			},
			want: 478754860,
		},
		"case_0.40": {
			args: args{
				prisoner: 147955933,
				candy:    450823700,
				start:    138416059,
			},
			want: 145371959,
		},
		// "case_0.41": {
		// 	args: args{
		// 		prisoner: 100046465,
		// 		candy:    587967212,
		// 		start:    32555061,
		// 	},
		// 	want: 20243482,
		// },
		"case_0.42": {
			args: args{
				prisoner: 233926824,
				candy:    996957988,
				start:    31809378,
			},
			want: 93060069,
		},
		"case_0.43": {
			args: args{
				prisoner: 785405778,
				candy:    835771758,
				start:    689182705,
			},
			want: 739548684,
		},
		// "case_0.44": {
		// 	args: args{
		// 		prisoner: 444389615,
		// 		candy:    870657324,
		// 		start:    72775880,
		// 	},
		// 	want: 54653973,
		// },
		"case_0.45": {
			args: args{
				prisoner: 702580369,
				candy:    895325888,
				start:    345053199,
			},
			want: 537798717,
		},
		"case_0.46": {
			args: args{
				prisoner: 968559651,
				candy:    974760313,
				start:    326732084,
			},
			want: 332932745,
		},
		// "case_0.47": {
		// 	args: args{
		// 		prisoner: 299437419,
		// 		candy:    514618345,
		// 		start:    254272806,
		// 	},
		// 	want: 170016312,
		// },
		"case_0.48": {
			args: args{
				prisoner: 901702945,
				candy:    930227426,
				start:    727030891,
			},
			want: 755555371,
		},
		"case_0.49": {
			args: args{
				prisoner: 721843209,
				candy:    736359383,
				start:    225283784,
			},
			want: 239799957,
		},
		// "case_0.50": {
		// 	args: args{
		// 		prisoner: 833018320,
		// 		candy:    883439261,
		// 		start:    806599246,
		// 	},
		// 	want: 24001866,
		// },
		"case_0.51": {
			args: args{
				prisoner: 267346244,
				candy:    307857609,
				start:    46989880,
			},
			want: 87501244,
		},
		// "case_0.52": {
		// 	args: args{
		// 		prisoner: 299901304,
		// 		candy:    892163356,
		// 		start:    210218436,
		// 	},
		// 	want: 202677879,
		// },
		"case_0.53": {
			args: args{
				prisoner: 565637506,
				candy:    795821687,
				start:    158300457,
			},
			want: 388484637,
		},
		"case_0.54": {
			args: args{
				prisoner: 107336562,
				candy:    781910357,
				start:    54488503,
			},
			want: 85042925,
		},
		// "case_0.55": {
		// 	args: args{
		// 		prisoner: 513281286,
		// 		candy:    916998022,
		// 		start:    254269425,
		// 	},
		// 	want: 144704874,
		// },
		"case_0.56": {
			args: args{
				prisoner: 413431205,
				candy:    611661371,
				start:    188998419,
			},
			want: 387228584,
		},
		// "case_0.57": {
		// 	args: args{
		// 		prisoner: 740163288,
		// 		candy:    938647312,
		// 		start:    571382392,
		// 	},
		// 	want: 29703127,
		// },
		// "case_0.58": {
		// 	args: args{
		// 		prisoner: 44702121,
		// 		candy:    296589002,
		// 		start:    43470596,
		// 	},
		// 	want: 27144750,
		// },
		"case_0.59": {
			args: args{
				prisoner: 771733011,
				candy:    919327188,
				start:    317638907,
			},
			want: 465233083,
		},
		"case_0.60": {
			args: args{
				prisoner: 718860003,
				candy:    815844769,
				start:    495144331,
			},
			want: 592129096,
		},
		"case_0.61": {
			args: args{
				prisoner: 377975600,
				candy:    438513404,
				start:    111085209,
			},
			want: 171623012,
		},
		"case_0.62": {
			args: args{
				prisoner: 424965480,
				candy:    928959619,
				start:    20776133,
			},
			want: 99804791,
		},
		"case_0.63": {
			args: args{
				prisoner: 234986523,
				candy:    732169039,
				start:    205952749,
			},
			want: 233162218,
		},
		// "case_0.64": {
		// 	args: args{
		// 		prisoner: 377078343,
		// 		candy:    981597420,
		// 		start:    219264561,
		// 	},
		// 	want: 69626951,
		// },
		// "case_0.65": {
		// 	args: args{
		// 		prisoner: 612269027,
		// 		candy:    791414524,
		// 		start:    580170106,
		// 	},
		// 	want: 147046575,
		// },
		// "case_0.66": {
		// 	args: args{
		// 		prisoner: 232336090,
		// 		candy:    616084008,
		// 		start:    81392003,
		// 	},
		// 	want: 467740,
		// },
		// "case_0.67": {
		// 	args: args{
		// 		prisoner: 75059328,
		// 		candy:    274029861,
		// 		start:    53524881,
		// 	},
		// 	want: 27317429,
		// },
		// "case_0.68": {
		// 	args: args{
		// 		prisoner: 239121359,
		// 		candy:    646856043,
		// 		start:    141349731,
		// 	},
		// 	want: 70841696,
		// },
		"case_0.69": {
			args: args{
				prisoner: 923193147,
				candy:    943368157,
				start:    206717532,
			},
			want: 226892541,
		},
		// "case_0.70": {
		// 	args: args{
		// 		prisoner: 12565064,
		// 		candy:    536852817,
		// 		start:    11557940,
		// 	},
		// 	want: 8113004,
		// },
		// "case_0.71": {
		// 	args: args{
		// 		prisoner: 360225746,
		// 		candy:    970375527,
		// 		start:    284883902,
		// 	},
		// 	want: 174582190,
		// },
		"case_0.72": {
			args: args{
				prisoner: 213705306,
				candy:    380885426,
				start:    14495860,
			},
			want: 181675979,
		},
		"case_0.73": {
			args: args{
				prisoner: 659446919,
				candy:    699401237,
				start:    73837176,
			},
			want: 113791493,
		},
		"case_0.74": {
			args: args{
				prisoner: 335372713,
				candy:    785363124,
				start:    7610828,
			},
			want: 122228525,
		},
		"case_0.75": {
			args: args{
				prisoner: 538388654,
				candy:    859196325,
				start:    110284314,
			},
			want: 431091984,
		},
		"case_0.76": {
			args: args{
				prisoner: 118558760,
				candy:    713483972,
				start:    83950807,
			},
			want: 86082218,
		},
		"case_0.77": {
			args: args{
				prisoner: 896959032,
				candy:    961368580,
				start:    8848444,
			},
			want: 73257991,
		},
		"case_0.78": {
			args: args{
				prisoner: 25543240,
				candy:    521123082,
				start:    2472730,
			},
			want: 12731011,
		},
		// "case_0.79": {
		// 	args: args{
		// 		prisoner: 258530682,
		// 		candy:    935834352,
		// 		start:    194732411,
		// 	},
		// 	want: 96444034,
		// },
		// "case_0.80": {
		// 	args: args{
		// 		prisoner: 465248213,
		// 		candy:    679599042,
		// 		start:    334563499,
		// 	},
		// 	want: 83666114,
		// },
		// "case_0.81": {
		// 	args: args{
		// 		prisoner: 331230504,
		// 		candy:    637771661,
		// 		start:    76778140,
		// 	},
		// 	want: 52088792,
		// },
		"case_0.82": {
			args: args{
				prisoner: 976340152,
				candy:    988657462,
				start:    243958260,
			},
			want: 256275569,
		},
		// "case_0.83": {
		// 	args: args{
		// 		prisoner: 552994811,
		// 		candy:    922743535,
		// 		start:    540135280,
		// 	},
		// 	want: 356889192,
		// },
		"case_0.84": {
			args: args{
				prisoner: 626831986,
				candy:    839281366,
				start:    24222267,
			},
			want: 236671646,
		},
		"case_0.85": {
			args: args{
				prisoner: 157704591,
				candy:    253731033,
				start:    59023773,
			},
			want: 155050214,
		},
		"case_0.86": {
			args: args{
				prisoner: 806377559,
				candy:    859228114,
				start:    238044289,
			},
			want: 290894843,
		},
		"case_0.87": {
			args: args{
				prisoner: 838798445,
				candy:    996851254,
				start:    268459446,
			},
			want: 426512254,
		},
		"case_0.88": {
			args: args{
				prisoner: 847646888,
				candy:    928001503,
				start:    755190846,
			},
			want: 835545460,
		},
		// "case_0.89": {
		// 	args: args{
		// 		prisoner: 877625009,
		// 		candy:    999275937,
		// 		start:    874127074,
		// 	},
		// 	want: 118152992,
		// },
		"case_0.90": {
			args: args{
				prisoner: 847949466,
				candy:    893343194,
				start:    10497830,
			},
			want: 55891557,
		},
		"case_0.91": {
			args: args{
				prisoner: 35029316,
				candy:    784675240,
				start:    8200127,
			},
			want: 22230414,
		},
		"case_0.92": {
			args: args{
				prisoner: 111807455,
				candy:    690309882,
				start:    23190325,
			},
			want: 42655476,
		},
		// "case_0.93": {
		// 	args: args{
		// 		prisoner: 355765714,
		// 		candy:    554560093,
		// 		start:    311565654,
		// 	},
		// 	want: 154594318,
		// },
		"case_0.94": {
			args: args{
				prisoner: 1890615,
				candy:    614626804,
				start:    976253,
			},
			want: 1153181,
		},
		"case_0.95": {
			args: args{
				prisoner: 132293206,
				candy:    165429783,
				start:    65360680,
			},
			want: 98497256,
		},
		// "case_0.96": {
		// 	args: args{
		// 		prisoner: 506726160,
		// 		candy:    934170235,
		// 		start:    455394293,
		// 	},
		// 	want: 376112207,
		// },
		// "case_0.97": {
		// 	args: args{
		// 		prisoner: 210041918,
		// 		candy:    328800789,
		// 		start:    159203369,
		// 	},
		// 	want: 67920321,
		// },
		"case_0.98": {
			args: args{
				prisoner: 499999999,
				candy:    999999997,
				start:    2,
			},
			want: 499999999,
		},
		"case_0.99": {
			args: args{
				prisoner: 499999999,
				candy:    999999998,
				start:    2,
			},
			want: 1,
		},
		// "case_0.100": {
		// 	args: args{
		// 		prisoner: 999999999,
		// 		candy:    999999999,
		// 		start:    1,
		// 	},
		// 	want: 999999999,
		// },
		"case_4.1": {
			args: args{
				prisoner: 5,
				candy:    2,
				start:    1,
			},
			want: 2,
		},
		"case_4.2": {
			args: args{
				prisoner: 5,
				candy:    2,
				start:    2,
			},
			want: 3,
		},
	}
	for name, test := range tests {
		tt := test
		t.Run(name, func(t *testing.T) {
			if got := saveThePrisoner(tt.args.prisoner, tt.args.candy, tt.args.start); got != tt.want {
				t.Errorf("saveThePrisoner() = %v, want %v", got, tt.want)
			}
		})
	}
}
