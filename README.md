
## Form3 Go Exam

### Developer

David Kotlirevsky

### Current Go Seniority

Just a beginner


### Development Steps

Main IDEA: learn from TDD how to solve the exercise
* Create a simple test
* Make it run 
* Refactor later.


Steps of developing
* Read Your definitions
* Check API documentations
* Start project with some basics folders
* Create some test to Know from where I can start
* Check all APIs requirements from postman ( Added postman file into project)
* Create a bunch of accounts and start for fetching one of them
  * Lear earn how to test / consume and validate it.  
  * Check Fail results with Postman to validate them with tests
* Work with create
* Work With delete
* Enhance code and refactor being based on tests working

 
Considerations
* How to run Test locally and running against console // check account api with postman
Change the environment variable :
ACCOUNTAPI_ADDR=http://localhost:8080

* How to test against docker
* start account api
* Check inside de container with ifconfig to get the Ip container
* Change the IP within env variable for account-lib ACCOUNTAPI_ADDR 

ACCOUNTAPI_ADDR=http://<IP Found>:8080

## Test runing by console:
![image](https://user-images.githubusercontent.com/1593856/204144212-68d48ce6-8653-4492-8987-b483b94b7e84.png)

## Test runing with docker console:
![image](https://user-images.githubusercontent.com/1593856/204279426-ef4de62d-c5b7-4823-93dd-cfdc478ffcf9.png)

