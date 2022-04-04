package essential

func Replace(s *string, old string, new string, j int) {
	var str []rune
	for i, rn := range *s {
		str[i] = rn
	}
	if j == -1 {
		/*
				s = "test"
		old = "st"
		new = "no"

		split old and new and s

		start test loop
		"t" "e" "s" "t"
		 ↑   ↑  x
		"s" "t"
		repeat move by one
		"t" "e" "s" "t"
		     ↑   ↑  x
		    "s" "t"
		repeat move by one
		"t" "e" "s" "t"
		         ↑   ↑  good
		        "s" "t" -> "n" "o"
		stop test loop

		replace
		"t" "e" "n" "o"
		join s
		s = "teno"
		*/
	} else {

	}
}
