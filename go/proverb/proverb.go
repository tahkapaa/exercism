/*
Package proverb - For want of a horseshoe nail, a kingdom was lost, or so the saying goes.

Given a list of inputs, generate the relevant proverb. For example, given the list `["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]`, you will output the full text of this proverbial rhyme:

```text
For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.
```

Note that the list of inputs may vary; your solution should be able to handle lists of arbitrary length and content. No line of the output text should be a static, unchanging string; all should vary according to the input given.
*/
package proverb

import (
	"fmt"
)

// Proverb creates the proverb.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	proverb := make([]string, len(rhyme))
	for i := 1; i < len(rhyme); i++ {
		proverb[i-1] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i])
	}
	proverb[len(rhyme)-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return proverb
}
