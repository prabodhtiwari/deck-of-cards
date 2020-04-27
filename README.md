
# deck of cards simulation
An API to simulate a deck of cards

Install
-------

From project root's cmd directory run following commands

```
go get
```

Test
-----

From project root run following commands

```
make test
```

Starting Server
-----

From project root run following command

```
make run
```


You can access it at http://127.0.0.1:3000


Requests
-----

Create a new deck

```
http GET :3000/deck/create
http GET :3000/deck/create cards==AS,KD,AC,2C,KH
http GET :3000/deck/create cards==AS,KD,AC,2C,KH shuffle==true
```

Open a deck

```
http GET :3000/deck/open deck_id==6d849b6a-cc06-4c39-b3df-18a3cf93700c
```

Draw a card
```
http GET :3000/deck/draw deck_id==174bef22-6de7-4095-ab73-0df0186a8041 count==5
```

> These commands assume [httpie](https://httpie.org/) is installed
