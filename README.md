# Commands
## add
add command will add `<file>` file to the given drive `<drive>`
```bash
qpfs add <file> <drive>
```

For example:
```bash
qpfs add readme.md mydrive
```
This operation will add file `readme.md` to an existing drive named `mydrive`.
If the drive or the file does not exist, a corresponding error message will show up.

## get
get command will download file with given name `<fname>` from the `<drive>` to `<path>`

```bash
qpfs get <drive>/<fname> <path>
```

For example:
```bash
qpfs get mydrive/readme.md .
```
This example will download `readme.md` from `mydrive` to current directoy.

## ls
ls will list all existing files under given `<drive>`

```bash
qpfs ls <drive>
```

## rm
rm command unpins the file with given `<fname>` on `<drive>`. Note that the file may not be removed immediately unless it has been garbage collected by the running IPFS daemon.
```bash
qpfs rm <drive>/<fname>
```

## drive
### new
It will create a new drive with given `name`.
```bash
qpfs drive new <name>
```

### ls
It will list all current created driver.
```bash
qpfs drive ls
```

### stat
It will stat the infomration about an existing drive with given `name`.
```bash
qpfs drive stat <name>
```

