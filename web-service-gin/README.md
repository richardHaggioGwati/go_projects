# Restful API

## Definition

API stands for Application Program Interface. It's main function is to communicate with another application in order to receive, edit
or delete data.

## Use Case

This API will provides access to a store selling vintage recordings on vinyl. So it needs to provide endpoints through which a client can get and add albums for users.

## Endpoints

/albums

- GET - Get all the list of albums and returns JSON
- POST - Add a new album from request data sent as JSON

/albums/:id

- GET - Get an unique album by ID, returning the album data as JSON

## Data Creation

For demonstration purposes, firstly I will store the data in memory and then move on to store the data in a backend for a more realistic API.

The string data in memory will mean that the data will be lost each time the server is stopped. Then recreated when the server starts.
