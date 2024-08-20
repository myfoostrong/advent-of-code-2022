# Day 2 Challenge

## Retrospective

I'm pretty sure I'm failing with Go maps somehow, hence the wonky instatiation [Line 17](./day2.go#L17)

Should have calculated the outcome map in code, not by hand. That way it could solve for a different strategy guide mapping, or maybe games with more options [Line 35](./day2.go#L35)

Though I'm not sure how RPSX would work... as kids people would throw in bomb or dynamite with a thumbs up fist then smack your hand.

### Efficiency Guess (YMMV)
O(n) time, where n is the length of the input list of "plays".
O(m^2) memory, where m is the number of "play" options. If locked to RPS, that's a constant of 3^2 or 9, and three characters per "play row", or 12 bytes, plus what's need by the imported libraries.

## Working Notes
* [Effective Go Maps](https://go.dev/doc/effective_go#maps)
* [Go 101: Arrays, Slices and Maps in Go](https://go101.org/article/container.html)
* [SO: Split whitespace](https://stackoverflow.com/questions/13737745/split-a-string-on-whitespace-in-go)
* [SO: Add data to maps of maps](https://stackoverflow.com/questions/64128440/how-to-add-data-to-a-map-of-maps-in-golang)

