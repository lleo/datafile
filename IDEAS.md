IDEAS
=====

0) Build a fully self contained API to store data blobs and have indexes
   defined all in a single file.
1) Blobs are a singley linked list of sub-blobs where sub-blobs are a
   continguous range of bytes an integer number of pages wide. However,
   the allocation algorithms favors a single sub-blobs sufficiently large
   until the allocation fails then trying the linked list scenario.


## Golang API to define data struct and build indexes.

```go
 ```

## File Structure:

	byte offset
	range          name          description
	--------------+-------------+-------------
	0-hdrsz       | header0     | page size/root
	hdrsz-2*hdrsz | header1     | ditto
	2*hdrsz-end   | data        |

## Header Structure

```go
type PageId uint64
```

	byte offset
	range           type      name          description
	--------------+---------+-------------+-------------
	0-3           | uint32  | pageSize
	4-11          | PageId  | rootPageId
	12-15         | uint32  | rootPageOffset
	16-hdrsz      | []byte  | unused
