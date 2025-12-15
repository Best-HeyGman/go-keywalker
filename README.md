# go-keywalker
Generate every possible keyboard walk on a given keyboard or check if a string is a valid keyboard walk.

## What is a keyboard walk?
Sometimes, when people want a "random" string of characters, they do what is called a keyboard walk. Don't know what to call your new movie? asdf. Very random. Passwords should be random, right? Oh no...

# Usage
## keywalk-gen
Generate all keyboard walks on a given keyboard.  
`keywalk-gen -k <path/to/your/keymap.nsk> -m <min_length> -M <max_length>`  
This will output every possible keywalk on the given keyboard and within the given lengths to stdout.

## keywalk-check
Pipe in your text file and print out every line that is a valid keyboard walk. You can also print every line that is **not** a keyboard walk by using `-r` (reverse).  
`cat wordlist.txt | keywalk-check -k <path/to/your/keymap.nsk>`

# The Keyboard Map file format (.nsk)
It's a `\n`-separated UTF-8 text file.  
The file format I created to map out a keyboard is as follows:  
- Every line of the text file lists one key and all of its adjacent keys.
- The first character of a line is the key we are mapping out.
- Every character afterward is a neighbouring key on the keyboard.

I have chosen this format to make mapping a new keyboard as easy as possible. I wanted a format where you don't have to look up from the keyboard while mapping it, because the task is boring enough for you to get easily distracted but complex enough that you will make mistakes if you do.

## Mapping out a new keyboard
I suggest doing it this way:
1. Start in one corner of the keyboard.
2. Press a key and then all the keys that are adjacent to that key.
3. If you have all of them, press Enter to make a new line and proceed to the next key.
4. Do this until you have mapped out all keys on the keyboard.

That's it. If you want to add a keyboard map, feel free to do so in a pull request. If you think your keyboard is fairly representative of the standard keyboard of a language, you can call it something like "de-full" or "de-tkl" or "us-full" or something. If you have some exotic layout, try to come up with a good name.

# libkeywalk
I made a library that is shared by these two programs. It can only do what I needed for the current functionality, but if you are interested and want to use it, for example, to check password quality in your program or something, just let me know and I'll make a real Go package out of it (however that works) and document it thoroughly.

# Contributing
Contributing a keyboard map should be very easy; the format is described above.  
The program can do everything I wanted it to do, but if you want some more features, feel free to either make a pull request (but please make sure you work with []rune instead of simple strings) or just ask if I can add it and I probably will.

## My ideas
I also have a couple of ideas myself, but I don't really need them personally, so I haven't added them yet. If someone else needs them, I'll definitely add them:  
- Also map alternative key assignments like capital letters and the like
	+ This is also why I made an extra file ending for the keyboard maps (.nsk), so that I can add a new format for alternative key assignments.
		* The idea is to give, for a single keyboard model, both files the same name and differentiate them by their ending, so that I don't have to make the .nsk format more complex but still keep both together.
		* `.nsk` stands for "newline separated key...something". I didn't think it completely through.
		* The new file format I would call something like `.nsc` (newline separated caps...something).
		* On a German keyboard, the mapping of the e key could look like this:
			- eE€
- Multithreading
	+ Should be very easy to implement, because every walk done from a particular starting key is independent of each other. 
	+ I've already split up the functions in libkeywalk in a way that I can easily add this feature.
	+ But the performance without multithreading is already really good, so I haven't bothered yet.