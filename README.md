# Project: datafile or datastore

## Goal:

Implement a (persistent & functional) (list of data structs & indexes over
those data structs).

## Subprojects:

* A backing store for the above Goal.
* An in memory data stucture to access backing store.
* [lower priority] a single file access to the backing store, which may be
  logically separated into coupled files aka a mini-filesystem.

## Deliverables:

* Golang API and/or toolkit for defining an Artifact that determines the data
  structs and which indexes and type of indexes to build/maintain.
* Golang library to access the above defined Artifact.
