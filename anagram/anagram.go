package anagram

import (
	"fmt"
	"strings"
)

/*
# Problems
wo words are anagrams of one another if their letters can be rearranged to form the other word.
In this challenge, you will be given a string. You must split it into two contiguous substrings, then determine the minimum number of characters to change to make the two substrings into anagrams of one another.
For example, given the string 'abccde', you would break it into two parts: 'abc' and 'cde'. Note that all letters have been used, the substrings are contiguous and their lengths are equal. Now you can change 'a' and 'b' in the first substring to 'd' and 'e' to have 'dec' and 'cde' which are anagrams. Two changes were necessary.

#sample input
aaabbb
ab
abc
mnop
xyyx
xaxbbbxx

#sample output
3
1
-1
2
0
1

#Explanation
Test Case #01: We split  into two strings ='aaa' and ='bbb'. We have to replace all three characters from the first string with 'b' to make the strings anagrams.

Test Case #02: You have to replace 'a' with 'b', which will generate "bb".

Test Case #03: It is not possible for two strings of unequal length to be anagrams of one another.

Test Case #04: We have to replace both the characters of first string ("mn") to make it an anagram of the other one.

Test Case #05:  and  are already anagrams of one another.

Test Case #06: Here S1 = "xaxb" and S2 = "bbxx". You must replace 'a' from S1 with 'b' so that S1 = "xbxb".
*/

func Anagram(s string) int32 {
	val1 := s[:len(s)/2]
	val2 := s[len(s)/2:]
	var count int32 = 0
	if len(s)%2 != 0 {
		return -1
	} else if string(val1) == string(val2) {
		return 0
	}

	for _, a := range val1 {
		index := strings.Index(val2, string(a))
		if index == -1 {
			count = count + 1
		} else {
			val2 = fmt.Sprintf("%s%s", val2[:index], val2[index+1:])
		}
	}
	return count
}
