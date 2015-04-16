# GeilMail

[![Build Status](https://travis-ci.org/GeilMail/geilmail.svg?branch=master)](https://travis-ci.org/GeilMail/geilmail)

## A word of warning

Please note that GeilMail is very much work in progress. There are tons of ugly code (first make it work, then make it nice), documentation is lacking very much and some things are broken or not existing.

## What is GeilMail aiming for?

GeilMail should make hosting an email server for you, your friends or your organisation as easy as possible. While there are a lot of solutions of setting up an email server, they all need a lot of documentation reading and understand about email. GeilMail is trying to be an all-in-one solution that can be installed and configured in 15 minutes.

It is not designed for endless scalability and thousands of users, because email is about decentrality and intercommunication and not huge clustering.

GeilMail will have SMTP and IMAP support with STARTTLS. There won't be support for legacy technology.

## Guidelines

### Test Coverage

A high test coverage is appreciated. In order to measure and inspect, type `go test -v -cover -covermode=count -coverprofile=cover.out && go tool cover -html=cover.out`
