
# deck of cards simulation  
An API to simulate a deck of cards  
  
Install  
-------  
  
From application's home directory run following commands   

```  
go get
```  
  
Test  
-----  
  
From application's home directory run following commands  
  
```  
cd test/unittests && go test -v  
cd test/integrationtests && go test -v  
```  
  
Starting Server  
-----  
  
From application's home directory run following command
  
```  
go run main.go  
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