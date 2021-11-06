# CLI Image Manager Phootecles

Phootecles is a Command Line Application that manages photo on your local machine. You can add an image, download it, store information such as title, description, online URL, you can extract the text from the image. You can search through your images, delete them, etc. It is written in Go.

Phootecles encodes object lists into json files, serializes maps, uses external packages and libraries such as Cobra, communicates with APIs to perform complex tasks, and uses Goroutines to run demanding operations in parallel.

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

### Creating new Images

Now try adding a new image by running

```bash
phootecles add "001" "Mount Fuji" "this active volcano is a very distinctive feature of the geography of Japan...." "https://upload.wikimedia.org/wikipedia/commons/1/1b/080103_hakkai_fuji.jpg" --download
```

You should get the responce:

```bash
Successfully added 002Mount Fuji
Successfully downloaded 002Mount Fuji
```

The `--download` flag means that Phootecles will download the image after adding it.

You can check that the image is indeed downloaded in the folder where Phootecles downloads new images; you can choose where you want that location to be; check Configuring Phootecles section.

Now check the image is added

```bash
phootecles get all
```

You should get the following responce:

```bash
id:001
title:Mount Fuji
description:this active volcano is a very distinctive feature of the geography of Japan....
url:https://upload.wikimedia.org/wikipedia/commons/1/1b/080103_hakkai_fuji.jpg
inside text:
---------------
Number of images is 1
```

#### Apply Optical Character Recognition to the image

When adding a new image, you can use the flag `--ocr` which will make Phootecles extract the text from the image and store it in its `inside text` attribute.

Later if you are searching in your images for a particular term, Phootecles will search the image's `inside text` as well as it's `title` and `description` for occurrences of that term.

Note: I use api `rapidapi.com/nadkabbani/api/ocrly-image-to-text/` to apply OCR to images and extract their text.

TODO: Look for another api such as `Google Cloud` that doesn't give only 50 calls per month in the free trial.

Note: The process of downloading an image and extracting its text through the aforementioned api run concurrently as I only the image url is sent to the api. I used Go routines to achieve this parallelism and gain performance increase as the api has around 3 seconds of latency.

### Configuring Phootecles

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

### Downloading The Images Again if You Want

You can use the command `download` to download the images that you added again when you want. So run:

```bash
phootecles download 001
```

You should get the following responce:

```bash
Successfully downloaded 001Mount Fuji
```

and you should find the image downloaded in the `imageSavePath`

You can also run:

```bash
phootecles download all
```

This will download all the images that are added.

### Deleting Images

You can delete an image using the command delete and passing the id of the image that you want to delete like this:

```bash
phootecles delete 001
```

If there was an image with the id `001` then it will be deleted.

You can also delete `all` images by typing all instead of the image id:

```bash
phootecles delete all
```

This will delete all the images that you have.
