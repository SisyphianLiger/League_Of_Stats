# Refactoring and One More Test

## General Design
So here was a thought, I know that my function before is able to access and test the Database, but here is what I do not know, does the memory that I need from the DB (types) match a in memory struct I need. In other words, if I wanted to say, use this struct to create some sort of JSON object to be sent to my front end...aka the next step. How would I be certain that the infor from the DB would in fact, be the data specified. 

We test it.
## Testing
Lets look at the "new test", I say new because there is a bit more to chew on but it's really just an extention of the first test, it's just that now we make an in memory struct that takes the data from the DB and puts it in main memory. Let's look at the code.

```Go

func TestPutPlayerCardToInMemStruct(t *testing.T) {
	type PlayerCards struct {
		Name         string
		Age          uint16
		Birthday     time.Time
		Country      string
		Team         string
		Position     string 
		WorldsTitles uint16
	}

	var players []PlayerCards
	// There is no .env here...we will get to that in Problems	
	dbURL := utils.EnvironmentVarExists("DB_URL")
	db := utils.OpenDB("postgres", dbURL)

	q := Queries{
		db: db,
	}
	// Make a note about Background for testing
	ctx := context.Background()
	p, err := q.FindAllPlayerCards(ctx)

	if err != nil {
		t.Fatalf("The Database could not be queried: %v", err)
	}

	for _, ps := range p {
		add_player := PlayerCards {
			Name: ps.Name,
			Age: uint16(ps.Age.Int16),
			Birthday: ps.Birthday.Time,
			Country: ps.Country.String,
			Team: ps.Team.String,
			Position: ps.Position.String,
			WorldsTitles: uint16(ps.WorldsTitles.Int16),
		}

		players = append(players, add_player)
	}

	if len(players) != 2 {
		t.Fatalf("The Query should have 2 in DB, found %d", len(players))
	}
}
```

So what is this testing. Well, if we for example wanted to query the Database that would then pull all of the players for my website, what we would need to do is test the length of structs that have been created. But for this to happen, we need a slice of structs, such that we can hold the `PlayerCards` struct. So what happens here, well just like the test before we make sure that we have access to the DB, and then we make a query which returns with sqlc a slice of db cards. So, we append the copied values into our `PlayerCards` struct and append this to the players slice or a slice of `PlayerCards[]`. Pretty straight forward. 

Now I know for certain that the length of this query is 2, I added the data myself, however, thinking more generally, it would probably be a good idea to either know the length before hand or test in some other way. For instance we could test a single hit on the database and confirm with an anonymous struct has the same values. Or alternatively, we could make a count(*) query for the table and see if it matches the length of the in memory slice struct. Either way lots of ways to crack the egg here.

## Problems 
So there was one problem that occured with these two tests here. And that that stoopud ".env" file from last time...yea it's back again with a vengence. See, it turns out that when you make a test that requires the ".env" file you only ever need to run it once. Running this file twice will in fact, lead to an error. So what do?

```Go
func TestPreLoadEnvFile(t *testing.T) {

	dirErr := os.Chdir("../../../")
	if dirErr != nil {
		t.Logf("If dirErr does not load here we are in trouble %v", dirErr)
	}
	utils.LoadEnvFile(".env")
}
```

Simple...we run this function first which gives us our in memory .env variable(s). And then we can extract this test out from the other tests. A lovely fix whenever we can make our tests more modular. 

Anyways, that will do it for this chapter, in the next chapter I am gonna take this newfound data and send it to the front end...so should be fun!
