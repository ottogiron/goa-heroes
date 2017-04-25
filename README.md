# GOA Heroes
Example of a simple GOA API. Implements an API with information about comics heroes.


## Implemented Endpoints


### Hero by ID

```
GET /heroes/:id
```

### List of available heroes

```
 GET /heroes
```


###  Create a new Hero

```
POST /heroes 
{
    name: <string>
}
```


## Generate Goa  Code 

```bash
goagen bootstrap -d github.com/ottogiron/goa-heroes/design
```