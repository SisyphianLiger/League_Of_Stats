# Refactoring and Pain

## General Design
I had a lot of fun with this project. And by that I mean sometimes you forget just how lovely and wonderful it is when all the plumbing is done for you and you can just Jump in! 

What do I mean you ask? Well I had a simple task. That was, could I make a unit test for my DB. What pray tell would I be testing for? Well mainly does my query succeed?

I know what you are thinking, Ryan, that sounds ostensibly simply, sure you could do that esily no? Yes, but in my adventures to become a better programme I learned quite a bit about Golang specific testing structure and a whole lot of pitfalls I would like to share. First...lets take a look at the test in question.
## Testing

```Go

func TestFindAllPlayerCardsSuccess(t *testing.T) {

	dirErr := os.Chdir("../../../")
	if dirErr != nil {
		t.Logf("If dirErr does not load here we are in trouble %v", dirErr)
	}
	
	utils.LoadEnvFile(".env")
	dbURL := utils.EnvironmentVarExists("DB_URL")
	db := utils.OpenDB("postgres", dbURL)

	q := Queries{
		db: db,
	}
	// Make a note about Background for testing
	ctx := context.Background()
	_, err := q.FindAllPlayerCards(ctx)

	if err != nil {
		t.Fatalf("The Database could not be queried: %v", err)
	}

}
```

Wow! What a terrific and simple testing funciton using `t *testing.T`. This file is located in the database folder section. My idea here was simple. I am making queries, then I am doing something with that in memory. So how can I make sure, for example that if I say, query a database, that I am guarenteed that what I have asked for exists?

Well, simply put I can use the function that gets all the DB information, and see if it was successful. (The next step would obviously be to compare the DB format to a struct, but for now I think this will suffice). Now, there are a couple of interesting notes I found with Golangs testing methodology. Firstly, the test file needs to have `_test` in it. Oddly specific but sure Golang live your best life. Next, the Function in the `_test.go` needs to lead with `Test...(){}`. Ok Golang, you really like particular stuff. 

All of these specific rules are fine in my opinion and harmless if they enforce testing. There was one thing from learning how to test that I found great, you can see the following struct can hold a blank context. 

```Go
type Context struct {
	ctx context.Context
}
```

This is pretty nice and allows for good flexibility. But! How then do we call the context into question!? We use `context.Background()`, you can think of it as a dummy context that allows for unit testing of our functions.

## Problems
If you really took a look at my function, you may notice this line.

```Go
	dirErr := os.Chdir("../../../")
	if dirErr != nil {
		t.Logf("If dirErr does not load here we are in trouble %v", dirErr)
	}
```
Yea...that is right, the problem I was having with this test, was loading the `.env` file,
 I don't know about y'all but I don't want to just leak my DB_URL Path. However, figuring out how I could get to the Top level (outside of Golang) proved difficult until I read about `os.Getwd()` or get working directory. With that I realized that just like when you `cd` too deep, I needed to go three levels up, such that my environment variables could be read to memory and passed to the `OpenDB` function. For my own personal pride...I do not want to mention how long that took me to figure out...but anyways, now I have a test, that can test with .env variables...Yay Me!
