# getsurei （月齢）

> I'm pretty sure I'm connected to the moon.
- David Lynch

lunar whimsy.
getsurei estimates primary moon phases and the age of the moon.

intentionally imprecise.

![getsurei in action](getsurei.gif)

## usage

```
now := time.Now()
switch getsurei.Gessou(now) {
case NewMoon:
  fmt.Println("hiding out all night")
case FirstQuarter:
  fmt.Println("still not full")
case FullMoon:
  fmt.Println("bold and bright")
case LastQuarter:
  fmt.Println("whisper goodbye with a quiet thrill")
}

age := getsurei.Getsurei(now)
fmt.Printf("%f days since the new moon\n", age)

fullmoon := getsurei.Next(getsurei.FullMoon, now)
fmt.Printf("net full moon: %s\n", fullmoon.Format("2006.01.02)")
```

## command line

```
$ date
Mon Jan 27 17:27:07 JST 2025
$ getsurei
下弦の月
$ getsurei 2025.01.01
新月
```
