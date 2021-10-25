# CLI Image Manager Phootecles

Phootecles is a Command Line Application that manages photo collections on your local machine. It downloads images and stores information like description and online URL. It is written in Go and uses Cobra library.

## Insatlling Phootecles

Before installing Phootecles, you need to have Go installed first.

### Installing Go

To install the latest version of Go, follow the instructions on their [website](https://golang.org/dl/).

### Installing Phootecles

After installing go, you can clone this repo, open a terminal, navigate to the cloned folder, and enter the command

```bash
go install
```

Now check that Phootecles is installed by typing

```bash
phootecles
```

anywhere in your terminal.

## Using Phootecles

Now try adding a new image by running

```bash
phootecles add "001" "Mount Fuji" "this active volcano is a very distinctive feature of the geography of Japan...." "https://upload.wikimedia.org/wikipedia/commons/1/1b/080103_hakkai_fuji.jpg"
```

Now check the image is indeed added

```bash
phootecles get all
```

You should get the following responce:

```bash
id:001
title:Mount Fuji
description:this active volcano is a very distinctive feature of the geography of Japan....
url:https://upload.wikimedia.org/wikipedia/commons/1/1b/080103_hakkai_fuji.jpg
---------------
Number of images is 1
```

You can change the directory that images are saved to by running

```bash
phootecles list
```

You should get the following responce:

```bash
imageSavePath: ./
```

The list commands lists all variables that Phootecles uses that you can configure.

To change the `imageSavePath` you can run

```bash
phootecles update imageSavePath C:/Users/Hp/Desktop/
```

You should get the following responce:

```bash
Successfully updated imageSavePath to become C:/Users/Hp/Desktop/
```

Now when a new image is downloaded it will be saved in the desktop
