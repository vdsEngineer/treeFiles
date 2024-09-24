## Review 

> treeFiles - minimalistic, fast and customizable cli applications for viewing files in the terminal in the form of a tree

<div align="center">
  <img alt="superfile LOGO" src="/asset/example_2.png" />
</div>

## Installion 

**to install, you will need the latest version of go**

```bash 
git clone https://github.com/vdsEngineer/treeFiles.git
cd treeFiles
go build -o ./bin/
```

Add the binary file to your $PATH, e.g., in `/usr/local/bin`:
```bash
sudo mv ./bin/treeFiles /usr/local/bin/ # For Linux
```

**It is important to note that it is much more convenient to use this utility using alias**

For the place of this:

```bash 
treeFiles -e 3 -l 10 -rc 35 -lc 63 -ic 212
```

That is:
```bash 
l # treeFiles -e 1 
ll # treeFiles -e 2
lll # treeFiles -e 3
```

Alias help for various OS can be found on the internet

## Examples

| parameter                | description                          |
| ------------------------ | ------------------------------------ |
| `-e or --enclosure`      | the nesting level for the tree       |
| `-l or --limitFiles`     | limit of files to show in the tree   |
| `-rc or --rootFontColor` | font color for the root of the tree  |
| `-lc or --colorLine`     | color for the lines of the tree      |
| `-ic or --itemFontColor` | font color for the items of the tree |

**parameters -rc -lc -ic  they can accept a color in the following format #FFFFFF  but it is worth transmitting without #, that is FFFFFF **

```bash 
treeFiles --enclosure 1 -l 10 -rc 35 -lc 63 -ic 120
```

<div align="center">
  <img alt="superfile LOGO" src="/asset/example_1.png" />
</div>

```bash 
treeFiles --enclosure 1 -l 10 -rc 35 -lc 63 -ic 220
```

<div align="center">
  <img alt="superfile LOGO" src="/asset/example_2.png" />
</div>


## Supported Systems
- [x] Linux
- [x]  MacOS
- [x]  Windows (Not fully supported yet)

