# mathsyst
conversions and operations between different numerical systems

Many languages, Go included, offer supports for different numerical systems,  
like binary, hexadecimal, and, of course, decimal. mathsyst help with custom  
systems. It is possible to define something like:

```
  alphabet := "0123456"
```
To get a heptary numerical system, or to define something like:  

```
 alphabet := "abcdefg"
```
Where the character "a" represents teh ZERO in this system and so on.  
It is also possible to define a binary system as:  
```
  alphabet := "ft"
```

and so on. It is possiboe to convert from a custom system to decimal and vice-versa:

```
  s := mathsyst.NewSystem("0123")
  s.ToDec("10") // this will return 4
  s.FromDec(4) // this will return "10"
```

It is also possible to perform the basic four operations within a system:

```
  s := mathsyst.NewSystem("0123")
  s.Sum("3", "1") // this will return "10"
  s.Diff("11", "2") // this will return "3"
  s.Mult("2", "2") // this will return "10"
  s.Div("10","2"). // this will return "2"
```

It is also possible to add some padding:

```
  s := mathsyst.NewSystem("abcd")
  s.WithLeadingZeros("d", 5) // this will return "aaaa5"
```

More features to come
