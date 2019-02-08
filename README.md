# Project: datafile or datastore

## Goal:

Implement a (persistent & functional) (list of data structs & indexes over
those data structs).

## Subprojects:

	* backing store for the above " (persistent & functional) (list of data structs AND indexes over those data structs)"
	* in memory data stucture to access the top level goal which can not or
	should not be loaded into memory.
	* [lower priority] a single file access to the above (list of data structs
	& indexes over those data structs) which may be logically separated into
	coupled files aka a mini-filesystem.

## Deliverables:

	* golang API and/or toolkit for defining the data structs and which indexes
	an type of indexes to build/maintain.
	* golang library to access the above defined artifact
