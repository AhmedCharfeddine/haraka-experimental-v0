# Haraka

Fast latin to arabic transliteration

## Overview

Haraka is a lightweight tool that lets you type in Arabic by turning your "mar7abaa" to a "مرحبا"

Transliteration is the process transforming text from one alphabet to another based on phonetic similarity, and Haraka does that for you by converting what you're typing.

## Who's this for?

You can use Haraka if you speak arabic, mainly type in Latin and you either:
- Type faster with Latin layouts
- Use a keyboard with missing Arabic labels
- Are too lazy to swap keyboard language

## Build

The project assumes you have Go installed on your system. The build steps are:

1. Building and installing the project:
```
$ go build
$ go intall
```
3. Testing the CLI tool:
```
$ haraka --help
$ haraka map salaam
سلام
```

## Contributions

You're always welcome to contribute by opening issues for suggestions or submitting pull requests
