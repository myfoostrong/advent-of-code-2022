# Day 11

## Notes

I cheated for Part 2. 

Realized `int` wasn't going to cut it for size, so tried `float`, `uint64`, and finally `big.Int`, with the latter being able to actually handle the size. But the compute time was way too long, though checking along the provided checkpoints in the challenge prompt showed the algorithm was at least working correctly.

I gave up and ran to the Reddit solution thread, which points out the [Chinese remainder theorem](https://en.wikipedia.org/wiki/Chinese_remainder_theorem) as the way to go.