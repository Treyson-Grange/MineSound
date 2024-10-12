# Ladies! Gentlemen! Both, and neither!

This is the dumbest thing I have ever programmed. Be warned.

I am in a big minecraft mood currently, and I constantly forget, Minecraft is terrifying. So many erie sounds that keep me on my toes. It keeps me attentive. Minecraft keeps my attention for hours on end, and with my ADHD like brain, that is impressive. I theorize that this effect comes from the soundtrack. (the sound design in general.) If I can keep the same attentive prowess in my minecraft playthrough, while I, I don't know, work? I think I will reach new levels of productivity.

This idea led to this program. I decided it had to be a command line application. No real reason. Fun intro to golang tho. Also a super fun spooky season project. Happy Halloween!

## Installation

The scariest thing of all! No installation instructions! ooooooOOoooooOOOooooh you're shaking out of fear, I know. 👻

## Instructions

- `run go . <min_in_sec> <max_in_sec>`
  - Override for randomization is necessary if you want more time between sounds.
  - By default, it is fast intervals, for the sake of showcase.
- Headphones on, sound up.
- Use the arrow keys to explore your options.
- Select an option with enter.
- Feel free to change your selection at any time.

## Care to expand?

Feel free. If you want to add a soundscape option, these are the steps.

1. Gather mp3s. The point of this program is to be spooky. Keep that in mind.
   1. For ease of access, all mp3s must be named x.mp3 where x is 1-# of mp3s.
2. Put them in a sub directory of the mp3 directory.
3. Naming matters here. The name of the directory must match the code.
4. Everything you need to change is in the initialModel.
   1. Add a string to the `choices` array
   2. Add a description of this choice to the `choicesText` mapping.
   3. Add the number of mp3s that you added.
5. On rerun, your option will appear.

For any other changes, you'll have to fend for yourself.

## Credits

- [Minecraft](https://www.minecraft.net/en-us), obviously.
