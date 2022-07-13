package main

type Name struct {
	Number    int
	Symbol    string
	Name      string
	Origin    string
	Group     int
	Period    int
	Weight    float64
	Density   float64
	Melt      float64
	Boil      float64
	C         string
	X         string
	Abundance float64
	Property  string
}

var Mendel = []Name{
	{
		Number:    1,
		Symbol:    "H",
		Name:      "Hydrogen",
		Origin:    "composed of the Greek Names <i>hydro-</i> and <i>-gen</i> meaning 'water-forming'",
		Group:     1,
		Period:    1,
		Weight:    1.008,
		Density:   0.00008988,
		Melt:      14.01,
		Boil:      20.28,
		C:         "14.304",
		X:         "2.20",
		Abundance: 1400,
		Property:  "diatomic_nonmetal",
	},
	{
		Number:    2,
		Symbol:    "He",
		Name:      "Helium",
		Origin:    "the Greek <i>helios</i>, 'sun'",
		Group:     18,
		Period:    1,
		Weight:    4.002602,
		Density:   0.0001785,
		Boil:      4.22,
		C:         "5.193",
		X:         "–",
		Abundance: 0.008,
		Property:  "noble_gas",
	}, {
		Number:    3,
		Symbol:    "Li",
		Name:      "Lithium",
		Origin:    "the Greek <i>lithos</i>, 'stone'",
		Group:     1,
		Period:    2,
		Weight:    6.94,
		Density:   0.534,
		Melt:      453.69,
		Boil:      1560,
		C:         "3.582",
		X:         "0.98",
		Abundance: 20,
		Property:  "alkali_metal",
	}, {
		Number:    4,
		Symbol:    "Be",
		Name:      "Beryllium",
		Origin:    "beryl, a mineral",
		Group:     2,
		Period:    2,
		Weight:    9.0121831,
		Density:   1.85,
		Melt:      1560,
		Boil:      2742,
		C:         "1.825",
		X:         "1.57",
		Abundance: 2.8,
		Property:  "alkaline_earth_metal",
	}, {
		Number:    5,
		Symbol:    "B",
		Name:      "Boron",
		Origin:    "borax, a mineral",
		Group:     13,
		Period:    2,
		Weight:    10.81,
		Density:   2.34,
		Melt:      2349,
		Boil:      4200,
		C:         "1.026",
		X:         "2.04",
		Abundance: 10,
		Property:  "metalloid",
	}, {
		Number:    6,
		Symbol:    "C",
		Name:      "Carbon",
		Origin:    "the Latin <i>carbo</i>, 'coal'",
		Group:     14,
		Period:    2,
		Weight:    12.011,
		Density:   2.267,
		Melt:      3800,
		Boil:      4300,
		C:         "0.709",
		X:         "2.55",
		Abundance: 200,
		Property:  "polyatomic_nonmetal",
	},
	/*
	   {
	   		Number:     7,
	     Symbol: "N",
	     Name: "Nitrogen",
	     Origin: "the Greek <i>nitron</i> and '-gen' meaning 'niter-forming'",
	     Group: 15,
	     Period: 2,
	     Weight: 14.007,
	     Density: 0.0012506,
	     Melt: 63.15,
	     Boil: 77.36,
	     C: "1.04",
	     X: "3.04",
	     Abundance: 19,
	     Property: "diatomic_nonmetal",
	   },{
	   		Number:     8,
	     Symbol: "O",
	     Name: "Oxygen",
	     Origin: "from the Greek <i>oxy-</i>, both 'sharp' and 'acid', and <i>-gen</i>, meaning 'acid-forming'",
	     Group: 16,
	     Period: 2,
	     Weight: 15.999,
	     Density: 0.001429,
	     Melt: 54.36,
	     Boil: 90.20,
	     C: "0.918",
	     X: "3.44",
	     Abundance: 461000,
	     Property: "diatomic_nonmetal",
	   },{
	   		Number:     9,
	     Symbol: "F",
	     Name: "Fluorine",
	     Origin: "the Latin <i>fluere</i>, 'to flow'",
	     Group: 17
	     Period: 2
	     Weight: 18.998403163(6)
	     Density: 0.001696
	     Melt: 53.53
	     Boil: 85.03
	     C: "0.824"
	     X: "3.98"
	     Abundance: 585
	     Property: "diatomic_nonmetal"
	   },{
	   		Number:     10
	     Symbol: "Ne"
	     Name: "Neon"
	     Origin: "the Greek <i>neos</i>, meaning 'new'"
	     Group: 18
	     Period: 2
	     Weight: 20.1797(6)<sup id="cite_ref-fn2_9-7" class="reference"><a href="#cite_note-fn2-9">[III]</sup><sup id="cite_ref-fn3_10-3" class="reference"><a href="#cite_note-fn3-10">[IV]</sup>
	     Density: 0.0008999
	     Melt: 24.56
	     Boil: 27.07
	     C: "1.03"
	     X: "–"
	     Abundance: 0.005
	     Property: "noble_gas"
	   },{
	   		Number:     11
	     Symbol: "Na"
	     Name: "Sodium"
	     Origin: "the English word <i>soda</i> (<i>natrium</i> in Latin)"
	     Group: 1
	     Period: 3
	     Weight: 22.98976928(2)
	     Density: 0.971
	     Melt: 370.87
	     Boil: 1156
	     C: "1.228"
	     X: "0.93"
	     Abundance: 23600
	     Property: "alkali_metal"
	   },{
	   		Number:     12
	     Symbol: "Mg"
	     Name: "Magnesium"
	     Origin: "Magnesia, a district of Eastern Thessaly in Greece"
	     Group: 2
	     Period: 3
	     Weight: 24.305
	     Density: 1.738
	     Melt: 923
	     Boil: 1363
	     C: "1.023"
	     X: "1.31"
	     Abundance: 23300
	     Property: alkaline_earth_metal
	   },{
	   		Number:     13
	     Symbol: "Al"
	     Name: "Aluminium"
	     Origin: "from <i>alumina</i>, a compound (Originally <i>aluminum</i>)"
	     Group: 13
	     Period: 3
	     Weight: 26.9815385(7)
	     Density: 2.698
	     Melt: 933.47
	     Boil: 2792
	     C: "0.897"
	     X: "1.61"
	     Abundance: 82300
	     Property: "post_​transition_metal"
	   },{
	   		Number:     14
	     Symbol: "Si"
	     Name: "Silicon"
	     Origin: "from the Latin <i>silex</i>, 'flint' (Originally <i>silicium</i>)"
	     Group: 14
	     Period: 3
	     Weight: 28.085
	     Density: 2.3296
	     Melt: 1687
	     Boil: 3538
	     C: "0.705"
	     X: "1.9"
	     Abundance: 282000
	     Property: metalloid
	   },{
	   		Number:     15
	     Symbol: "P"
	     Name: "Phosphorus"
	     Origin: "the Greek <i>phoosphoros</i>, 'carrying light'"
	     Group: 15
	     Period: 3
	     Weight: 30.973761998(5)
	     Density: 1.82
	     Melt: 317.30
	     Boil: 550
	     C: "0.769"
	     X: "2.19"
	     Abundance: 1050
	     Property: polyatomic_nonmetal
	   },{
	   		Number:     16
	     Symbol: "S"
	     Name: "Sulfur"
	     Origin: "the Latin <i>sulphur</i>, 'fire and brimstone'"
	     Group: 16
	     Period: 3
	     Weight: 32.06
	     Density: 2.067
	     Melt: 388.36
	     Boil: 717.87
	     C: "0.71"
	     X: "2.58"
	     Abundance: 350
	     Property: polyatomic_nonmetal
	   },{
	   		Number:     17
	     Symbol: "Cl"
	     Name: "Chlorine"
	     Origin: "the Greek <i>chloros</i>, 'greenish yellow'"
	     Group: 17
	     Period: 3
	     Weight: 35.45
	     Density: 0.003214
	     Melt: 171.6
	     Boil: 239.11
	     C: "0.479"
	     X: "3.16"
	     Abundance: 145
	     Property: "diatomic_nonmetal"
	   },{
	   		Number:     18
	     Symbol: "Ar"
	     Name: "Argon"
	     Origin: "the Greek <i>argos</i>, 'idle'"
	     Group: 18
	     Period: 3
	     Weight: 39.948(1)
	     Density: 0.0017837
	     Melt: 83.80
	     Boil: 87.30
	     C: "0.52"
	   #  X: "–"
	     Abundance: 3.5
	     Property: "noble_gas"
	   },{
	   		Number:     19
	     Symbol: "K"
	     Name: "Potassium"
	     Origin: "New Latin <i>potassa</i>, 'potash' (<i>kalium</i> in Latin)"
	     Group: 1
	     Period: 4
	     Weight: 39.0983(1)
	     Density: 0.862
	     Melt: 336.53
	     Boil: 1032
	     C: "0.757"
	     X: "0.82"
	     Abundance: 20900
	     Property: "alkali_metal"
	   },{
	   		Number:     20
	     Symbol: "Ca"
	     Name: "Calcium"
	     Origin: "the Latin <i>calx</i>, 'lime'"
	     Group: 2
	     Period: 4
	     Weight: 40.078(4)
	     Density: 1.54
	     Melt: 1115
	     Boil: 1757
	     C: "0.647"
	     X: "1"
	     Abundance: 41500
	     Property: alkaline_earth_metal
	   },{
	   		Number:     21
	     Symbol: "Sc"
	     Name: "Scandium"
	     Origin: "<i>Scandia</i>, the Latin name for Scandinavia"
	     Group: 3
	     Period: 4
	     Weight: 44.955908(5)
	     Density: 2.989
	     Melt: 1814
	     Boil: 3109
	     C: "0.568"
	     X: "1.36"
	     Abundance: 22
	     Property: "transition_metal"
	   },{
	   		Number:     22
	     Symbol: "Ti"
	     Name: "Titanium"
	     Origin: "Titans, the sons of the Earth goddess of Greek mythology"
	     Group: 4
	     Period: 4
	     Weight: 47.867(1)
	     Density: 4.54
	     Melt: 1941
	     Boil: 3560
	     C: "0.523"
	     X: "1.54"
	     Abundance: 5650
	     Property: "transition_metal"
	   },{
	   		Number:     23
	     Symbol: "V"
	     Name: "Vanadium"
	     Origin: "Vanadis, an Old Norse name for the Scandinavian goddess Freyja"
	     Group: 5
	     Period: 4
	     Weight: 50.9415(1)
	     Density: 6.11
	     Melt: 2183
	     Boil: 3680
	     C: "0.489"
	     X: "1.63"
	     Abundance: 120
	     Property: "transition_metal"
	   },{
	   		Number:     24
	     Symbol: "Cr"
	     Name: "Chromium"
	     Origin: "the Greek <i>chroma</i>, 'color'"
	     Group: 6
	     Period: 4
	     Weight: 51.9961(6)
	     Density: 7.15
	     Melt: 2180
	     Boil: 2944
	     C: "0.449"
	     X: "1.66"
	     Abundance: 102
	     Property: "transition_metal"
	   },{
	   		Number:     25
	     Symbol: "Mn"
	     Name: "Manganese"
	     Origin: "corrupted from <i>magnesia negra</i>, see Magnesium"
	     Group: 7
	     Period: 4
	     Weight: 54.938044(3)
	     Density: 7.44
	     Melt: 1519
	     Boil: 2334
	     C: "0.479"
	     X: "1.55"
	     Abundance: 950
	     Property: "transition_metal"
	   },{
	   		Number:     26
	     Symbol: "Fe"
	     Name: "Iron"
	     Origin: "English word (<i>ferrum</i> in Latin)"
	     Group: 8
	     Period: 4
	     Weight: 55.845(2)
	     Density: 7.874
	     Melt: 1811
	     Boil: 3134
	     C: "0.449"
	     X: "1.83"
	     Abundance: 56300
	     Property: "transition_metal"
	   },{
	   		Number:     27
	     Symbol: "Co"
	     Name: "Cobalt"
	     Origin: "the German word <i>Kobold</i>, 'goblin'"
	     Group: 9
	     Period: 4
	     Weight: 58.933194(4)
	     Density: 8.86
	     Melt: 1768
	     Boil: 3200
	     C: "0.421"
	     X: "1.88"
	     Abundance: 25
	     Property: "transition_metal"
	   },{
	   		Number:     28
	     Symbol: "Ni"
	     Name: "Nickel"
	     Origin: "from a mischievous sprite of German miner mythology , Nickel"
	     Group: 10
	     Period: 4
	     Weight: 58.6934(4)
	     Density: 8.912
	     Melt: 1728
	     Boil: 3186
	     C: "0.444"
	     X: "1.91"
	     Abundance: 84
	     Property: "transition_metal"
	   },{
	   		Number:     29
	     Symbol: "Cu"
	     Name: "Copper"
	     Origin: "English word (Latin <i>cuprum</i>)"
	     Group: 11
	     Period: 4
	     Weight: 63.546(3)
	     Density: 8.96
	     Melt: 1357.77
	     Boil: 2835
	     C: "0.385"
	     X: "1.9"
	     Abundance: 60
	     Property: "transition_metal"
	   },{
	   		Number:     30
	     Symbol: "Zn"
	     Name: "Zinc"
	     Origin: "German word <i>Zinke</i> (prong, tooth)"
	     Group: 12
	     Period: 4
	     Weight: 65.38(2)
	     Density: 7.134
	     Melt: 692.88
	     Boil: 1180
	     C: "0.388"
	     X: "1.65"
	     Abundance: 70
	     Property: "post_​transition_metal"
	   },{
	   		Number:     31
	     Symbol: "Ga"
	     Name: "Gallium"
	     Origin: "<i>Gallia</i>, the Latin name for France"
	     Group: 13
	     Period: 4
	     Weight: 69.723(1)
	     Density: 5.907
	     Melt: 302.9146
	     Boil: 2673
	     C: "0.371"
	     X: "1.81"
	     Abundance: 19
	     Property: "post_​transition_metal"
	   },{
	   		Number:     32
	     Symbol: "Ge"
	     Name: "Germanium"
	     Origin: "<i>Germania</i>, the Latin name for Germany"
	     Group: 14
	     Period: 4
	     Weight: 72.630(8)
	     Density: 5.323
	     Melt: 1211.40
	     Boil: 3106
	     C: "0.32"
	     X: "2.01"
	     Abundance: 1.5
	     Property: metalloid
	   },{
	   		Number:     33
	     Symbol: "As"
	     Name: "Arsenic"
	     Origin: "English word (Latin <i>arsenicum</i>)"
	     Group: 15
	     Period: 4
	     Weight: 74.921595(6)
	     Density: 5.776
	     Melt: 1090 <sup id="cite_ref-fn7_15-0" class="reference"><a href="#cite_note-fn7-15">[IX]</sup>
	     Boil: 887
	     C: "0.329"
	     X: "2.18"
	     Abundance: 1.8
	     Property: metalloid
	   },{
	   		Number:     34
	     Symbol: "Se"
	     Name: "Selenium"
	     Origin: "the Greek <i>selene</i>, 'moon'"
	     Group: 16
	     Period: 4
	     Weight: 78.971(8)
	     Density: 4.809
	     Melt: 453
	     Boil: 958
	     C: "0.321"
	     X: "2.55"
	     Abundance: 0.05
	     Property: polyatomic_nonmetal
	   },{
	   		Number:     35
	     Symbol: "Br"
	     Name: "Bromine"
	     Origin: "the Greek <i>bromos</i>, 'stench'"
	     Group: 17
	     Period: 4
	     Weight: 79.904
	     Density: 3.122
	     Melt: 265.8
	     Boil: 332.0
	     C: "0.474"
	     X: "2.96"
	     Abundance: 2.4
	     Property: "diatomic_nonmetal"
	   },{
	   		Number:     36
	     Symbol: "Kr"
	     Name: "Krypton"
	     Origin: "the Greek <i>kryptos</i>, 'hidden'"
	     Group: 18
	     Period: 4
	     Weight: 83.798(2)
	     Density: 0.003733
	     Melt: 115.79
	     Boil: 119.93
	     C: "0.248"
	     X: "3"
	     Abundance: 1×10<sup>−4</sup>
	     Property: "noble_gas"
	   },{
	   		Number:     37
	     Symbol: "Rb"
	     Name: "Rubidium"
	     Origin: "the Latin <i>rubidus</i>, 'deep red'"
	     Group: 1
	     Period: 5
	     Weight: 85.4678(3)
	     Density: 1.532
	     Melt: 312.46
	     Boil: 961
	     C: "0.363"
	     X: "0.82"
	     Abundance: 90
	     Property: "alkali_metal"
	   },{
	   		Number:     38
	     Symbol: "Sr"
	     Name: "Strontium"
	     Origin: "Strontian, a small town in Scotland"
	     Group: 2
	     Period: 5
	     Weight: 87.62(1)
	     Density: 2.64
	     Melt: 1050
	     Boil: 1655
	     C: "0.301"
	     X: "0.95"
	     Abundance: 370
	     Property: alkaline_earth_metal
	   },{
	   		Number:     39
	     Symbol: "Y"
	     Name: "Yttrium"
	     Origin: "Ytterby, Sweden"
	     Group: 3
	     Period: 5
	     Weight: 88.90584(2)
	     Density: 4.469
	     Melt: 1799
	     Boil: 3609
	     C: "0.298"
	     X: "1.22"
	     Abundance: 33


	   },{
	   		Number:     40
	     Symbol: "Zr"
	     Name: "Zirconium"
	     Origin: "Persian <i>Zargun</i>, 'gold-colored'; German <i>Zirkoon</i>, 'jargoon'"
	     Group: 4
	     Period: 5
	     Weight: 91.224(2)
	     Density: 6.506
	     Melt: 2128
	     Boil: 4682
	     C: "0.278"
	     X: "1.33"
	     Abundance: 165
	     Property: "transition_metal"
	   },{
	   		Number:     41
	     Symbol: "Nb"
	     Name: "Niobium"
	     Origin: "Niobe, daughter of king Tantalus from Greek mythology"
	     Group: 5
	     Period: 5
	     Weight: 92.90637(2)
	     Density: 8.57
	     Melt: 2750
	     Boil: 5017
	     C: "0.265"
	     X: "1.6"
	     Abundance: 20
	     Property: "transition_metal"
	   },{
	   		Number:     42
	     Symbol: "Mo"
	     Name: "Molybdenum"
	     Origin: "the Greek <i>molybdos</i> meaning 'lead'"
	     Group: 6
	     Period: 5
	     Weight: 95.95(1)
	     Density: 10.22
	     Melt: 2896
	     Boil: 4912
	     C: "0.251"
	     X: "2.16"
	     Abundance: 1.2
	     Property: "transition_metal"
	   },{
	   		Number:     43
	     Symbol: "Tc"
	     Name: "Technetium"
	     Origin: "the Greek <i>tekhnètos</i> meaning 'artificial'"
	     Group: 7
	     Period: 5
	     Weight: "[98]"
	     Density: 11.5
	     Melt: 2430
	     Boil: 4538
	     C: "–"
	     X: "1.9"
	     Abundance: ~&#160;3×10<sup>−9</sup>
	     Property: "transition_metal"
	   },{
	   		Number:     44
	     Symbol: "Ru"
	     Name: "Ruthenium"
	     Origin: "<i>Ruthenia</i>, the New Latin name for Russia"
	     Group: 8
	     Period: 5
	     Weight: 101.07(2)<sup id="cite_ref-fn2_9-17" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 12.37
	     Melt: 2607
	     Boil: 4423
	     C: "0.238"
	     X: "2.2"
	     Abundance: 0.001
	     Property: "transition_metal"
	   },{
	   		Number:     45
	     Symbol: "Rh"
	     Name: "Rhodium"
	     Origin: "the Greek <i>rhodos</i>, meaning 'rose coloured'"
	     Group: 9
	     Period: 5
	     Weight: 102.90550(2)
	     Density: 12.41
	     Melt: 2237
	     Boil: 3968
	     C: "0.243"
	     X: "2.28"
	     Abundance: 0.001
	     Property: "transition_metal"
	   },{
	   		Number:     46
	     Symbol: "Pd"
	     Name: "Palladium"
	     Origin: "the then recently discovered asteroid Pallas, considered a planet at the time"
	     Group: 10
	     Period: 5
	     Weight: 106.42(1)<sup id="cite_ref-fn2_9-18" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 12.02
	     Melt: 1828.05
	     Boil: 3236
	     C: "0.244"
	     X: "2.2"
	     Abundance: 0.015
	     Property: "transition_metal"
	   },{
	   		Number:     47
	     Symbol: "Ag"
	     Name: "Silver"
	     Origin: "English word (<i>argentum</i> in Latin)"
	     Group: 11
	     Period: 5
	     Weight: 107.8682(2)<sup id="cite_ref-fn2_9-19" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 10.501
	     Melt: 1234.93
	     Boil: 2435
	     C: "0.235"
	     X: "1.93"
	     Abundance: 0.075
	     Property: "transition_metal"
	   },{
	   		Number:     48
	     Symbol: "Cd"
	     Name: "Cadmium"
	     Origin: "the New Latin <i>cadmia</i>, from King Kadmos"
	     Group: 12
	     Period: 5
	     Weight: 112.414(4)<sup id="cite_ref-fn2_9-20" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 8.69
	     Melt: 594.22
	     Boil: 1040
	     C: "0.232"
	     X: "1.69"
	     Abundance: 0.159
	     Property: "post_​transition_metal"
	   },{
	   		Number:     49
	     Symbol: "In"
	     Name: "Indium"
	     Origin: "<i>indigo</i>"
	     Group: 13
	     Period: 5
	     Weight: 114.818(1)
	     Density: 7.31
	     Melt: 429.75
	     Boil: 2345
	     C: "0.233"
	     X: "1.78"
	     Abundance: 0.25
	     Property: "post_​transition_metal"
	   },{
	   		Number:     50
	     Symbol: "Sn"
	     Name: "Tin"
	     Origin: "English word (<i>stannum</i> in Latin)"
	     Group: 14
	     Period: 5
	     Weight: 118.710(7)<sup id="cite_ref-fn2_9-21" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 7.287
	     Melt: 505.08
	     Boil: 2875
	     C: "0.228"
	     X: "1.96"
	     Abundance: 2.3
	     Property: "post_​transition_metal"
	   },{
	   		Number:     51
	     Symbol: "Sb"
	     Name: "Antimony"
	     Origin: "uncertain: perhaps from the Greek <i>anti</i>, 'against', and <i>monos</i>, 'alone', or the Old French <i>antimoine</i>, 'Monk's bane' (<i>stibium</i> in Latin)"
	     Group: 15
	     Period: 5
	     Weight: 121.760(1)
	     Density: 6.685
	     Melt: 903.78
	     Boil: 1860
	     C: "0.207"
	     X: "2.05"
	     Abundance: 0.2
	     Property: metalloid
	   },{
	   		Number:     52
	     Symbol: "Te"
	     Name: "Tellurium"
	     Origin: "Latin <i>tellus</i>, 'earth'"
	     Group: 16
	     Period: 5
	     Weight: 127.60(3)<sup id="cite_ref-fn2_9-23" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 6.232
	     Melt: 722.66
	     Boil: 1261
	     C: "0.202"
	     X: "2.1"
	     Abundance: 0.001
	     Property: metalloid
	   },{
	   		Number:     53
	     Symbol: "I"
	     Name: "Iodine"
	     Origin: "French <i>iode</i> (after the Greek <i>ioeides</i>, 'violet')"
	     Group: 17
	     Period: 5
	     Weight: 126.90447(3)
	     Density: 4.93
	     Melt: 386.85
	     Boil: 457.4
	     C: "0.214"
	     X: "2.66"
	     Abundance: 0.45
	     Property: "diatomic_nonmetal"
	   },{
	   		Number:     54
	     Symbol: "Xe"
	     Name: "Xenon"
	     Origin: "the Greek <i>xenos</i>, 'strange'"
	     Group: 18
	     Period: 5
	     Weight: 131.293(6)<sup id="cite_ref-fn2_9-24" class="reference"><a href="#cite_note-fn2-9">[III]</sup><sup id="cite_ref-fn3_10-6" class="reference"><a href="#cite_note-fn3-10">[IV]</sup>
	     Density: 0.005887
	     Melt: 161.4
	     Boil: 165.03
	     C: "0.158"
	     X: "2.6"
	     Abundance: 3×10<sup>−5</sup>
	     Property: "noble_gas"
	   },{
	   		Number:     55
	     Symbol: "Cs"
	     Name: "Caesium"
	     Origin: "the Latin <i>caesius</i>, 'sky blue'"
	     Group: 1
	     Period: 6
	     Weight: 132.90545196(6)
	     Density: 1.873
	     Melt: 301.59
	     Boil: 944
	     C: "0.242"
	     X: "0.79"
	     Abundance: 3
	     Property: "alkali_metal"
	   },{
	   		Number:     56
	     Symbol: "Ba"
	     Name: "Barium"
	     Origin: "the Greek <i>barys</i>, 'heavy'"
	     Group: 2
	     Period: 6
	     Weight: 137.327(7)
	     Density: 3.594
	     Melt: 1000
	     Boil: 2170
	     C: "0.204"
	     X: "0.89"
	     Abundance: 425
	     Property: alkaline_earth_metal
	   },{
	   		Number:     57
	     Symbol: "La"
	     Name: "Lanthanum"
	     Origin: "the Greek <i>lanthanein</i>, 'to lie hidden'"
	     Group: 3
	     Period: 6
	     Weight: 138.90547(7)<sup id="cite_ref-fn2_9-25" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 6.145
	     Melt: 1193
	     Boil: 3737
	     C: "0.195"
	     X: "1.1"
	     Abundance: 39
	     Property: "lan­thanide"
	   },{
	   		Number:     58
	     Symbol: "Ce"
	     Name: "Cerium"
	     Origin: "the then recently discovered asteroid Ceres, considered a planet at the time"
	     Group: -4 # not a Group, but used as coordinates in the UI grid
	     Period: 6
	     Weight: 140.116(1)<sup id="cite_ref-fn2_9-26" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 6.77
	     Melt: 1068
	     Boil: 3716
	     C: "0.192"
	     X: "1.12"
	     Abundance: 66.5
	     Property: "lan­thanide"
	   },{
	   		Number:     59
	     Symbol: "Pr"
	     Name: "Praseodymium"
	     Origin: "the Greek <i>praseios didymos</i> meaning 'green twin'"
	     Group: -5 # not a Group, but used as coordinates in the UI grid
	     Period: 6
	     Weight: 140.90766(2)
	     Density: 6.773
	     Melt: 1208
	     Boil: 3793
	     C: "0.193"
	     X: "1.13"
	     Abundance: 9.2
	     Property: "lan­thanide"
	   },{
	   		Number:     60
	     Symbol: "Nd"
	     Name: "Neodymium"
	     Origin: "the Greek <i>neos didymos</i> meaning 'new twin'"
	     Group: -6 # not a Group, but used as coordinates in the UI grid
	     Period: 6
	     Weight: 144.242(3)<sup id="cite_ref-fn2_9-27" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 7.007
	     Melt: 1297
	     Boil: 3347
	     C: "0.19"
	     X: "1.14"
	     Abundance: 41.5
	     Property: "lan­thanide"
	   },{
	   		Number:     61
	     Symbol: "Pm"
	     Name: "Promethium"
	     Origin: "Prometheus of Greek mythology who stole fire from the Gods and gave it to humans"
	   #  Group: "-"
	     Period: 6
	     Weight: "[145]"
	     Density: 7.26
	     Melt: 1315
	     Boil: 3273
	     C: "–"
	     X: "1.13"
	     Abundance: 2×10<sup>−19</sup>
	     Property: "lan­thanide"
	   },{
	   		Number:     62
	     Symbol: "Sm"
	     Name: "Samarium"
	     Origin: "Samarskite, the name of the mineral from which it was first isolated"
	   #  Group: "-"
	     Period: 6
	     Weight: 150.36(2)<sup id="cite_ref-fn2_9-28" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 7.52
	     Melt: 1345
	     Boil: 2067
	     C: "0.197"
	     X: "1.17"
	     Abundance: 7.05
	     Property: "lan­thanide"
	   },{
	   		Number:     63
	     Symbol: "Eu"
	     Name: "Europium"
	     Origin: "Europe"
	   #  Group: "-"
	     Period: 6
	     Weight: 151.964(1)<sup id="cite_ref-fn2_9-29" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 5.243
	     Melt: 1099
	     Boil: 1802
	     C: "0.182"
	     X: "1.2"
	     Abundance: 2
	     Property: "lan­thanide"
	   },{
	   		Number:     64
	     Symbol: "Gd"
	     Name: "Gadolinium"
	     Origin: "Johan Gadolin, chemist, physicist and mineralogist"
	   #  Group: "-"
	     Period: 6
	     Weight: 157.25(3)<sup id="cite_ref-fn2_9-30" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 7.895
	     Melt: 1585
	     Boil: 3546
	     C: "0.236"
	     X: "1.2"
	     Abundance: 6.2
	     Property: "lan­thanide"
	   },{
	   		Number:     65
	     Symbol: "Tb"
	     Name: "Terbium"
	     Origin: "Ytterby, Sweden"
	   #  Group: "-"
	     Period: 6
	     Weight: 158.92535(2)
	     Density: 8.229
	     Melt: 1629
	     Boil: 3503
	     C: "0.182"
	     X: "1.2"
	     Abundance: 1.2
	     Property: "lan­thanide"
	   },{
	   		Number:     66
	     Symbol: "Dy"
	     Name: "Dysprosium"
	     Origin: "the Greek <i>dysprositos</i>, 'hard to get'"
	   #  Group: "-"
	     Period: 6
	     Weight: 162.500(1)<sup id="cite_ref-fn2_9-31" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 8.55
	     Melt: 1680
	     Boil: 2840
	     C: "0.17"
	     X: "1.22"
	     Abundance: 5.2
	     Property: "lan­thanide"
	   },{
	   		Number:     67
	     Symbol: "Ho"
	     Name: "Holmium"
	     Origin: "<i>Holmia</i>, the New Latin name for Stockholm"
	   #  Group: "-"
	     Period: 6
	     Weight: 164.93033(2)
	     Density: 8.795
	     Melt: 1734
	     Boil: 2993
	     C: "0.165"
	     X: "1.23"
	     Abundance: 1.3
	     Property: "lan­thanide"
	   },{
	   		Number:     68
	     Symbol: "Er"
	     Name: "Erbium"
	     Origin: "Ytterby, Sweden"
	   #  Group: "-"
	     Period: 6
	     Weight: 167.259(3)<sup id="cite_ref-fn2_9-32" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 9.066
	     Melt: 1802
	     Boil: 3141
	     C: "0.168"
	     X: "1.24"
	     Abundance: 3.5
	     Property: "lan­thanide"
	   },{
	   		Number:     69
	     Symbol: "Tm"
	     Name: "Thulium"
	     Origin: "Thule, the ancient name for Scandinavia"
	   #  Group: "-"
	     Period: 6
	     Weight: 168.93422(2)
	     Density: 9.321
	     Melt: 1818
	     Boil: 2223
	     C: "0.16"
	     X: "1.25"
	     Abundance: 0.52
	     Property: "lan­thanide"
	   },{
	   		Number:     70
	     Symbol: "Yb"
	     Name: "Ytterbium"
	     Origin: "Ytterby, Sweden"
	   #  Group: "-"
	     Period: 6
	     Weight: 173.045(10)<sup id="cite_ref-fn2_9-33" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 6.965
	     Melt: 1097
	     Boil: 1469
	     C: "0.155"
	     X: "1.1"
	     Abundance: 3.2
	     Property: "lan­thanide"
	   },{
	   		Number:     71
	     Symbol: "Lu"
	     Name: "Lutetium"
	     Origin: "<i>Lutetia</i>, the Latin name for Paris"
	   #  Group: "-"
	     Period: 6
	     Weight: 174.9668(1)<sup id="cite_ref-fn2_9-34" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 9.84
	     Melt: 1925
	     Boil: 3675
	     C: "0.154"
	     X: "1.27"
	     Abundance: 0.8
	     Property: "lan­thanide"
	   },{
	   		Number:     72
	     Symbol: "Hf"
	     Name: "Hafnium"
	     Origin: "<i>Hafnia</i>, the New Latin name for Copenhagen"
	     Group: 4
	     Period: 6
	     Weight: 178.49(2)
	     Density: 13.31
	     Melt: 2506
	     Boil: 4876
	     C: "0.144"
	     X: "1.3"
	     Abundance: 3
	     Property: "transition_metal"
	   },{
	   		Number:     73
	     Symbol: "Ta"
	     Name: "Tantalum"
	     Origin: "King Tantalus, father of Niobe from Greek mythology"
	     Group: 5
	     Period: 6
	     Weight: 180.94788(2)
	     Density: 16.654
	     Melt: 3290
	     Boil: 5731
	     C: "0.14"
	     X: "1.5"
	     Abundance: 2
	     Property: "transition_metal"
	   },{
	   		Number:     74
	     Symbol: "W"
	     Name: "Tungsten"
	     Origin: "the Swedish <i>tung sten</i>, 'heavy stone' (W is <i>wolfram</i>, the old name of the tungsten mineral wolframite)"
	     Group: 6
	     Period: 6
	     Weight: 183.84(1)
	     Density: 19.25
	     Melt: 3695
	     Boil: 5828
	     C: "0.132"
	     X: "2.36"
	     Abundance: 1.3
	     Property: "transition_metal"
	   },{
	   		Number:     75
	     Symbol: "Re"
	     Name: "Rhenium"
	     Origin: "<i>Rhenus</i>, the Latin name for the river Rhine"
	     Group: 7
	     Period: 6
	     Weight: 186.207(1)
	     Density: 21.02
	     Melt: 3459
	     Boil: 5869
	     C: "0.137"
	     X: "1.9"
	     Abundance: 7×10<sup>−4</sup>
	     Property: "transition_metal"
	   },{
	   		Number:     76
	     Symbol: "Os"
	     Name: "Osmium"
	     Origin: "the Greek <i>osmè</i>, meaning 'smell'"
	     Group: 8
	     Period: 6
	     Weight: 190.23(3)<sup id="cite_ref-fn2_9-35" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 22.61
	     Melt: 3306
	     Boil: 5285
	     C: "0.13"
	     X: "2.2"
	     Abundance: 0.002
	     Property: "transition_metal"
	   },{
	   		Number:     77
	     Symbol: "Ir"
	     Name: "Iridium"
	     Origin: "Iris, the Greek goddess of the rainbow"
	     Group: 9
	     Period: 6
	     Weight: 192.217(3)
	     Density: 22.56
	     Melt: 2719
	     Boil: 4701
	     C: "0.131"
	     X: "2.2"
	     Abundance: 0.001
	     Property: "transition_metal"
	   },{
	   		Number:     78
	     Symbol: "Pt"
	     Name: "Platinum"
	     Origin: "the Spanish <i>platina</i>, meaning 'little silver'"
	     Group: 10
	     Period: 6
	     Weight: 195.084(9)
	     Density: 21.46
	     Melt: 2041.4
	     Boil: 4098
	     C: "0.133"
	     X: "2.28"
	     Abundance: 0.005
	     Property: "transition_metal"
	   },{
	   		Number:     79
	     Symbol: "Au"
	     Name: "Gold"
	     Origin: "English word (<i>aurum</i> in Latin)"
	     Group: 11
	     Period: 6
	     Weight: 196.966569(5)
	     Density: 19.282
	     Melt: 1337.33
	     Boil: 3129
	     C: "0.129"
	     X: "2.54"
	     Abundance: 0.004
	     Property: "transition_metal"
	   },{
	   		Number:     80
	     Symbol: "Hg"
	     Name: "Mercury"
	     Origin: "the New Latin name <i>mercurius</i>, named after the Roman god (Hg from former name <i>hydrargyrum</i>, from Greek <i>hydr-</i>, 'water', and <i>argyros</i>, 'silver')"
	     Group: 12
	     Period: 6
	     Weight: 200.592(3)
	     Density: 13.5336
	     Melt: 234.43
	     Boil: 629.88
	     C: "0.14"
	     X: "2"
	     Abundance: 0.085


	   },{
	   		Number:     81
	     Symbol: "Tl"
	     Name: "Thallium"
	     Origin: "the Greek <i>thallos</i>, 'green twig'"
	     Group: 13
	     Period: 6
	     Weight: 204.38<sup id="cite_ref-fn9_12-11" class="reference"><a href="#cite_note-fn9-12">[VI]</sup>
	     Density: 11.85
	     Melt: 577
	     Boil: 1746
	     C: "0.129"
	     X: "1.62"
	     Abundance: 0.85


	   },{
	   		Number:     82
	     Symbol: "Pb"
	     Name: "Lead"
	     Origin: "English word (<i>plumbum</i> in Latin)"
	     Group: 14
	     Period: 6
	     Weight: 207.2(1)<sup id="cite_ref-fn2_9-36" class="reference"><a href="#cite_note-fn2-9">[III]</sup><sup id="cite_ref-fn4_11-14" class="reference"><a href="#cite_note-fn4-11">[V]</sup>
	     Density: 11.342
	     Melt: 600.61
	     Boil: 2022
	     C: "0.129"
	     X: "1.87"
	     Abundance: 14
	     Property: "post_​transition_metal"
	   },{
	   		Number:     83
	     Symbol: "Bi"
	     Name: "Bismuth"
	     Origin: "Uncertain, possibly Arabic or German"
	     Group: 15
	     Period: 6
	     Weight: 208.98040(1)<sup id="cite_ref-fn1_16-2" class="reference"><a href="#cite_note-fn1-16">[X]</sup>
	     Density: 9.807
	     Melt: 544.7
	     Boil: 1837
	     C: "0.122"
	     X: "2.02"
	     Abundance: 0.009
	     Property: "post_​transition_metal"
	   },{
	   		Number:     84
	     Symbol: "Po"
	     Name: "Polonium"
	     Origin: "Named after the home country of Marie Curie (<i>Polonia</i>, Latin for Poland), who is also the discoverer of Radium"
	     Group: 16
	     Period: 6
	     Weight: "[209]"
	     Density: 9.32
	     Melt: 527
	     Boil: 1235
	     C: "–"
	     X: "2.0"
	     Abundance: 2×10<sup>−10</sup>
	     Property: "post_​transition_metal"
	   },{
	   		Number:     85
	     Symbol: "At"
	     Name: "Astatine"
	     Origin: "the Greek <i>astatos</i>, 'unstable'"
	     Group: 17
	     Period: 6
	     Weight: "[210]"
	     Density: 7
	     Melt: 575
	     Boil: 610
	     C: "–"
	     X: "2.2"
	     Abundance: 3×10<sup>−20</sup>
	     Property: metalloid
	   },{
	   		Number:     86
	     Symbol: "Rn"
	     Name: "Radon"
	     Origin: "From <i>radium</i>, as it was first detected as an emission from radium during radioactive decay"
	     Group: 18
	     Period: 6
	     Weight: "[222]"
	     Density: 0.00973
	     Melt: 202
	     Boil: 211.3
	     C: "0.094"
	     X: "2.2"
	     Abundance: 4×10<sup>−13</sup>
	     Property: "noble_gas"
	   },{
	   		Number:     87
	     Symbol: "Fr"
	     Name: "Francium"
	     Origin: "<i>Francia</i>, the New Latin name for France"
	     Group: 1
	     Period: 7
	     Weight: "[223]"
	     Density: 1.87
	     Melt: 300
	     Boil: 950
	     C: "–"
	     X: "0.7"
	     Abundance: ~&#160;1×10<sup>−18</sup>
	     Property: "alkali_metal"
	   },{
	   		Number:     88
	     Symbol: "Ra"
	     Name: "Radium"
	     Origin: "the Latin <i>radius</i>, 'ray'"
	     Group: 2
	     Period: 7
	     Weight: "[226]"
	     Density: 5.5
	     Melt: 973
	     Boil: 2010
	     C: "0.094"
	     X: "0.9"
	     Abundance: 9×10<sup>−7</sup>
	     Property: alkaline_earth_metal
	   },{
	   		Number:     89
	     Symbol: "Ac"
	     Name: "Actinium"
	     Origin: "the Greek <i>aktis</i>, 'ray'"
	     Group: 3
	     Period: 7
	     Weight: "[227]"
	     Density: 10.07
	     Melt: 1323
	     Boil: 3471
	     C: "0.12"
	     X: "1.1"
	     Abundance: 5.5×10<sup>−10</sup>
	     Property: "actinide"
	   },{
	   		Number:     90
	     Symbol: "Th"
	     Name: "Thorium"
	     Origin: "Thor, the Scandinavian god of thunder"
	   #  Group: "-"
	     Period: 7
	     Weight: 232.0377(4)<sup id="cite_ref-fn1_16-9" class="reference"><a href="#cite_note-fn1-16">[X]</sup><sup id="cite_ref-fn2_9-37" class="reference"><a href="#cite_note-fn2-9">[III]</sup>
	     Density: 11.72
	     Melt: 2115
	     Boil: 5061
	     C: "0.113"
	     X: "1.3"
	     Abundance: 9.6
	     Property: "actinide"
	   },{
	   		Number:     91
	     Symbol: "Pa"
	     Name: "Protactinium"
	     Origin: "the Greek <i>protos</i>, 'first', and actinium, which is produced through the radioactive decay of protactinium"
	   #  Group: "-"
	     Period: 7
	     Weight: 231.03588(2)<sup id="cite_ref-fn1_16-10" class="reference"><a href="#cite_note-fn1-16">[X]</sup>
	     Density: 15.37
	     Melt: 1841
	     Boil: 4300
	     C: "–"
	     X: "1.5"
	     Abundance: 1.4×10<sup>−6</sup>
	     Property: "actinide"
	   },{
	   		Number:     92
	     Symbol: "U"
	     Name: "Uranium"
	     Origin: "Uranus, the seventh planet in the Solar System"
	   #  Group: "-"
	     Period: 7
	     Weight: 238.02891(3)<sup id="cite_ref-fn1_16-11" class="reference"><a href="#cite_note-fn1-16">[X]</sup>
	     Density: 18.95
	     Melt: 1405.3
	     Boil: 4404
	     C: "0.116"
	     X: "1.38"
	     Abundance: 2.7
	     Property: "actinide"
	   },{
	   		Number:     93
	     Symbol: "Np"
	     Name: "Neptunium"
	     Origin: "Neptune, the eighth planet in the Solar System"
	   #  Group: "-"
	     Period: 7
	     Weight: "[237]"
	     Density: 20.45
	     Melt: 917
	     Boil: 4273
	     C: "–"
	     X: "1.36"
	     Abundance: ≤&#160;3×10<sup>−12</sup>
	     Property: "actinide"
	   },{
	   		Number:     94
	     Symbol: "Pu"
	     Name: "Plutonium"
	     Origin: "Pluto, a dwarf planet in the Solar System (then considered the ninth planet)"
	   #  Group: "-"
	     Period: 7
	     Weight: "[244]"
	     Density: 19.84
	     Melt: 912.5
	     Boil: 3501
	     C: "–"
	     X: "1.28"
	     Abundance: ≤&#160;3×10<sup>−11</sup>
	     Property: "actinide"
	   },{
	   		Number:     95
	     Symbol: "Am"
	     Name: "Americium"
	     Origin: "The Americas, as the Name was first synthesized on the continent, by analogy with europium"
	   #  Group: "-"
	     Period: 7
	     Weight: "[243]"
	     Density: 13.69
	     Melt: 1449
	     Boil: 2880
	     C: "–"
	     X: "1.13"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     96
	     Symbol: "Cm"
	     Name: "Curium"
	     Origin: "Pierre Curie, a physicist, and Marie Curie, a physicist and chemist, named after great scientists by analogy with gadolinium"
	   #  Group: "-"
	     Period: 7
	     Weight: "[247]"
	     Density: 13.51
	     Melt: 1613
	     Boil: 3383
	     C: "–"
	     X: "1.28"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     97
	     Symbol: "Bk"
	     Name: "Berkelium"
	     Origin: "Berkeley, California, where the Name was first synthesized, by analogy with terbium"
	   #  Group: "-"
	     Period: 7
	     Weight: "[247]"
	     Density: 14.79
	     Melt: 1259
	     Boil: 2900
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     98
	     Symbol: "Cf"
	     Name: "Californium"
	     Origin: "California, where the Name was first synthesized"
	   #  Group: "-"
	     Period: 7
	     Weight: "[251]"
	     Density: 15.1
	     Melt: 1173
	     Boil: (1743)
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     99
	     Symbol: "Es"
	     Name: "Einsteinium"
	     Origin: "Albert Einstein, physicist"
	   #  Group: "-"
	     Period: 7
	     Weight: "[252]"
	     Density: 8.84
	     Melt: 1133
	     Boil: (1269)
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     100
	     Symbol: "Fm"
	     Name: "Fermium"
	     Origin: "Enrico Fermi, physicist"
	   #  Group: "-"
	     Period: 7
	     Weight: "[257]"
	     Density: (9.7)
	     Melt: (1125)
	     Boil: –
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     101
	     Symbol: "Md"
	     Name: "Mendelevium"
	     Origin: "Dmitri Mendeleev, chemist and inventor"
	   #  Group: "-"
	     Period: 7
	     Weight: "[258]"
	     Density: (10.3)
	     Melt: (1100)
	     Boil: –
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     102
	     Symbol: "No"
	     Name: "Nobelium"
	     Origin: "Alfred Nobel, chemist, engineer, innovator, and armaments manufacturer"
	   #  Group: "-"
	     Period: 7
	     Weight: "[259]"
	     Density: (9.9)
	     Melt: (1100)
	     Boil: –
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     103
	     Symbol: "Lr"
	     Name: "Lawrencium"
	     Origin: "Ernest O. Lawrence, physicist"
	   #  Group: "-"
	     Period: 7
	     Weight: "[266]"
	     Density: (15.6)
	     Melt: (1900)
	     Boil: –
	     C: "–"
	     X: "1.3"
	     Abundance: 0
	     Property: "actinide"
	   },{
	   		Number:     104
	     Symbol: "Rf"
	     Name: "Rutherfordium"
	     Origin: "Ernest Rutherford, chemist and physicist"
	     Group: 4
	     Period: 7
	     Weight: "[267]"
	     Density: (23.2)
	     Melt: (2400)
	     Boil: (5800)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "transition_metal"
	   },{
	   		Number:     105
	     Symbol: "Db"
	     Name: "Dubnium"
	     Origin: "Dubna, Russia"
	     Group: 5
	     Period: 7
	     Weight: "[268]"
	     Density: (29.3)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "transition_metal"
	   },{
	   		Number:     106
	     Symbol: "Sg"
	     Name: "Seaborgium"
	     Origin: "Glenn T. Seaborg, scientist"
	     Group: 6
	     Period: 7
	     Weight: "[269]"
	     Density: (35.0)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "transition_metal"
	   },{
	   		Number:     107
	     Symbol: "Bh"
	     Name: "Bohrium"
	     Origin: "Niels Bohr, physicist"
	     Group: 7
	     Period: 7
	     Weight: "[270]"
	     Density: (37.1)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "transition_metal"
	   },{
	   		Number:     108
	     Symbol: "Hs"
	     Name: "Hassium"
	     Origin: "Hesse, Germany, where the Name was first synthesized"
	     Group: 8
	     Period: 7
	     Weight: "[277]"
	     Density: (40.7)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "transition_metal"
	   },{
	   		Number:     109
	     Symbol: "Mt"
	     Name: "Meitnerium"
	     Origin: "Lise Meitner, physicist"
	     Group: 9
	     Period: 7
	     Weight: "[278]"
	     Density: (37.4)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     110
	     Symbol: "Ds"
	     Name: "Darmstadtium"
	     Origin: "Darmstadt, Germany, where the Name was first synthesized"
	     Group: 10
	     Period: 7
	     Weight: "[281]"
	     Density: (34.8)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     111
	     Symbol: "Rg"
	     Name: "Roentgenium"
	     Origin: "Wilhelm Conrad Röntgen, physicist"
	     Group: 11
	     Period: 7
	     Weight: "[282]"
	     Density: (28.7)
	     Melt: –
	     Boil: –
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     112
	     Symbol: "Cn"
	     Name: "Copernicium"
	     Origin: "Nicolaus Copernicus, astronomer"
	     Group: 12
	     Period: 7
	     Weight: "[285]"
	     Density: (23.7)
	     Melt: –
	     Boil: ~357
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "post_​transition_metal"
	   },{
	   		Number:     113
	     Symbol: "Nh"
	     Name: "Nihonium"
	     Origin: "the Japanese name for Japan, <i>Nihon</i>, where the Name was first synthesized"
	     Group: 13
	     Period: 7
	     Weight: "[286]"
	     Density: (16)
	     Melt: (700)
	     Boil: (1400)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     114
	     Symbol: "Fl"
	     Name: "Flerovium"
	     Origin: "Flerov Laboratory of Nuclear Reactions, part of JINR where the Name was synthesized; itself named for Georgy Flyorov, physicist"
	     Group: 14
	     Period: 7
	     Weight: "[289]"
	     Density: (14)
	     Melt: –
	     Boil: ~210
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     115
	     Symbol: "Mc"
	     Name: "Moscovium"
	     Origin: "Moscow Oblast, Russia, where the Name was first synthesized"
	     Group: 15
	     Period: 7
	     Weight: "[290]"
	     Density: (13.5)
	     Melt: (700)
	     Boil: (1400)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     116
	     Symbol: "Lv"
	     Name: "Livermorium"
	     Origin: "Lawrence Livermore National Laboratory (in Livermore, California) which collaborated with JINR on its synthesis"
	     Group: 16
	     Period: 7
	     Weight: "[293]"
	     Density: (12.9)
	     Melt: (709)
	     Boil: (1085)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     117
	     Symbol: "Ts"
	     Name: "Tennessine"
	     Origin: "Tennessee, United States"
	     Group: 17
	     Period: 7
	     Weight: "[294]"
	     Density: (7.2)
	     Melt: (723)
	     Boil: (883)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     118
	     Symbol: "Og"
	     Name: "Oganesson"
	     Origin: "Yuri Oganessian, physicist"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   #properties:
	   #  alkali_metal: Alkali metal
	   #  alkaline_earth_metal: Alkaline earth metal
	   #  lan­thanide: Lan­thanide
	   #  actinide: Actinide
	   #  transition_metal: Transition metal
	   #  post_​transition_metal: Post-​transition metal
	   #  metalloid: Metalloid
	   #  polyatomic_nonmetal: Polyatomic nonmetal
	   #  diatomic_nonmetal: Diatomic nonmetal
	   #  noble_gas: Noble gas
	   #  unknown: Unknown
	   },{
	   		Number:     500
	     Symbol: "Ui"
	     Name: "Userium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     501
	     Symbol: "Ai"
	     Name: "Apium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     201
	     Symbol: "Bo"
	     Name: "Boneium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     202
	     Symbol: "Bw"
	     Name: "Bontwoium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     203
	     Symbol: "Bt"
	     Name: "Bontreium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     204
	     Symbol: "Bu"
	     Name: "Bonfourium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     205
	     Symbol: "Bf"
	     Name: "Bonfiveium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"
	   },{
	   		Number:     299
	     Symbol: "Bn"
	     Name: "Bonium"
	     Origin: "Telemetry Reactor"
	     Group: 18
	     Period: 7
	     Weight: "[294]"
	     Density: (5.0)
	     Melt: –
	     Boil: (350)
	     C: "–"
	     X: "–"
	     Abundance: 0
	     Property: "unknown"

	*/
}
