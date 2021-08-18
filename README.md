# assignme

a task assigning platform that allows for others on your team to assign actionable tickets demanding accountability

### TODO List

1. The following need to be returned in a get all request
    1. `totalRows, perPage, pageOptions, filter, filterOn, sortBy, sortDesc`
2. Where to store AWS credentials
3. Logging
4. refactor - clean

### Feature List
* Start timer for when a ticket moves from TODO to In-Progress
* Table to capture the life-cycle of a ticket - audit log
* analytics surrounding tickets

### Best Practices

1. `Accept interfaces, return structs`:
    1. ```go
         func funcName(a INTERFACETYPE) CONCRETETYPE
       
       type Person struct{
        firstName string
        lastName string
       }
       
       func loadPerson(r io.Reader) (Person, error) {

            var p Person

            err := json.NewDecoder(r).Decode(&p)

            if err != nil {

                return p, err

            }

            return p, err

       }
       ```
    2. The main focus here is being lazy and follow Postel's
       Law `Be conservative with what you do, be liberal with what you accept`
       Idiomatic go would define this as a non-Java way (define an interface, define one type to fulfill the interface,
       define the methods that satisfies implementation of the interface)
       this is accomplished by:
        1. Define the types
        2. Define the interface at point of use
