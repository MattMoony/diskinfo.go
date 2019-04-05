# diskinfo.go
_Stay informed about your storage situation_

---

Just a fun, little thing. You can use several functions
to get info about the storage space on your computer's 
disk.

*Note:* It's able to get the diskinfo on both Windows and
Linux.

---

## Setup

In order to use _diskinfo_ in your project, you need to
import it first: 

```Go
import "diskinfo"

// ... 
```

## Usage 

There's only one function you need to pay attention to,
the `GetDiskInfo(...)` function.

It will return a struct of type `DiskInfo` which has the
following structure:

```Go

type DiskInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

```

And you should use the function like this:

```Go

// on windows ... 
c_part_info := GetDiskInfo("C:")

// on linux ... 
r_part_info := GetDiskInfo("/")

```

## Conclusion

I hope you enjoy this project and that
it is in some way, shape or form useful to you.

If you have any questions, [here](https://twitter.com/Matthia23184857) is my twitter.

---

... MattMoony (April, 2019)