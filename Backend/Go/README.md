# Go note


### array vs slice
- cap
- len
- append


#####
pass an array to function will result a new copy of array

##### Memory Leak problem
slice a existing array or a slice will cause underline array keep unnecessary elements as long as slice keep.



### type

```pointer receiver
func (p *int) method() { }
```

```value receiver
func (p int) method() { }
```

### Mutex

```
type T Struct {
    lock Sync.Mutex
}
var t T
t.lock.Lock()
t.lock.Unlock()
```


### init function
when app runs, func init will be run ahead of main func.
It could have multiple init functions.


### fmt.Print
This will take Stringer interface(String method) to print.
However, if there is error interface(Error method) implement. error interface has higher priority.