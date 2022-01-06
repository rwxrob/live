# Twitch Live Streamer Utility/Library in Go

This project seeks to make the life of a Twitch live streamer easier by
taking a modal approach to streaming. Most streamers have a finite set
of states or modes or things they do. One moment they are gaming, then
another they are just chatting, rocking out, playing an instrument,
hiking outside, or coding and co-working while they do their day job.
Managing these modes can be hard for the streamer and the community.

## What is a *mode*?

A *mode* is type of activity with a specific topic at a given time. On
Twitch it is represented by a combination of the following:

* unique mode emoji (first character in status title)
* status (title) current message
* status (title) default message
* category
* tags
