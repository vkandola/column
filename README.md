<p align="center">
<img width="330" height="110" src=".github/logo.png" border="0" alt="kelindar/column">
<br>
<img src="https://img.shields.io/github/go-mod/go-version/kelindar/column" alt="Go Version">
<a href="https://pkg.go.dev/github.com/kelindar/column"><img src="https://pkg.go.dev/badge/github.com/kelindar/column" alt="PkgGoDev"></a>
<a href="https://goreportcard.com/report/github.com/kelindar/column"><img src="https://goreportcard.com/badge/github.com/kelindar/column" alt="Go Report Card"></a>
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
<a href="https://coveralls.io/github/kelindar/column"><img src="https://coveralls.io/repos/github/kelindar/column/badge.svg" alt="Coverage"></a>
</p>

## Columnar In-Memory Store with Bitmap Indexing

This package contains a **hihg-performance, columnar, in-memory storage engine** that supports fast querying, update and iteration with zero-allocations and bitmap indexing.

## Features

 * Optimized, cache-friendly **columnar data layout** that minimizes cache-misses.
 * Optimized for **zero heap allocation** during querying (see benchmarks below).
 * Optimized **batch updates/deletes**, an update during a transaction takes around `12ns`.
 * Support for **SIMD-enabled filtering** (i.e. "where" clause) by leveraging [bitmap indexing](https://github.com/kelindar/bitmap).
 * Support for **columnar projection**  (i.e. "select" clause) for fast retrieval.
 * Support for **computed indexes** that are dynamically calculated based on provided predicate.
 * Support for **concurrent updates** per-column (e.g. 2 goroutines can update 2 columns concurrently).
 * Support for **transaction isolation**, allowing you to create transactions and commit/rollback.
 * Support for **expiration** of rows based on time-to-live or expiration column.
 * Support for **atomic increment/decrement** of numerical values, transactionally.

## Documentation

The general idea is to leverage cache-friendly ways of organizing data in [structures of arrays (SoA)](https://en.wikipedia.org/wiki/AoS_and_SoA) otherwise known "columnar" storage in database design. This, in turn allows us to iterate and filter over columns very efficiently. On top of that, this package also adds [bitmap indexing](https://en.wikipedia.org/wiki/Bitmap_index) to the columnar storage, allowing to build filter queries using binary `and`, `and not`, `or` and `xor` (see [kelindar/bitmap](https://github.com/kelindar/bitmap) with SIMD support). 

- [Collection and Columns](#collection-and-columns)
- [Querying and Indexing](#querying-and-indexing)
- [Iterating over Results](#iterating-over-results)
- [Updating Values](#updating-values)
- [Expiring Values](#expiring-values)
- [Transaction Commit and Rollback](#transaction-commit-and-rollback)
- [Complete Example](#complete-example)
- [Benchmarks](#benchmarks)
- [Contributing](#contributing)

## Collection and Columns

In order to get data into the store, you'll need to first create a `Collection` by calling `NewCollection()` method. Each collection requires a schema, which can be either specified manually by calling `CreateColumn()` multiple times or automatically inferred from an object by calling `CreateColumnsOf()` function. 

In the example below we're loading some `JSON` data by using `json.Unmarshal()` and auto-creating colums based on the first element on the loaded slice. After this is done, we can then load our data by inserting the objects one by one into the collection. This is accomplished by calling `Insert()` method on the collection itself repeatedly.

```go
data := loadFromJson("players.json")

// Create a new columnar collection
players := column.NewCollection()
players.CreateColumnsOf(data[0])

// Insert every item from our loaded data
for _, v := range data {
	players.Insert(v)
}
```

Now, let's say we only want specific columns to be added. We can do this by calling `CreateColumn()` method on the collection manually to create the required columns.

```go
// Create a new columnar collection with pre-defined columns
players := column.NewCollection()
players.CreateColumn("name", column.ForString())
players.CreateColumn("class", column.ForString())
players.CreateColumn("balance", column.ForFloat64())
players.CreateColumn("age", column.ForInt16())

// Insert every item from our loaded data
for _, v := range loadFromJson("players.json") {
	players.Insert(v)
}
```

## Querying and Indexing

The store allows you to query the data based on a presence of certain attributes or their values. In the example below we are querying our collection and applying a *filtering* operation bu using `WithValue()` method on the transaction. This method scans the values and checks whether a certain predicate evaluates to `true`. In this case, we're scanning through all of the players and looking up their `class`, if their class is equal to "rogue", we'll take it. At the end, we're calling `Count()` method that simply counts the result set.

```go
// This query performs a full scan of "class" column
players.Query(func(txn *column.Txn) error {
	count := txn.WithValue("class", func(v interface{}) bool {
		return v == "rogue"
	}).Count()
	return nil
})
```

Now, what if we'll need to do this query very often? It is possible to simply *create an index* with the same predicate and have this computation being applied every time (a) an object is inserted into the collection and (b) an value of the dependent column is updated. Let's look at the example below, we're fist creating a `rogue` index which depends on "class" column. This index applies the same predicate which only returns `true` if a class is "rogue". We then can query this by simply calling `With()` method and providing the index name. 

An index is essentially akin to a boolean column, so you could technically also select it's value when querying it. Now, in this example the query would be around `10-100x` faster to execute as behind the scenes it uses [bitmap indexing](https://github.com/kelindar/bitmap) for the "rogue" index and performs a simple logical `AND` operation on two bitmaps when querying. This avoid the entire scanning and applying of a predicate during the `Query`.

```go
// Create the index "rogue" in advance
out.CreateIndex("rogue", "class", func(v interface{}) bool {
	return v == "rogue"
})

// This returns the same result as the query before, but much faster
players.Query(func(txn *column.Txn) error {
	count := txn.With("rogue").Count()
	return nil
})
```

The query can be further expanded as it allows indexed `intersection`, `difference` and `union` operations. This allows you to ask more complex questions of a collection. In the examples below let's assume we have a bunch of indexes on the `class` column and we want to ask different questions.

First, let's try to merge two queries by applying a `Union()` operation with the method named the same. Here, we first select only rogues but then merge them together with mages, resulting in selection containing both rogues and mages.

```go
// How many rogues and mages?
players.Query(func(txn *Txn) error {
	txn.With("rogue").Union("mage").Count()
	return nil
})
```

Next, let's count everyone who isn't a rogue, for that we can use a `Without()` method which performs a difference (i.e. binary `AND NOT` operation) on the collection. This will result in a count of all players in the collection except the rogues.

```go
// How many rogues and mages?
players.Query(func(txn *Txn) error {
	txn.Without("rogue").Count()
	return nil
})
```

Now, you can combine all of the methods and keep building more complex queries. When querying indexed and non-indexed fields together it is important to know that as every scan will apply to only the selection, speeding up the query. So if you have a filter on a specific index that selects 50% of players and then you perform a scan on that (e.g. `WithValue()`), it will only scan 50% of users and hence will be 2x faster. 

```go
// How many rogues that are over 30 years old?
players.Query(func(txn *Txn) error {
	txn.With("rogue").WithFloat("age", func(v float64) bool {
		return v >= 30
	}).Count()
	return nil
})
```

## Iterating over Results

In all of the previous examples, we've only been doing `Count()` operation which counts the number of elements in the result set. In this section we'll look how we can iterate over the result set. In short, there's 2 main methods that allow us to do it:

 1. `Range()` method which takes in a column name as an argument and allows faster get/set of the values for that column.
 2. `Select()` method which doesn't pre-select any specific column, so it's usually a bit slower if you're actually selecting the data. 

Let's first examine the `Range()` method. In the example below we select all of the rogues from our collection and print out their name by using the `Range()` method and providing "name" column to it. The callback containing the `Cursor` allows us to quickly get the value of the column by calling `String()` method to retrieve a string value. It also contains methods such as `Int()`, `Uint()`, `Float()` or more generic `Value()` to pull data of different types.

```go
players.Query(func(txn *Txn) error {
	txn.With("rogue").Range("name", func(v column.Cursor) bool {
		println("rogue name ", v.String()) // Prints the name
		return true
	})
	return nil
})
```

Now, what if you need two columns? The range only allows you to quickly select a single column, but you can still retrieve other columns by their name during the iteration. This can be accomplished by corresponding `StringAt()`, `FloatAt()`, `IntAt()`, `UintAt()` or `ValueAt()` methods as shown below.

```go
players.Query(func(txn *Txn) error {
	txn.With("rogue").Range("name", func(v column.Cursor) bool {
		println("rogue name ", v.String())    // Prints the name
		println("rogue age ", v.IntAt("age")) // Prints the age
		return true
	})
	return nil
})
```

Now, what if you need to quickly delete all of the rogues from the collection? In this case `Select()` method comes in handy specifically because it does not require pre-selected columns and it will be slightly faster to do a batch deletion with it. In the example below we filter all of the rogues and delete them. 

```go
players.Query(func(txn *Txn) error {
	txn.With("rogue").Select(func(v column.Selector) bool {
		v.Delete()
		return true
	})
	return nil
})
```

## Updating Values

In order to update certain items in the collection, you can simply call `Range()` method and the corresponding `Cursor`'s `Update()` or `UpdateAt()` methods that allow to update a value of a certain column atomically. The updates won't be directly reflected given that the store supports transactions and only when transaction is commited, then the update will be applied to the collection. This allows for isolation and rollbacks.

In the example below we're selecting all of the rogues and updating both their balance and age to certain values. The transaction returns `nil`, hence it will be automatically committed when `Query()` method returns.

```go
players.Query(func(txn *Txn) error {
	txn.With("rogue").Range("balance", func(v column.Cursor) bool {
		v.Update(10.0)        // Update the "balance" to 10.0
		v.UpdateAt("age", 50) // Update the "age" to 50
		return true
	}) // Select the balance
	return nil
})
```

In certain cases, you might want to atomically increment or decrement numerical values. In order to accomplish this you can use the provided `Add()` or `AddAt()` operations of the `Cursor` or `Selector`. Note that the indexes will also be updated accordingly and the predicates re-evaluated with the most up-to-date values. In the below example we're incrementing the balance of all our rogues by *500* atomically.

```go
players.Query(func(txn *Txn) error {
	txn.With("rogue").Range("balance", func(v column.Cursor) bool {
		v.Add(500.0) // Increment the "balance" by 500
		return true
	})
	return nil
})
```

## Expiring Values

Sometimes, it is useful to automatically delete certain rows when you do not need them anymore. In order to do this, the library automatically adds an `expire` column to each new collection and starts a cleanup goroutine aynchronously that runs periodically and cleans up the expired objects. In order to set this, you can simply use `InsertWithTTL()` method on the collection that allows to insert an object with a time-to-live duration defined.

In the example below we are inserting an object to the collection and setting the time-to-live to *5 seconds* from the current time. After this time, the object will be automatically evicted from the collection and its space can be reclaimed.

```go
players.InsertWithTTL(map[string]interface{}{
	"name": "Merlin",
	"class": "mage",
	"age": 55,
	"balance": 500,
}, 5 * time.Second) // The time-to-live of 5 seconds
```

On an interestig node, since `expire` column which is automatically added to each collection is an actual normal column, you can query and even update it. In the example below we query and conditionally update the expiration column. The example loads a time, adds one hour and updates it, but in practice if you want to do it you should use `Add()` method which can perform this atomically.

```go
players.Query(func(txn *column.Txn) error {
	return txn.Range("expire", func(v column.Cursor) bool {
		oldExpire := time.Unix(0, v.Int()) // Convert expiration to time.Time
		newExpire := expireAt.Add(1 * time.Hour).UnixNano()  // Add some time
		v.Update(newExpire)
		return true
	})
})
```

## Transaction Commit and Rollback

Transactions allow for isolation between two concurrent operations. In fact, all of the batch queries must go through a transaction in this library. The `Query` method requires a function which takes in a `column.Txn` pointer which contains various helper methods that support querying. In the example below we're trying to iterate over all of the players and update their balance by setting it to `10.0`. The `Query` method automatically calls `txn.Commit()` if the function returns without any error. On the flip side, if the provided function returns an error, the query will automatically call `txn.Rollback()` so none of the changes will be applied.

```go
// Range over all of the players and update (successfully their balance)
players.Query(func(txn *column.Txn) error {
	txn.Range("balance", func(v column.Cursor) bool {
		v.Update(10.0) // Update the "balance" to 10.0
		return true
	})

	// No error, txn.Commit() will be called
	return nil
})
```

Now, in this example, we try to update balance but a query callback returns an error, in which case none of the updates will be actually reflected in the underlying collection.

```go
// Range over all of the players and update (successfully their balance)
players.Query(func(txn *column.Txn) error {
	txn.Range("balance", func(v column.Cursor) bool {
		v.Update(10.0) // Update the "balance" to 10.0
		return true
	})

	// Returns an error, txn.Rollback() will be called
	return fmt.Errorf("bug") 
})
```

You can (but probablty won't need to) call `Commit()` or `Rollback()` manually, as many times as required. This could be handy to do partial updates but calling them too often will have a performance hit on your application.

```go
// Range over all of the players and update (successfully their balance)
players.Query(func(txn *column.Txn) error {
	txn.Range("balance", func(v column.Cursor) bool {
		v.Update(10.0) // Update the "balance" to 10.0
		return true
	})

	txn.Commit() // Manually commit all of the changes
	return nil   // This will call txn.Commit() again, but will be a no-op
})
```

## Complete Example

```go
func main(){

	// Create a new columnar collection
	players := column.NewCollection()

	// index on humans
	players.CreateIndex("human", "race", func(v interface{}) bool {
		return v == "human"
	})

	// index for mages
	players.CreateIndex("mage", "class", func(v interface{}) bool {
		return v == "mage"
	})

	// index for old
	players.CreateIndex("old", "age", func(v interface{}) bool {
		return v.(float64) >= 30
	})

	// Load the items into the collection
	loaded := loadFixture("players.json")
	players.CreateColumnsOf(loaded[0])
	for _, v := range loaded {
		players.Insert(v)
	}

	// This performs a full scan on 3 different columns and compares them given the 
	// specified predicates. This is not indexed, but does a columnar scan which is
	// cache-friendly.
	players.Query(func(txn *column.Txn) error {
		println(txn.WithString("race", func(v string) bool {
			return v == "human"
		}).WithString("class", func(v string) bool {
			return v == "mage"
		}).WithFloat("age", func(v float64) bool {
			return v >= 30
		}).Count()) // prints the count
		return nil
	})

	// This performs a cound, but instead of scanning through the entire dataset, it scans
	// over pre-built indexes and combines them using a logical AND operation. The result
	// will be the same as the query above but the performance of the query is 10x-100x
	// faster depending on the size of the underlying data.
	players.Query(func(txn *column.Txn) error {
		println(txn.With("human", "mage", "old").Count()) // prints the count
		return nil
	})

	// Same condition as above, but we also select the actual names of those 
	// players and iterate through them.
	players.Query(func(txn *column.Txn) error {
		txn.With("human", "mage", "old").Range("name", func(v column.Cursor) bool {
			println(v.String()) // prints the name
			return true
		}) // The column to select
		return nil
	})
}
```

## Benchmarks

The benchmarks below were ran on a collection of *500 items* containing a dozen columns. Feel free to explore the benchmarks but I strongly recommend testing it on your actual dataset.

```
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkCollection/insert-8         27589314       43.05 ns/op     1 B/op     0 allocs/op
BenchmarkCollection/fetch-8          21041593       56.84 ns/op     0 B/op     0 allocs/op
BenchmarkCollection/count-slow-8       109107    11001 ns/op        0 B/op     0 allocs/op
BenchmarkCollection/count-8           9300270      128.6 ns/op      0 B/op     0 allocs/op
BenchmarkCollection/range-8           1871557      641.0 ns/op      0 B/op     0 allocs/op
BenchmarkCollection/select-8          1214799      975.8 ns/op      0 B/op     0 allocs/op
BenchmarkCollection/update-at-8      28573945       41.99 ns/op     0 B/op     0 allocs/op
BenchmarkCollection/update-all-8       184694      6481 ns/op       0 B/op     0 allocs/op
BenchmarkCollection/delete-at-8       2613982      459.1 ns/op      0 B/op     0 allocs/op
BenchmarkCollection/delete-all-8       296901      3762 ns/op       0 B/op     0 allocs/op
```

## Contributing

We are open to contributions, feel free to submit a pull request and we'll review it as quickly as we can. This library is maintained by [Roman Atachiants](https://www.linkedin.com/in/atachiants/)

## License

Tile is licensed under the [MIT License](LICENSE.md).