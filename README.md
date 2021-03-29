# Golang-Unit-Test

Instructions<br/><br/>

1- Open to terminal <br/>
2- go version<br/><br/>
3- go help <br/>
4- 	cd HTMLFolder<br/><br/><br/><br/><br/>
	go run web.go<br/><br/><br/><br/>
	(Now out localhost:3000 port started to listening)<br/><br/><br/>
	(Open a browser and enter this site:)<br/><br/>
	localhost:3000/example.html<br/>
5- Open a new terminal and make sure that you are in the root<br/><br/><br/><br/><br/><br/>
	(I mean first folder when you open to folder. That path should <br/><br/><br/><br/><br/>
	include main.go and main_test.go)<br/><br/><br/><br/>
6- go run main.go<br/><br/><br/>
7- go test<br/><br/>

--------------------------------------------------------------

You can get to libraries using this code in case you need:
$ go get <library>

--------------------------------------------------------------

web.go ---> provide to open localhost and run the html

main.go ---> 	get_html() : accessing to http://localhost:3000/example.html/ 
			and getting HTML codes from there and write in the
			body.txt file. 
		ReadBodyTxt(): Reading body.txt. In this way we will test
			that is url openning correctly. And we will look codes of 
			come from url same with example.html
		ReadExampleHtml(): Reading example.html
		GetSubResult(): It returns us result of the subtract from example.HTML
			We will use this function to test that could we get 
			result of the sub from <p> correctly.
		GetAddResult(): It returns us result of the add from example.HTML
			We will use this function to test that could we get 
			result of the add from <p> correctly.

main_test.go --->	TestCompare(): testing that url is openning correctly. And compares
			that codes are same between body.txt and example.html
			TestSubtract(): testing that result of subtract is true 	
			TestAdd(): testing that result of add is true
		

		
		
			
