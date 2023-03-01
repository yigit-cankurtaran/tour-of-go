package main

// import of the exercise test didn't work.
// i'll run it in the playground to check if it works.
import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	wc := WordCount(`[Intro: Lil Uzi Vert]
	Yes, I'm a rich nigga
	(Dubba-AA flex)
	Yes, I'm a rich nigga
	Yes, I'm a rich nigga
	
	[Chorus: Lil Uzi Vert]
	Yes I'm a rich nigga, rich nigga, rich nigga
	I heard that you is a bitch nigga, bitch nigga
	Keep a 30 stick, nigga
	Bitches on my dick, nigga
	I patiently waited, for your ass to trip, nigga
	Fuck nigga, flip nigga
	You know that I'm lit, nigga
	Walkin' with the bands, so I could pay for the hit, nigga
	Diamond chain choker bitch, so I could not trip, nigga
	They changed on me so many times, I cannot trust no one
	[Verse 1: Lil Uzi Vert]
	No worry, no you is a bitch, so you ain't goin' up now
	You ain't never bust a strap so we call that a Rust Gun
	I'm the type of nigga eat the ice when the whole cup done
	Diamonds they so big up in my chain, they look like mushrooms
	These niggas keep talking 'bout that smoke, but they don't want none
	Heavy artillery on me, heard that you got one gun
	You know my chain is gon' shine even when the whole sun-done
	If I wear all this shit at the same time, I'd weigh a whole ton
	
	[Chorus: Lil Uzi Vert]
	Yes I'm a rich nigga, rich nigga, rich nigga
	I heard that you is a bitch nigga, bitch nigga
	Keep a 30 stick, nigga
	Bitches on my dick, nigga
	I patiently waited, for your ass to trip, nigga
	Fuck nigga, flip nigga
	You know that I'm lit, nigga
	Walkin' with the bands, so I could pay for the hit, nigga
	Diamond chain choker bitch, so I could not trip, nigga
	They changed on me so many times, I cannot trust no one
	
	[Verse 2: YoungBoy Never Broke Again]
	Where the whole team?
	Wanted you to know that I copped it
	Niggas they know that I stay with the rocket
	Still they be hatin', but know I ain't stoppin'
	Shit we do these rappers copy
	Sauced up, I'm a trendin' topic
	Step on the scene, they lookin' and watchin'
	Throw up that gang like fuck what the hollering (gang)
	Used to be jumpin' them fences
	Put my face down, and now that I'm in it
	You know I stay strapped with that glizzy
	17, I got locked in that prison
	Wait, hold on one minute (hol' on)
	You go wit' your moves, you know you gon' get it
	Make me drop a 50
	Leavin' the scene with blood on my tennis
	You might also like
	Time Out
	YoungBoy Never Broke Again
	Locked & Loaded
	YoungBoy Never Broke Again
	Thug Cry
	YoungBoy Never Broke Again
	[Chorus: Lil Uzi Vert]
	Yes I'm a rich nigga, rich nigga, rich nigga
	I heard that you is a bitch nigga, bitch nigga
	Keep a 30 stick, nigga
	Bitches on my dick, nigga
	I patiently waited, for your ass to trip, nigga
	Fuck nigga, flip nigga
	You know that I'm lit, nigga
	Walkin' with the bands, so I could pay for the hit, nigga
	Diamond chain choker bitch, so I could not trip, nigga
	They changed on me so many times, I cannot trust no one
	`)
	fmt.Println(wc)
	// prints map[hello:2 world:1]
	// wc.Test(WordCount)
	// returns an error normally because local file doesn't have a wc import
	// uncomment it to run on the playground.
}

/*
the result i got is:

​
PASS
 f("I am learning Go!") =
  map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
PASS
 f("The quick brown fox jumped over the lazy dog.") =
  map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
PASS
 f("I ate a donut. Then I ate another donut.") =
  map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
PASS
 f("A man a plan a canal panama.") =
  map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

  ALL PASSES!

*/
