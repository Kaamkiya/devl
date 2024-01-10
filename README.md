main file will be 		  `/main.go`
file running quizzes will be      `/quiz.go`
quizzes will be stored in 	  `/quizzes/<language>.txt`
file running challenges will be   `/challenge.go`
challenges will be stored in      `/challenges/<challenge#>.<language>`


Structure:
```
/cli-learn
    |____ main.go
    |____ challenge.go
    |____ quiz.go
    |____ Makefile
    |____ quizzes/
	     |____ python.txt
             |____ go.txt
    |____ challenges/
             |____ 1.py
             |____ 1.c
```

`cli-learn`: main project directory
`main.go`: main file. This will be passed arguments and such. 
`quiz.go`: if the requested thing is a quiz, run the quiz
